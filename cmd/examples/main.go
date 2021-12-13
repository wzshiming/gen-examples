package main

import (
	fmt "fmt"
	http "net/http"
	os "os"

	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	"github.com/wzshiming/gen-examples/openapi"
	"github.com/wzshiming/gen-examples/route"
	openapiui "github.com/wzshiming/openapiui"
	redoc "github.com/wzshiming/openapiui/redoc"
	swaggereditor "github.com/wzshiming/openapiui/swaggereditor"
	swaggerui "github.com/wzshiming/openapiui/swaggerui"
)

// RouteOpenAPI
func RouteOpenAPI(router *mux.Router) *mux.Router {
	openapi := map[string][]byte{
		"openapi.json": openapi.OpenAPI,
	}
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", openapiui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui", openapiui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggereditor/").Handler(http.StripPrefix("/swaggereditor", openapiui.HandleWithFiles(openapi, swaggereditor.Asset)))
	router.PathPrefix("/redoc/").Handler(http.StripPrefix("/redoc", openapiui.HandleWithFiles(openapi, redoc.Asset)))
	return router
}

func main() {
	m := route.Router()
	m = RouteOpenAPI(m.(*mux.Router))
	mux0 := handlers.RecoveryHandler()(m)
	mux0 = handlers.CombinedLoggingHandler(os.Stdout, mux0)

	fmt.Printf("Open http://127.0.0.1:8080/swagger/#\n")
	fmt.Printf("  or http://127.0.0.1:8080/swaggerui/#\n")
	fmt.Printf("  or http://127.0.0.1:8080/swaggereditor/#\n")
	fmt.Printf("  or http://127.0.0.1:8080/redoc/#\n")
	fmt.Printf("  with your browser.\n")

	err := http.ListenAndServe(":8080", mux0)
	if err != nil {
		fmt.Println(err)
	}
	return
}
