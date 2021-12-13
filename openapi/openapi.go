package openapi

import (
	_ "embed"
)

//go:embed openapi.json
var OpenAPI []byte
