package group

import (
	"github.com/wzshiming/gen-examples/service/auth"
	"github.com/wzshiming/gen-examples/service/file"
	"github.com/wzshiming/gen-examples/service/item"
	"github.com/wzshiming/gen-examples/service/midd"
)

// Group #path:"/group"#
type Group struct {

	// #path:"/item"#
	Item item.ItemService
	// #path:"/file"#
	File file.FileService
	// #path:"/auth"#
	Auth auth.AuthService
	// #path:"/midd"#
	Midd midd.MiddService
}
