package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <go-file>", os.Args[0])
	}

	filename := os.Args[1]

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	// Walk all declarations
	ast.Inspect(node, func(n ast.Node) bool {
		st, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		// Iterate through each field in the struct
		for _, field := range st.Fields.List {
			if len(field.Names) == 0 {
				// Embedded fields have no names; you can choose to skip them.
				continue
			}

			fieldName := field.Names[0].Name
			bsonTag := fmt.Sprintf(`bson:"%s"`, toBSONName(fieldName))

			// Handle tags
			if field.Tag == nil {
				// No tag → create one
				field.Tag = &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("`%s`", bsonTag),
				}
				continue
			}

			// Existing tag → append (if bson doesn't already exist)
			tag := field.Tag.Value // includes backticks
			if !strings.Contains(tag, "bson:") {
				tag = strings.Trim(tag, "`")
				tag = fmt.Sprintf("`%s %s`", tag, bsonTag)
				field.Tag.Value = tag
			}
		}

		return true
	})

	// Print the modified file back to original file
	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, node)
	if err != nil {
		log.Fatalf("Failed printing file: %v", err)
	}

	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed writing file: %v", err)
	}

	fmt.Println("BSON tags added successfully.")
}

// Converts field name to default bson naming: lowerCamelCase
func toBSONName(s string) string {
	if s == "" {
		return s
	}
	// Make first letter lower-case
	return strings.ToLower(s[:1]) + s[1:]
}
