// Code generated; DO NOT EDIT.
// file ./route/route_gen.go

package route

import (
	bytes "bytes"
	json "encoding/json"
	errors "errors"
	"fmt"
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	strconv "strconv"
	strings "strings"

	mux "github.com/gorilla/mux"
	githubComWzshimingGenExamplesServiceAuth "github.com/wzshiming/gen-examples/service/auth"
	githubComWzshimingGenExamplesServiceErrresp "github.com/wzshiming/gen-examples/service/errresp"
	githubComWzshimingGenExamplesServiceFile "github.com/wzshiming/gen-examples/service/file"
	githubComWzshimingGenExamplesServiceGroup "github.com/wzshiming/gen-examples/service/group"
	githubComWzshimingGenExamplesServiceItem "github.com/wzshiming/gen-examples/service/item"
	githubComWzshimingGenExamplesServiceMidd "github.com/wzshiming/gen-examples/service/midd"
)

// notFoundHandler Is the not found of handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	err := fmt.Errorf("Not found '%s %s'", r.Method, r.URL.Path)

	_wrappingErrErrorResp(w, r, 404, err)

}

// Router is all routing for package
// generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// AuthService Define the method scope
	var _authService githubComWzshimingGenExamplesServiceAuth.AuthService
	RouteAuthService(router, &_authService)

	// FileService Define the method scope
	var _fileService githubComWzshimingGenExamplesServiceFile.FileService
	RouteFileService(router, &_fileService)

	// Group Define the method scope
	var _group githubComWzshimingGenExamplesServiceGroup.Group
	RouteGroup(router, &_group)

	// ItemService Define the method scope
	var _itemService githubComWzshimingGenExamplesServiceItem.ItemService
	RouteItemService(router, &_itemService)

	// MiddService Define the method scope
	var _middService githubComWzshimingGenExamplesServiceMidd.MiddService
	RouteMiddService(router, &_middService)

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return router
}

// RouteAuthService is routing for AuthService
func RouteAuthService(router *mux.Router, _authService *githubComWzshimingGenExamplesServiceAuth.AuthService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routeAuth := router.PathPrefix("/auth").Subrouter()
	if len(fs) != 0 {
		_routeAuth.Use(fs...)
	}

	// Registered routing POST /auth
	var __operationPostAuth http.Handler
	__operationPostAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostAuth(_authService, w, r)
	})
	_routeAuth.Methods("POST").Path("").Handler(__operationPostAuth)

	// Registered routing GET /auth
	var __operationGetAuth http.Handler
	__operationGetAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAuth(_authService, w, r)
	})
	_routeAuth.Methods("GET").Path("").Handler(__operationGetAuth)

	// Registered routing DELETE /auth/{auth_id}
	var __operationDeleteAuthAuthID http.Handler
	__operationDeleteAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteAuthAuthID(_authService, w, r)
	})
	_routeAuth.Methods("DELETE").Path("/{auth_id}").Handler(__operationDeleteAuthAuthID)

	// Registered routing GET /auth/{auth_id}
	var __operationGetAuthAuthID http.Handler
	__operationGetAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAuthAuthID(_authService, w, r)
	})
	_routeAuth.Methods("GET").Path("/{auth_id}").Handler(__operationGetAuthAuthID)

	// Registered routing PUT /auth/{auth_id}
	var __operationPutAuthAuthID http.Handler
	__operationPutAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutAuthAuthID(_authService, w, r)
	})
	_routeAuth.Methods("PUT").Path("/{auth_id}").Handler(__operationPutAuthAuthID)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// RouteFileService is routing for FileService
func RouteFileService(router *mux.Router, _fileService *githubComWzshimingGenExamplesServiceFile.FileService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routeFile := router.PathPrefix("/file").Subrouter()
	if len(fs) != 0 {
		_routeFile.Use(fs...)
	}

	// Registered routing POST /file
	var __operationPostFile http.Handler
	__operationPostFile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostFile(_fileService, w, r)
	})
	_routeFile.Methods("POST").Path("").Handler(__operationPostFile)

	// Registered routing GET /file/{filename}
	var __operationGetFileFilename http.Handler
	__operationGetFileFilename = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetFileFilename(_fileService, w, r)
	})
	_routeFile.Methods("GET").Path("/{filename}").Handler(__operationGetFileFilename)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// RouteGroup is routing for Group
func RouteGroup(router *mux.Router, _group *githubComWzshimingGenExamplesServiceGroup.Group, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routeGroupMidd := router.PathPrefix("/group/midd").Subrouter()
	if len(fs) != 0 {
		_routeGroupMidd.Use(fs...)
	}

	// Registered routing POST /group/midd
	var __operationPostGroupMidd http.Handler
	__operationPostGroupMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostGroupMidd(_group, w, r)
	})
	_routeGroupMidd.Methods("POST").Path("").Handler(__operationPostGroupMidd)

	// Registered routing GET /group/midd
	var __operationGetGroupMidd http.Handler
	__operationGetGroupMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupMidd(_group, w, r)
	})
	_routeGroupMidd.Methods("GET").Path("").Handler(__operationGetGroupMidd)

	// Registered routing DELETE /group/midd/{midd_id}
	var __operationDeleteGroupMiddMiddID http.Handler
	__operationDeleteGroupMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteGroupMiddMiddID(_group, w, r)
	})
	_routeGroupMidd.Methods("DELETE").Path("/{midd_id}").Handler(__operationDeleteGroupMiddMiddID)

	// Registered routing GET /group/midd/{midd_id}
	var __operationGetGroupMiddMiddID http.Handler
	__operationGetGroupMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupMiddMiddID(_group, w, r)
	})
	_routeGroupMidd.Methods("GET").Path("/{midd_id}").Handler(__operationGetGroupMiddMiddID)

	// Registered routing PUT /group/midd/{midd_id}
	var __operationPutGroupMiddMiddID http.Handler
	__operationPutGroupMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutGroupMiddMiddID(_group, w, r)
	})
	_routeGroupMidd.Methods("PUT").Path("/{midd_id}").Handler(__operationPutGroupMiddMiddID)

	_routeGroupItem := router.PathPrefix("/group/item").Subrouter()
	if len(fs) != 0 {
		_routeGroupItem.Use(fs...)
	}

	// Registered routing POST /group/item
	var __operationPostGroupItem http.Handler
	__operationPostGroupItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostGroupItem(_group, w, r)
	})
	_routeGroupItem.Methods("POST").Path("").Handler(__operationPostGroupItem)

	// Registered routing GET /group/item
	var __operationGetGroupItem http.Handler
	__operationGetGroupItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupItem(_group, w, r)
	})
	_routeGroupItem.Methods("GET").Path("").Handler(__operationGetGroupItem)

	// Registered routing DELETE /group/item/{item_id}
	var __operationDeleteGroupItemItemID http.Handler
	__operationDeleteGroupItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteGroupItemItemID(_group, w, r)
	})
	_routeGroupItem.Methods("DELETE").Path("/{item_id}").Handler(__operationDeleteGroupItemItemID)

	// Registered routing GET /group/item/{item_id}
	var __operationGetGroupItemItemID http.Handler
	__operationGetGroupItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupItemItemID(_group, w, r)
	})
	_routeGroupItem.Methods("GET").Path("/{item_id}").Handler(__operationGetGroupItemItemID)

	// Registered routing PUT /group/item/{item_id}
	var __operationPutGroupItemItemID http.Handler
	__operationPutGroupItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutGroupItemItemID(_group, w, r)
	})
	_routeGroupItem.Methods("PUT").Path("/{item_id}").Handler(__operationPutGroupItemItemID)

	_routeGroupFile := router.PathPrefix("/group/file").Subrouter()
	if len(fs) != 0 {
		_routeGroupFile.Use(fs...)
	}

	// Registered routing POST /group/file
	var __operationPostGroupFile http.Handler
	__operationPostGroupFile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostGroupFile(_group, w, r)
	})
	_routeGroupFile.Methods("POST").Path("").Handler(__operationPostGroupFile)

	// Registered routing GET /group/file/{filename}
	var __operationGetGroupFileFilename http.Handler
	__operationGetGroupFileFilename = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupFileFilename(_group, w, r)
	})
	_routeGroupFile.Methods("GET").Path("/{filename}").Handler(__operationGetGroupFileFilename)

	_routeGroupAuth := router.PathPrefix("/group/auth").Subrouter()
	if len(fs) != 0 {
		_routeGroupAuth.Use(fs...)
	}

	// Registered routing POST /group/auth
	var __operationPostGroupAuth http.Handler
	__operationPostGroupAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostGroupAuth(_group, w, r)
	})
	_routeGroupAuth.Methods("POST").Path("").Handler(__operationPostGroupAuth)

	// Registered routing GET /group/auth
	var __operationGetGroupAuth http.Handler
	__operationGetGroupAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupAuth(_group, w, r)
	})
	_routeGroupAuth.Methods("GET").Path("").Handler(__operationGetGroupAuth)

	// Registered routing DELETE /group/auth/{auth_id}
	var __operationDeleteGroupAuthAuthID http.Handler
	__operationDeleteGroupAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteGroupAuthAuthID(_group, w, r)
	})
	_routeGroupAuth.Methods("DELETE").Path("/{auth_id}").Handler(__operationDeleteGroupAuthAuthID)

	// Registered routing GET /group/auth/{auth_id}
	var __operationGetGroupAuthAuthID http.Handler
	__operationGetGroupAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetGroupAuthAuthID(_group, w, r)
	})
	_routeGroupAuth.Methods("GET").Path("/{auth_id}").Handler(__operationGetGroupAuthAuthID)

	// Registered routing PUT /group/auth/{auth_id}
	var __operationPutGroupAuthAuthID http.Handler
	__operationPutGroupAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutGroupAuthAuthID(_group, w, r)
	})
	_routeGroupAuth.Methods("PUT").Path("/{auth_id}").Handler(__operationPutGroupAuthAuthID)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// RouteItemService is routing for ItemService
func RouteItemService(router *mux.Router, _itemService *githubComWzshimingGenExamplesServiceItem.ItemService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routeItem := router.PathPrefix("/item").Subrouter()
	if len(fs) != 0 {
		_routeItem.Use(fs...)
	}

	// Registered routing POST /item
	var __operationPostItem http.Handler
	__operationPostItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostItem(_itemService, w, r)
	})
	_routeItem.Methods("POST").Path("").Handler(__operationPostItem)

	// Registered routing GET /item
	var __operationGetItem http.Handler
	__operationGetItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetItem(_itemService, w, r)
	})
	_routeItem.Methods("GET").Path("").Handler(__operationGetItem)

	// Registered routing DELETE /item/{item_id}
	var __operationDeleteItemItemID http.Handler
	__operationDeleteItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteItemItemID(_itemService, w, r)
	})
	_routeItem.Methods("DELETE").Path("/{item_id}").Handler(__operationDeleteItemItemID)

	// Registered routing GET /item/{item_id}
	var __operationGetItemItemID http.Handler
	__operationGetItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetItemItemID(_itemService, w, r)
	})
	_routeItem.Methods("GET").Path("/{item_id}").Handler(__operationGetItemItemID)

	// Registered routing PUT /item/{item_id}
	var __operationPutItemItemID http.Handler
	__operationPutItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutItemItemID(_itemService, w, r)
	})
	_routeItem.Methods("PUT").Path("/{item_id}").Handler(__operationPutItemItemID)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// RouteMiddService is routing for MiddService
func RouteMiddService(router *mux.Router, _middService *githubComWzshimingGenExamplesServiceMidd.MiddService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routeMidd := router.PathPrefix("/midd").Subrouter()
	if len(fs) != 0 {
		_routeMidd.Use(fs...)
	}

	// Registered routing POST /midd
	var __operationPostMidd http.Handler
	__operationPostMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostMidd(_middService, w, r)
	})
	_routeMidd.Methods("POST").Path("").Handler(__operationPostMidd)

	// Registered routing GET /midd
	var __operationGetMidd http.Handler
	__operationGetMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetMidd(_middService, w, r)
	})
	_routeMidd.Methods("GET").Path("").Handler(__operationGetMidd)

	// Registered routing DELETE /midd/{midd_id}
	var __operationDeleteMiddMiddID http.Handler
	__operationDeleteMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteMiddMiddID(_middService, w, r)
	})
	_routeMidd.Methods("DELETE").Path("/{midd_id}").Handler(__operationDeleteMiddMiddID)

	// Registered routing GET /midd/{midd_id}
	var __operationGetMiddMiddID http.Handler
	__operationGetMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetMiddMiddID(_middService, w, r)
	})
	_routeMidd.Methods("GET").Path("/{midd_id}").Handler(__operationGetMiddMiddID)

	// Registered routing PUT /midd/{midd_id}
	var __operationPutMiddMiddID http.Handler
	__operationPutMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutMiddMiddID(_middService, w, r)
	})
	_routeMidd.Methods("PUT").Path("/{midd_id}").Handler(__operationPutMiddMiddID)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// _requestBodyAuth Parsing the body for of auth
func _requestBodyAuth(w http.ResponseWriter, r *http.Request) (_auth *githubComWzshimingGenExamplesServiceAuth.Auth, err error) {

	defer r.Body.Close()

	var __auth []byte
	__auth, err = ioutil.ReadAll(r.Body)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	err = json.Unmarshal(__auth, &_auth)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestPathAuthID Parsing the path for of auth_id
func _requestPathAuthID(w http.ResponseWriter, r *http.Request) (_authID int, err error) {

	var _rawAuthID = mux.Vars(r)["auth_id"]

	if __authID, err := strconv.ParseInt(_rawAuthID, 0, 0); err == nil {
		_authID = int(__authID)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestBodyFile Parsing the body for of file
func _requestBodyFile(w http.ResponseWriter, r *http.Request) (_file io.Reader, err error) {

	defer r.Body.Close()

	body := r.Body
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		if r.MultipartForm == nil {
			err = r.ParseMultipartForm(10 << 20)
			if err != nil {
				return
			}
		}
		file := r.MultipartForm.File["file"]
		if len(file) != 0 {
			body, err = file[0].Open()
			if err != nil {
				return
			}
		}
	}

	__file := bytes.NewBuffer(nil)
	_, err = io.Copy(__file, body)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	_file = __file

	return
}

// _requestPathFilename Parsing the path for of filename
func _requestPathFilename(w http.ResponseWriter, r *http.Request) (_filename string, err error) {

	var _rawFilename = mux.Vars(r)["filename"]
	_filename = string(_rawFilename)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestBodyItem Parsing the body for of item
func _requestBodyItem(w http.ResponseWriter, r *http.Request) (_item *githubComWzshimingGenExamplesServiceItem.Item, err error) {

	defer r.Body.Close()

	var __item []byte
	__item, err = ioutil.ReadAll(r.Body)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	err = json.Unmarshal(__item, &_item)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestPathItemID Parsing the path for of item_id
func _requestPathItemID(w http.ResponseWriter, r *http.Request) (_itemID int, err error) {

	var _rawItemID = mux.Vars(r)["item_id"]

	if __itemID, err := strconv.ParseInt(_rawItemID, 0, 0); err == nil {
		_itemID = int(__itemID)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryLimit Parsing the query for of limit
func _requestQueryLimit(w http.ResponseWriter, r *http.Request) (_limit int, err error) {

	var _rawLimit = r.URL.Query()["limit"]

	if len(_rawLimit) == 0 {
		return
	}

	__rawLimit_0 := _rawLimit[0]

	if __limit, err := strconv.ParseInt(__rawLimit_0, 0, 0); err == nil {
		_limit = int(__limit)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryLimit_1 Parsing the query for of limit
func _requestQueryLimit_1(w http.ResponseWriter, r *http.Request) (_limit_1 int, err error) {

	var _rawLimit1 = r.URL.Query()["limit"]

	if len(_rawLimit1) == 0 {
		return
	}

	__rawLimit1_0 := _rawLimit1[0]

	if __limit_1, err := strconv.ParseInt(__rawLimit1_0, 0, 0); err == nil {
		_limit_1 = int(__limit_1)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryLimit_2 Parsing the query for of limit
func _requestQueryLimit_2(w http.ResponseWriter, r *http.Request) (_limit_2 int, err error) {

	var _rawLimit2 = r.URL.Query()["limit"]

	if len(_rawLimit2) == 0 {
		return
	}

	__rawLimit2_0 := _rawLimit2[0]

	if __limit_2, err := strconv.ParseInt(__rawLimit2_0, 0, 0); err == nil {
		_limit_2 = int(__limit_2)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestBodyMidd Parsing the body for of midd
func _requestBodyMidd(w http.ResponseWriter, r *http.Request) (_midd *githubComWzshimingGenExamplesServiceMidd.Midd, err error) {

	defer r.Body.Close()

	var __midd []byte
	__midd, err = ioutil.ReadAll(r.Body)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	err = json.Unmarshal(__midd, &_midd)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestPathMiddID Parsing the path for of midd_id
func _requestPathMiddID(w http.ResponseWriter, r *http.Request) (_middID int, err error) {

	var _rawMiddID = mux.Vars(r)["midd_id"]

	if __middID, err := strconv.ParseInt(_rawMiddID, 0, 0); err == nil {
		_middID = int(__middID)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryOffset Parsing the query for of offset
func _requestQueryOffset(w http.ResponseWriter, r *http.Request) (_offset int, err error) {

	var _rawOffset = r.URL.Query()["offset"]

	if len(_rawOffset) == 0 {
		return
	}

	__rawOffset_0 := _rawOffset[0]

	if __offset, err := strconv.ParseInt(__rawOffset_0, 0, 0); err == nil {
		_offset = int(__offset)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryOffset_1 Parsing the query for of offset
func _requestQueryOffset_1(w http.ResponseWriter, r *http.Request) (_offset_1 int, err error) {

	var _rawOffset1 = r.URL.Query()["offset"]

	if len(_rawOffset1) == 0 {
		return
	}

	__rawOffset1_0 := _rawOffset1[0]

	if __offset_1, err := strconv.ParseInt(__rawOffset1_0, 0, 0); err == nil {
		_offset_1 = int(__offset_1)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestQueryOffset_2 Parsing the query for of offset
func _requestQueryOffset_2(w http.ResponseWriter, r *http.Request) (_offset_2 int, err error) {

	var _rawOffset2 = r.URL.Query()["offset"]

	if len(_rawOffset2) == 0 {
		return
	}

	__rawOffset2_0 := _rawOffset2[0]

	if __offset_2, err := strconv.ParseInt(__rawOffset2_0, 0, 0); err == nil {
		_offset_2 = int(__offset_2)
	}

	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestHeaderToken Parsing the header for of token
func _requestHeaderToken(w http.ResponseWriter, r *http.Request) (_token string, err error) {

	var _rawToken = r.Header.Get("token")
	_token = string(_rawToken)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _requestHeaderXToken Parsing the header for of x-token
func _requestHeaderXToken(w http.ResponseWriter, r *http.Request) (_xToken string, err error) {

	var _rawXToken = r.Header.Get("x-token")
	_xToken = string(_rawXToken)
	if err != nil {
		_wrappingErrErrorResp(w, r, 400, err)

		return
	}

	return
}

// _middlewareSessionMiddelSession Is the middleware of MiddelSession
func _middlewareSessionMiddelSession(w http.ResponseWriter, r *http.Request) (_session *githubComWzshimingGenExamplesServiceMidd.Session, _err error) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddelSession.x-token
	var _xToken string

	// Parsing x-token.
	_xToken, _err = _requestHeaderXToken(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddelSession.
	_session, _err = githubComWzshimingGenExamplesServiceMidd.MiddelSession(_xToken)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	return
}

// _middlewareTokenMiddelToken Is the middleware of MiddelToken
func _middlewareTokenMiddelToken(w http.ResponseWriter, r *http.Request) (_token_1 string, _err error) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddelToken.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddelToken.
	_token_1, _err = githubComWzshimingGenExamplesServiceMidd.MiddelToken(_session_1)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	return
}

// _wrappingErrErrorResp Is the wrapping of ErrorResp
func _wrappingErrErrorResp(w http.ResponseWriter, r *http.Request, _code int, _rawErr error) {
	// responses github.com/wzshiming/gen-examples/service/errresp ErrorResp.e
	var _e *githubComWzshimingGenExamplesServiceErrresp.Error
	// responses github.com/wzshiming/gen-examples/service/errresp ErrorResp.err
	var _err_1 error

	// Call github.com/wzshiming/gen-examples/service/errresp ErrorResp.
	_e, _err_1 = githubComWzshimingGenExamplesServiceErrresp.ErrorResp(_rawErr)

	var __e []byte
	__e, _err_1 = json.Marshal(_e)
	if _err_1 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_1)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(_code)
	w.Write(__e)

	return
}

// _securityAPIKeyVerifyAPIKey Is the security of VerifyApiKey
func _securityAPIKeyVerifyAPIKey(w http.ResponseWriter, r *http.Request) (_userID uint64, _err_2 error) {
	// requests github.com/wzshiming/gen-examples/service/auth VerifyApiKey.token
	var _token string

	// Parsing token.
	_token, _err_2 = _requestHeaderToken(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth VerifyApiKey.
	_userID, _err_2 = githubComWzshimingGenExamplesServiceAuth.VerifyApiKey(_token)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	return
}

// _securityBasicVerifyBasic Is the security of VerifyBasic
func _securityBasicVerifyBasic(w http.ResponseWriter, r *http.Request) (_userID uint64, _err_2 error) {

	// Call github.com/wzshiming/gen-examples/service/auth VerifyBasic.
	_userID, _err_2 = githubComWzshimingGenExamplesServiceAuth.VerifyBasic(r)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	return
}

// _operationPutAuthAuthID Is the route of Update
func _operationPutAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Update.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Update.auth_id
	var _authID int
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Update.auth
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.Update.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Parsing auth.
	_auth, _err_2 = _requestBodyAuth(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Update.
	_err_2 = s.Update(_userID_1, _authID, _auth)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetAuth Is the route of List
func _operationGetAuth(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.List.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.List.offset
	var _offset int
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.List.limit
	var _limit int
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.List.auths
	var _auths []*githubComWzshimingGenExamplesServiceAuth.AuthWithID
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.List.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing offset.
	_offset, _err_2 = _requestQueryOffset(w, r)
	if _err_2 != nil {
		return
	}

	// Parsing limit.
	_limit, _err_2 = _requestQueryLimit(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.List.
	_auths, _err_2 = s.List(_userID_1, _offset, _limit)

	// Response code 200 OK for auths.
	if _auths != nil {
		var __auths []byte
		__auths, _err_2 = json.Marshal(_auths)
		if _err_2 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_2)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auths)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	var __auths []byte
	__auths, _err_2 = json.Marshal(_auths)
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_2)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auths)

	return
}

// _operationGetAuthAuthID Is the route of Get
func _operationGetAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Get.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Get.auth_id
	var _authID int
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.Get.auth
	var _auth_1 *githubComWzshimingGenExamplesServiceAuth.AuthWithID
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.Get.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Get.
	_auth_1, _err_2 = s.Get(_userID_1, _authID)

	// Response code 200 OK for auth.
	if _auth_1 != nil {
		var __auth_1 []byte
		__auth_1, _err_2 = json.Marshal(_auth_1)
		if _err_2 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_2)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auth_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	var __auth_1 []byte
	__auth_1, _err_2 = json.Marshal(_auth_1)
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_2)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auth_1)

	return
}

// _operationDeleteAuthAuthID Is the route of Delete
func _operationDeleteAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Delete.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Delete.auth_id
	var _authID int
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.Delete.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Delete.
	_err_2 = s.Delete(_userID_1, _authID)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostAuth Is the route of Create
func _operationPostAuth(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Create.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth AuthService.Create.auth
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth
	// responses github.com/wzshiming/gen-examples/service/auth AuthService.Create.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth.
	_auth, _err_2 = _requestBodyAuth(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Create.
	_err_2 = s.Create(_userID_1, _auth)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostFile Is the route of Upload
func _operationPostFile(s *githubComWzshimingGenExamplesServiceFile.FileService, w http.ResponseWriter, r *http.Request) {
	// requests  FileService.Upload.file
	var _file io.Reader
	// responses github.com/wzshiming/gen-examples/service/file FileService.Upload.filename
	var _filename_1 string
	// responses github.com/wzshiming/gen-examples/service/file FileService.Upload.err
	var _err_3 error

	// Parsing file.
	_file, _err_3 = _requestBodyFile(w, r)
	if _err_3 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Upload.
	_filename_1, _err_3 = s.Upload(_file)

	// Response code 200 OK for filename.
	if _filename_1 != "" {
		var __filename_1 []byte
		__filename_1, _err_3 = json.Marshal(_filename_1)
		if _err_3 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_3)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__filename_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_3)
		return
	}

	var __filename_1 []byte
	__filename_1, _err_3 = json.Marshal(_filename_1)
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_3)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__filename_1)

	return
}

// _operationGetFileFilename Is the route of Get
func _operationGetFileFilename(s *githubComWzshimingGenExamplesServiceFile.FileService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/file FileService.Get.filename
	var _filename string
	// responses github.com/wzshiming/gen-examples/service/file FileService.Get.file
	var _file_1 []byte
	// responses github.com/wzshiming/gen-examples/service/file FileService.Get.err
	var _err_3 error

	// Parsing filename.
	_filename, _err_3 = _requestPathFilename(w, r)
	if _err_3 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Get.
	_file_1, _err_3 = s.Get(_filename)

	// Response code 200 OK for file.
	if _file_1 != nil {
		var __file_1 []byte
		__file_1 = _file_1

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		w.Write(__file_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_3)
		return
	}

	var __file_1 []byte
	__file_1 = _file_1

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(__file_1)

	return
}

// _operationPutGroupItemItemID Is the route of Update
func _operationPutGroupItemItemID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item Group.Update.item_id
	var _itemID int
	// requests github.com/wzshiming/gen-examples/service/item Group.Update.item
	var _item *githubComWzshimingGenExamplesServiceItem.Item
	// responses github.com/wzshiming/gen-examples/service/item Group.Update.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Parsing item.
	_item, _err_4 = _requestBodyItem(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item Group.Item.Update.
	_err_4 = s.Item.Update(_itemID, _item)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetGroupItem Is the route of List
func _operationGetGroupItem(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item Group.List.offset
	var _offset_1 int
	// requests github.com/wzshiming/gen-examples/service/item Group.List.limit
	var _limit_1 int
	// responses github.com/wzshiming/gen-examples/service/item Group.List.items
	var _items []*githubComWzshimingGenExamplesServiceItem.ItemWithID
	// responses github.com/wzshiming/gen-examples/service/item Group.List.err
	var _err_4 error

	// Parsing offset.
	_offset_1, _err_4 = _requestQueryOffset_1(w, r)
	if _err_4 != nil {
		return
	}

	// Parsing limit.
	_limit_1, _err_4 = _requestQueryLimit_1(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item Group.Item.List.
	_items, _err_4 = s.Item.List(_offset_1, _limit_1)

	// Response code 200 OK for items.
	if _items != nil {
		var __items []byte
		__items, _err_4 = json.Marshal(_items)
		if _err_4 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_4)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__items)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	var __items []byte
	__items, _err_4 = json.Marshal(_items)
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_4)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__items)

	return
}

// _operationGetGroupItemItemID Is the route of Get
func _operationGetGroupItemItemID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item Group.Get.item_id
	var _itemID int
	// responses github.com/wzshiming/gen-examples/service/item Group.Get.item
	var _item_1 *githubComWzshimingGenExamplesServiceItem.ItemWithID
	// responses github.com/wzshiming/gen-examples/service/item Group.Get.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item Group.Item.Get.
	_item_1, _err_4 = s.Item.Get(_itemID)

	// Response code 200 OK for item.
	if _item_1 != nil {
		var __item_1 []byte
		__item_1, _err_4 = json.Marshal(_item_1)
		if _err_4 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_4)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__item_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	var __item_1 []byte
	__item_1, _err_4 = json.Marshal(_item_1)
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_4)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__item_1)

	return
}

// _operationDeleteGroupItemItemID Is the route of Delete
func _operationDeleteGroupItemItemID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item Group.Delete.item_id
	var _itemID int
	// responses github.com/wzshiming/gen-examples/service/item Group.Delete.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item Group.Item.Delete.
	_err_4 = s.Item.Delete(_itemID)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostGroupItem Is the route of Create
func _operationPostGroupItem(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item Group.Create.item
	var _item *githubComWzshimingGenExamplesServiceItem.Item
	// responses github.com/wzshiming/gen-examples/service/item Group.Create.err
	var _err_4 error

	// Parsing item.
	_item, _err_4 = _requestBodyItem(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item Group.Item.Create.
	_err_4 = s.Item.Create(_item)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostGroupFile Is the route of Upload
func _operationPostGroupFile(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests  Group.Upload.file
	var _file io.Reader
	// responses github.com/wzshiming/gen-examples/service/file Group.Upload.filename
	var _filename_1 string
	// responses github.com/wzshiming/gen-examples/service/file Group.Upload.err
	var _err_3 error

	// Parsing file.
	_file, _err_3 = _requestBodyFile(w, r)
	if _err_3 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file Group.File.Upload.
	_filename_1, _err_3 = s.File.Upload(_file)

	// Response code 200 OK for filename.
	if _filename_1 != "" {
		var __filename_1 []byte
		__filename_1, _err_3 = json.Marshal(_filename_1)
		if _err_3 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_3)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__filename_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_3)
		return
	}

	var __filename_1 []byte
	__filename_1, _err_3 = json.Marshal(_filename_1)
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_3)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__filename_1)

	return
}

// _operationGetGroupFileFilename Is the route of Get
func _operationGetGroupFileFilename(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/file Group.Get.filename
	var _filename string
	// responses github.com/wzshiming/gen-examples/service/file Group.Get.file
	var _file_1 []byte
	// responses github.com/wzshiming/gen-examples/service/file Group.Get.err
	var _err_3 error

	// Parsing filename.
	_filename, _err_3 = _requestPathFilename(w, r)
	if _err_3 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file Group.File.Get.
	_file_1, _err_3 = s.File.Get(_filename)

	// Response code 200 OK for file.
	if _file_1 != nil {
		var __file_1 []byte
		__file_1 = _file_1

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		w.Write(__file_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_3 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_3)
		return
	}

	var __file_1 []byte
	__file_1 = _file_1

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(__file_1)

	return
}

// _operationPutGroupAuthAuthID Is the route of Update
func _operationPutGroupAuthAuthID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth Group.Update.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth Group.Update.auth_id
	var _authID int
	// requests github.com/wzshiming/gen-examples/service/auth Group.Update.auth
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth
	// responses github.com/wzshiming/gen-examples/service/auth Group.Update.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Parsing auth.
	_auth, _err_2 = _requestBodyAuth(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Group.Auth.Update.
	_err_2 = s.Auth.Update(_userID_1, _authID, _auth)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetGroupAuth Is the route of List
func _operationGetGroupAuth(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth Group.List.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth Group.List.offset
	var _offset int
	// requests github.com/wzshiming/gen-examples/service/auth Group.List.limit
	var _limit int
	// responses github.com/wzshiming/gen-examples/service/auth Group.List.auths
	var _auths []*githubComWzshimingGenExamplesServiceAuth.AuthWithID
	// responses github.com/wzshiming/gen-examples/service/auth Group.List.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing offset.
	_offset, _err_2 = _requestQueryOffset(w, r)
	if _err_2 != nil {
		return
	}

	// Parsing limit.
	_limit, _err_2 = _requestQueryLimit(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Group.Auth.List.
	_auths, _err_2 = s.Auth.List(_userID_1, _offset, _limit)

	// Response code 200 OK for auths.
	if _auths != nil {
		var __auths []byte
		__auths, _err_2 = json.Marshal(_auths)
		if _err_2 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_2)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auths)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	var __auths []byte
	__auths, _err_2 = json.Marshal(_auths)
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_2)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auths)

	return
}

// _operationGetGroupAuthAuthID Is the route of Get
func _operationGetGroupAuthAuthID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth Group.Get.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth Group.Get.auth_id
	var _authID int
	// responses github.com/wzshiming/gen-examples/service/auth Group.Get.auth
	var _auth_1 *githubComWzshimingGenExamplesServiceAuth.AuthWithID
	// responses github.com/wzshiming/gen-examples/service/auth Group.Get.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Group.Auth.Get.
	_auth_1, _err_2 = s.Auth.Get(_userID_1, _authID)

	// Response code 200 OK for auth.
	if _auth_1 != nil {
		var __auth_1 []byte
		__auth_1, _err_2 = json.Marshal(_auth_1)
		if _err_2 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_2)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auth_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	var __auth_1 []byte
	__auth_1, _err_2 = json.Marshal(_auth_1)
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_2)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auth_1)

	return
}

// _operationDeleteGroupAuthAuthID Is the route of Delete
func _operationDeleteGroupAuthAuthID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth Group.Delete.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth Group.Delete.auth_id
	var _authID int
	// responses github.com/wzshiming/gen-examples/service/auth Group.Delete.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth_id.
	_authID, _err_2 = _requestPathAuthID(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Group.Auth.Delete.
	_err_2 = s.Auth.Delete(_userID_1, _authID)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostGroupAuth Is the route of Create
func _operationPostGroupAuth(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/auth Group.Create.userID
	var _userID_1 uint64
	// requests github.com/wzshiming/gen-examples/service/auth Group.Create.auth
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth
	// responses github.com/wzshiming/gen-examples/service/auth Group.Create.err
	var _err_2 error

	// Permission verification
	if r.Header.Get("token") != "" { // Call VerifyApiKey.
		_userID_1, _err_2 = _securityAPIKeyVerifyAPIKey(w, r)
	} else if strings.HasPrefix(r.Header.Get("Authorization"), "Basic ") { // Call VerifyBasic.
		_userID_1, _err_2 = _securityBasicVerifyBasic(w, r)
	} else {
		_err_2 = errors.New("Unauthorized")
	}
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 401, _err_2)

		return
	}

	// Parsing auth.
	_auth, _err_2 = _requestBodyAuth(w, r)
	if _err_2 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Group.Auth.Create.
	_err_2 = s.Auth.Create(_userID_1, _auth)

	// Response code 400 Bad Request for err.
	if _err_2 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_2)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPutGroupMiddMiddID Is the route of Update
func _operationPutGroupMiddMiddID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd Group.Update.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd Group.Update.midd_id
	var _middID int
	// requests github.com/wzshiming/gen-examples/service/midd Group.Update.midd
	var _midd *githubComWzshimingGenExamplesServiceMidd.Midd
	// responses github.com/wzshiming/gen-examples/service/midd Group.Update.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Parsing midd.
	_midd, _err = _requestBodyMidd(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd Group.Midd.Update.
	_err = s.Midd.Update(_session_1, _middID, _midd)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetGroupMidd Is the route of List
func _operationGetGroupMidd(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd Group.List.token
	var _token_2 string
	// requests github.com/wzshiming/gen-examples/service/midd Group.List.offset
	var _offset_2 int
	// requests github.com/wzshiming/gen-examples/service/midd Group.List.limit
	var _limit_2 int
	// responses github.com/wzshiming/gen-examples/service/midd Group.List.midds
	var _midds []*githubComWzshimingGenExamplesServiceMidd.MiddWithID
	// responses github.com/wzshiming/gen-examples/service/midd Group.List.err
	var _err error

	// Permission middlewares call MiddelToken.
	_token_2, _err = _middlewareTokenMiddelToken(w, r)
	if _err != nil {
		return
	}

	// Parsing offset.
	_offset_2, _err = _requestQueryOffset_2(w, r)
	if _err != nil {
		return
	}

	// Parsing limit.
	_limit_2, _err = _requestQueryLimit_2(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd Group.Midd.List.
	_midds, _err = s.Midd.List(_token_2, _offset_2, _limit_2)

	// Response code 200 OK for midds.
	if _midds != nil {
		var __midds []byte
		__midds, _err = json.Marshal(_midds)
		if _err != nil {
			_wrappingErrErrorResp(w, r, 500, _err)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midds)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	var __midds []byte
	__midds, _err = json.Marshal(_midds)
	if _err != nil {
		_wrappingErrErrorResp(w, r, 500, _err)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midds)

	return
}

// _operationGetGroupMiddMiddID Is the route of Get
func _operationGetGroupMiddMiddID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd Group.Get.token
	var _token_2 string
	// requests github.com/wzshiming/gen-examples/service/midd Group.Get.midd_id
	var _middID int
	// responses github.com/wzshiming/gen-examples/service/midd Group.Get.midd
	var _midd_1 *githubComWzshimingGenExamplesServiceMidd.MiddWithID
	// responses github.com/wzshiming/gen-examples/service/midd Group.Get.err
	var _err error

	// Permission middlewares call MiddelToken.
	_token_2, _err = _middlewareTokenMiddelToken(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd Group.Midd.Get.
	_midd_1, _err = s.Midd.Get(_token_2, _middID)

	// Response code 200 OK for midd.
	if _midd_1 != nil {
		var __midd_1 []byte
		__midd_1, _err = json.Marshal(_midd_1)
		if _err != nil {
			_wrappingErrErrorResp(w, r, 500, _err)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midd_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	var __midd_1 []byte
	__midd_1, _err = json.Marshal(_midd_1)
	if _err != nil {
		_wrappingErrErrorResp(w, r, 500, _err)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midd_1)

	return
}

// _operationDeleteGroupMiddMiddID Is the route of Delete
func _operationDeleteGroupMiddMiddID(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd Group.Delete.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd Group.Delete.midd_id
	var _middID int
	// responses github.com/wzshiming/gen-examples/service/midd Group.Delete.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd Group.Midd.Delete.
	_err = s.Midd.Delete(_session_1, _middID)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostGroupMidd Is the route of Create
func _operationPostGroupMidd(s *githubComWzshimingGenExamplesServiceGroup.Group, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd Group.Create.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd Group.Create.midd
	var _midd *githubComWzshimingGenExamplesServiceMidd.Midd
	// responses github.com/wzshiming/gen-examples/service/midd Group.Create.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd.
	_midd, _err = _requestBodyMidd(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd Group.Midd.Create.
	_err = s.Midd.Create(_session_1, _midd)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPutItemItemID Is the route of Update
func _operationPutItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item ItemService.Update.item_id
	var _itemID int
	// requests github.com/wzshiming/gen-examples/service/item ItemService.Update.item
	var _item *githubComWzshimingGenExamplesServiceItem.Item
	// responses github.com/wzshiming/gen-examples/service/item ItemService.Update.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Parsing item.
	_item, _err_4 = _requestBodyItem(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Update.
	_err_4 = s.Update(_itemID, _item)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetItem Is the route of List
func _operationGetItem(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item ItemService.List.offset
	var _offset_1 int
	// requests github.com/wzshiming/gen-examples/service/item ItemService.List.limit
	var _limit_1 int
	// responses github.com/wzshiming/gen-examples/service/item ItemService.List.items
	var _items []*githubComWzshimingGenExamplesServiceItem.ItemWithID
	// responses github.com/wzshiming/gen-examples/service/item ItemService.List.err
	var _err_4 error

	// Parsing offset.
	_offset_1, _err_4 = _requestQueryOffset_1(w, r)
	if _err_4 != nil {
		return
	}

	// Parsing limit.
	_limit_1, _err_4 = _requestQueryLimit_1(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.List.
	_items, _err_4 = s.List(_offset_1, _limit_1)

	// Response code 200 OK for items.
	if _items != nil {
		var __items []byte
		__items, _err_4 = json.Marshal(_items)
		if _err_4 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_4)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__items)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	var __items []byte
	__items, _err_4 = json.Marshal(_items)
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_4)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__items)

	return
}

// _operationGetItemItemID Is the route of Get
func _operationGetItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item ItemService.Get.item_id
	var _itemID int
	// responses github.com/wzshiming/gen-examples/service/item ItemService.Get.item
	var _item_1 *githubComWzshimingGenExamplesServiceItem.ItemWithID
	// responses github.com/wzshiming/gen-examples/service/item ItemService.Get.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Get.
	_item_1, _err_4 = s.Get(_itemID)

	// Response code 200 OK for item.
	if _item_1 != nil {
		var __item_1 []byte
		__item_1, _err_4 = json.Marshal(_item_1)
		if _err_4 != nil {
			_wrappingErrErrorResp(w, r, 500, _err_4)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__item_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	var __item_1 []byte
	__item_1, _err_4 = json.Marshal(_item_1)
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 500, _err_4)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__item_1)

	return
}

// _operationDeleteItemItemID Is the route of Delete
func _operationDeleteItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item ItemService.Delete.item_id
	var _itemID int
	// responses github.com/wzshiming/gen-examples/service/item ItemService.Delete.err
	var _err_4 error

	// Parsing item_id.
	_itemID, _err_4 = _requestPathItemID(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Delete.
	_err_4 = s.Delete(_itemID)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostItem Is the route of Create
func _operationPostItem(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/item ItemService.Create.item
	var _item *githubComWzshimingGenExamplesServiceItem.Item
	// responses github.com/wzshiming/gen-examples/service/item ItemService.Create.err
	var _err_4 error

	// Parsing item.
	_item, _err_4 = _requestBodyItem(w, r)
	if _err_4 != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Create.
	_err_4 = s.Create(_item)

	// Response code 400 Bad Request for err.
	if _err_4 != nil {
		_wrappingErrErrorResp(w, r, 400, _err_4)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPutMiddMiddID Is the route of Update
func _operationPutMiddMiddID(s *githubComWzshimingGenExamplesServiceMidd.MiddService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Update.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Update.midd_id
	var _middID int
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Update.midd
	var _midd *githubComWzshimingGenExamplesServiceMidd.Midd
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.Update.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Parsing midd.
	_midd, _err = _requestBodyMidd(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddService.Update.
	_err = s.Update(_session_1, _middID, _midd)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetMidd Is the route of List
func _operationGetMidd(s *githubComWzshimingGenExamplesServiceMidd.MiddService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.List.token
	var _token_2 string
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.List.offset
	var _offset_2 int
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.List.limit
	var _limit_2 int
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.List.midds
	var _midds []*githubComWzshimingGenExamplesServiceMidd.MiddWithID
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.List.err
	var _err error

	// Permission middlewares call MiddelToken.
	_token_2, _err = _middlewareTokenMiddelToken(w, r)
	if _err != nil {
		return
	}

	// Parsing offset.
	_offset_2, _err = _requestQueryOffset_2(w, r)
	if _err != nil {
		return
	}

	// Parsing limit.
	_limit_2, _err = _requestQueryLimit_2(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddService.List.
	_midds, _err = s.List(_token_2, _offset_2, _limit_2)

	// Response code 200 OK for midds.
	if _midds != nil {
		var __midds []byte
		__midds, _err = json.Marshal(_midds)
		if _err != nil {
			_wrappingErrErrorResp(w, r, 500, _err)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midds)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	var __midds []byte
	__midds, _err = json.Marshal(_midds)
	if _err != nil {
		_wrappingErrErrorResp(w, r, 500, _err)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midds)

	return
}

// _operationGetMiddMiddID Is the route of Get
func _operationGetMiddMiddID(s *githubComWzshimingGenExamplesServiceMidd.MiddService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Get.token
	var _token_2 string
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Get.midd_id
	var _middID int
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.Get.midd
	var _midd_1 *githubComWzshimingGenExamplesServiceMidd.MiddWithID
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.Get.err
	var _err error

	// Permission middlewares call MiddelToken.
	_token_2, _err = _middlewareTokenMiddelToken(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddService.Get.
	_midd_1, _err = s.Get(_token_2, _middID)

	// Response code 200 OK for midd.
	if _midd_1 != nil {
		var __midd_1 []byte
		__midd_1, _err = json.Marshal(_midd_1)
		if _err != nil {
			_wrappingErrErrorResp(w, r, 500, _err)

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midd_1)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	var __midd_1 []byte
	__midd_1, _err = json.Marshal(_midd_1)
	if _err != nil {
		_wrappingErrErrorResp(w, r, 500, _err)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midd_1)

	return
}

// _operationDeleteMiddMiddID Is the route of Delete
func _operationDeleteMiddMiddID(s *githubComWzshimingGenExamplesServiceMidd.MiddService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Delete.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Delete.midd_id
	var _middID int
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.Delete.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd_id.
	_middID, _err = _requestPathMiddID(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddService.Delete.
	_err = s.Delete(_session_1, _middID)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostMidd Is the route of Create
func _operationPostMidd(s *githubComWzshimingGenExamplesServiceMidd.MiddService, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Create.session
	var _session_1 *githubComWzshimingGenExamplesServiceMidd.Session
	// requests github.com/wzshiming/gen-examples/service/midd MiddService.Create.midd
	var _midd *githubComWzshimingGenExamplesServiceMidd.Midd
	// responses github.com/wzshiming/gen-examples/service/midd MiddService.Create.err
	var _err error

	// Permission middlewares call MiddelSession.
	_session_1, _err = _middlewareSessionMiddelSession(w, r)
	if _err != nil {
		return
	}

	// Parsing midd.
	_midd, _err = _requestBodyMidd(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/midd MiddService.Create.
	_err = s.Create(_session_1, _midd)

	// Response code 400 Bad Request for err.
	if _err != nil {
		_wrappingErrErrorResp(w, r, 400, _err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}
