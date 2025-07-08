package mangadex

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// --- generateKeyGob: Uses gob for binary serialization ---
// generateKeyGob creates a cache key using gob for binary serialization.
// It returns a hex-encoded string of the binary data.
func generateKeyGob(method string, args ...interface{}) string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	// Encode the method name
	if err := enc.Encode(method); err != nil {
		log.Printf("Error encoding method with gob: %v", err)
		return "" // Return empty string or handle error as appropriate for your application
	}

	// Encode the arguments
	if err := enc.Encode(args); err != nil {
		log.Printf("Error encoding args with gob: %v", err)
		return "" // Return empty string or handle error as appropriate for your application
	}
	return fmt.Sprintf("%x", buf.Bytes()) // Hex encode the binary data for a string key
}

// --- generateKeyHashed: Uses json.Marshal and then SHA256 hash ---
// generateKeyHashed creates a cache key by marshaling arguments to JSON
// and then hashing the resulting JSON string with SHA256.
func generateKeyHashed(method string, args ...interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		log.Printf("Error marshaling args to JSON: %v", err)
		// Return a key indicating an error, or handle as appropriate
		return fmt.Sprintf("%s:json_marshal_error:%s", method, err.Error())
	}

	hasher := sha256.New()
	hasher.Write(b)
	hash := hex.EncodeToString(hasher.Sum(nil))

	return fmt.Sprintf("%s:%s", method, hash)
}

// --- Test Data Structure ---
// ComplexArgs represents a typical set of arguments that might be passed
// to a cached function.
type ComplexArgs struct {
	UserID    int
	Query     string
	Enabled   bool
	Tags      []string
	Settings  map[string]interface{}
	NestedObj struct {
		ID   string
		Rate float64
	}
}

// init is used to register types with gob.
// This is crucial for gob to correctly encode/decode interface{} types.
func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{}) // For args ...interface{}
	gob.Register(ComplexArgs{})
	gob.Register(struct {
		ID   string
		Rate float64
	}{})
}

// --- Benchmark Functions ---

// BenchmarkGenerateKeyGob benchmarks the generateKeyGob function.
func BenchmarkGenerateKeyGob(b *testing.B) {
	testArgs := ComplexArgs{
		UserID:  12345,
		Query:   "example search query for data",
		Enabled: true,
		Tags:    []string{"tag1", "tag2", "tag3", "long-tag-name-for-testing"},
		Settings: map[string]interface{}{
			"limit":  100,
			"offset": 0,
			"filter": "active",
		},
		NestedObj: struct {
			ID   string
			Rate float64
		}{ID: "nested-item-abc", Rate: 99.99},
	}

	b.ResetTimer() // Reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		_ = generateKeyGob("GetUserData", testArgs, "some_extra_param", 42)
	}
}

// BenchmarkGenerateKeyHashed benchmarks the generateKeyHashed function.
func BenchmarkGenerateKeyHashed(b *testing.B) {
	testArgs := ComplexArgs{
		UserID:  12345,
		Query:   "example search query for data",
		Enabled: true,
		Tags:    []string{"tag1", "tag2", "tag3", "long-tag-name-for-testing"},
		Settings: map[string]interface{}{
			"limit":  100,
			"offset": 0,
			"filter": "active",
		},
		NestedObj: struct {
			ID   string
			Rate float64
		}{ID: "nested-item-abc", Rate: 99.99},
	}

	b.ResetTimer() // Reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		_ = generateKeyHashed("GetUserData", testArgs, "some_extra_param", 42)
	}
}
