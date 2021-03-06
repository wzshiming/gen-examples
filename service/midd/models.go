// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/wzshiming/gen-examples/service/middle`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

package midd

// Midd is a midd
type Midd struct {
	MiddID  int    `json:"midd_id"` // MiddID is the Midd ID
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}
