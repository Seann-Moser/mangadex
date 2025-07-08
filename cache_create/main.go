// gen_cache_wrapper.go
// Command-line tool to generate a cache wrapper for a given interface in a Go source file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

var tpl = template.Must(template.New("wrapper").Parse(`package {{.Package}}

import (
{{range .Imports}}	{{.}}
{{end}})

// Cached{{.IfaceName}} wraps {{.IfaceName}} adding caching layer.
type Cached{{.IfaceName}} struct {
	client {{.IfaceName}}
	cache  Cache
}

// NewCached{{.IfaceName}} constructs a new Cached{{.IfaceName}}.
func NewCached{{.IfaceName}}(client {{.IfaceName}}, cache Cache) *Cached{{.IfaceName}} {
	return &Cached{{.IfaceName}}{client: client, cache: cache}
}

// generateKey creates a cache key from the method name and argument values.
func generateKey(method string, args ...interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		// Handle error: perhaps log it and return a non-cacheable key or panic
		return fmt.Sprintf("%s:json_marshal_error:%s", method, err.Error())
	}

	hasher := sha256.New()
	hasher.Write(b)
	hash := hex.EncodeToString(hasher.Sum(nil))

	return fmt.Sprintf("%s:%s", method, hash)
}

{{range .Methods}}
// {{.Name}} applies caching before delegating to the underlying client.
func (c *Cached{{$.IfaceName}}) {{.Name}}({{.ParamDecls}}) ({{.ReturnType}}, error) {
	// Build cache key
	key := generateKey("{{.Name}}"{{- range .ArgNames}}, {{.}}{{- end }})
	if v, ok := c.cache.Get(key); ok {
		output := {{.VarType}}{}
		err := json.Unmarshal(v, &output)
		if err == nil{
			return {{if .IsPointer}}&{{end}}output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.{{.Name}}({{.CallArgs}})
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}
{{end}}
`))

func main() {
	inPath := flag.String("input", "", "input Go file containing interface")
	iface := flag.String("interface", "", "interface name to wrap")
	outPath := flag.String("output", "wrapper_cache.go", "output file for generated wrapper")
	pkg := flag.String("package", "cachedclient", "package name for generated file")
	flag.Parse()

	if *inPath == "" || *iface == "" {
		flag.Usage()
		os.Exit(1)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, *inPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parsing input: %v", err)
	}

	// Gather original imports

	type impSpec struct{ Alias, Path string }
	var origImports []impSpec
	for _, imp := range node.Imports {
		pathValue := strings.Trim(imp.Path.Value, `"`)
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		} else {
			alias = path.Base(pathValue)
		}
		origImports = append(origImports, impSpec{Alias: alias, Path: pathValue})
	}

	// Track used imports
	used := make(map[string]bool)

	// Standard imports always needed
	stdImps := []string{`"context"`, `"encoding/json"`, `"fmt"`, `"io"`, `"crypto/sha256"`, `"encoding/hex"`}

	type methodInfo struct {
		Name       string
		ParamDecls string
		ArgNames   []string
		CallArgs   string
		ReturnType string
		VarType    string
		IsPointer  bool
	}
	var methods []methodInfo

	for _, decl := range node.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts := spec.(*ast.TypeSpec)
			if ts.Name.Name != *iface {
				continue
			}
			ifaceType, ok := ts.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			for _, m := range ifaceType.Methods.List {
				fn := m.Type.(*ast.FuncType)
				// Inspect selectors for imports
				ast.Inspect(fn, func(n ast.Node) bool {
					if sel, ok := n.(*ast.SelectorExpr); ok {
						if x, ok := sel.X.(*ast.Ident); ok {
							for _, imp := range origImports {
								if x.Name == imp.Alias {
									used[imp.Path] = true
								}
							}
						}
					}
					return true
				})
				// Collect params and call args
				var decls []string
				var argNames, callArgs []string
				for _, param := range fn.Params.List {
					var typeBuf bytes.Buffer
					_ = format.Node(&typeBuf, fset, param.Type)
					typ := typeBuf.String()
					isVar := false
					if _, ok := param.Type.(*ast.Ellipsis); ok {
						isVar = true
					}
					for _, n := range param.Names {
						decls = append(decls, fmt.Sprintf("%s %s", n.Name, typ))
						argNames = append(argNames, n.Name)
						if isVar {
							callArgs = append(callArgs, n.Name+"...")
						} else {
							callArgs = append(callArgs, n.Name)
						}
					}
				}
				// Collect return imports
				if fn.Results != nil {
					for _, res := range fn.Results.List {
						ast.Inspect(res.Type, func(n ast.Node) bool {
							if sel, ok := n.(*ast.SelectorExpr); ok {
								if x, ok := sel.X.(*ast.Ident); ok {
									for _, imp := range origImports {
										if x.Name == imp.Alias {
											used[imp.Path] = true
										}
									}
								}
							}
							return true
						})
					}
				}
				// Determine return type string
				ret := ""
				if fn.Results != nil && len(fn.Results.List) > 0 {
					var retBuf bytes.Buffer
					_ = format.Node(&retBuf, fset, fn.Results.List[0].Type)
					ret = retBuf.String()
				}
				for _, name := range m.Names {
					methods = append(methods, methodInfo{
						Name:       name.Name,
						ParamDecls: strings.Join(decls, ", "),
						ArgNames:   argNames,
						CallArgs:   strings.Join(callArgs, ", "),
						ReturnType: ret,
						VarType:    strings.ReplaceAll(ret, "*", ""),
						IsPointer:  strings.Contains(ret, "*"),
					})
				}
			}
		}
	}

	// Build the final import list
	var imports []string
	imports = append(imports, stdImps...)
	for _, imp := range origImports {
		if used[imp.Path] {
			if imp.Alias != path.Base(imp.Path) {
				imports = append(imports, fmt.Sprintf("%s %q", imp.Alias, imp.Path))
			} else {
				imports = append(imports, fmt.Sprintf("%q", imp.Path))
			}
		}
	}

	tmplData := map[string]interface{}{
		"Package":   *pkg,
		"IfaceName": *iface,
		"Methods":   methods,
		"Imports":   imports,
	}

	var out bytes.Buffer
	if err := tpl.Execute(&out, tmplData); err != nil {
		log.Fatalf("template error: %v", err)
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		log.Printf("warning: could not format source: %v", err)
		formatted = out.Bytes()
	}

	if err := os.WriteFile(*outPath, formatted, 0644); err != nil {
		log.Fatalf("writing output: %v", err)
	}
	fmt.Printf("Generated %s successfully.\n", *outPath)
}
