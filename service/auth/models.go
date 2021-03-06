// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/wzshiming/gen-examples/service/auth`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

package auth

import (
	"net/http"
	"strconv"
)

// Auth is a auth
type Auth struct {
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// VerifyApiKey #security:"apiKey"#
func VerifyApiKey(token string /* #in:"header"# */) (userID uint64, err error) {
	return strconv.ParseUint(token, 0, 0)
}

// VerifyBasic #security:"basic"#
func VerifyBasic(r *http.Request) (userID uint64, err error) {
	u, _, _ := r.BasicAuth()
	return strconv.ParseUint(u, 0, 0)
}
