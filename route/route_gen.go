// Code generated; DO NOT EDIT.
// file ./route/route_gen.go

package route

import (
	bytes "bytes"
	json "encoding/json"
	mux "github.com/gorilla/mux"
	githubComWzshimingGenExamplesServiceAuth "github.com/wzshiming/gen-examples/service/auth"
	githubComWzshimingGenExamplesServiceFile "github.com/wzshiming/gen-examples/service/file"
	githubComWzshimingGenExamplesServiceItem "github.com/wzshiming/gen-examples/service/item"
	githubComWzshimingGenExamplesServiceMiddle "github.com/wzshiming/gen-examples/service/middle"
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	strconv "strconv"
	strings "strings"
)

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

	// ItemService Define the method scope
	var _itemService githubComWzshimingGenExamplesServiceItem.ItemService
	RouteItemService(router, &_itemService)

	// MiddService Define the method scope
	var _middService githubComWzshimingGenExamplesServiceMiddle.MiddService
	RouteMiddService(router, &_middService)

	return router
}

// RouteAuthService is routing for AuthService
func RouteAuthService(router *mux.Router, _authService *githubComWzshimingGenExamplesServiceAuth.AuthService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/auth").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing GET /auth
	var __operationGetAuth http.Handler
	__operationGetAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAuth(_authService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetAuth)

	// Registered routing POST /auth
	var __operationPostAuth http.Handler
	__operationPostAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostAuth(_authService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostAuth)

	// Registered routing PUT /auth/{auth_id}
	var __operationPutAuthAuthID http.Handler
	__operationPutAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutAuthAuthID(_authService, w, r)
	})
	subrouter.Methods("PUT").Path("/{auth_id}").Handler(__operationPutAuthAuthID)

	// Registered routing GET /auth/{auth_id}
	var __operationGetAuthAuthID http.Handler
	__operationGetAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAuthAuthID(_authService, w, r)
	})
	subrouter.Methods("GET").Path("/{auth_id}").Handler(__operationGetAuthAuthID)

	// Registered routing DELETE /auth/{auth_id}
	var __operationDeleteAuthAuthID http.Handler
	__operationDeleteAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteAuthAuthID(_authService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{auth_id}").Handler(__operationDeleteAuthAuthID)

	return router
}

// RouteFileService is routing for FileService
func RouteFileService(router *mux.Router, _fileService *githubComWzshimingGenExamplesServiceFile.FileService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/file").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing POST /file
	var __operationPostFile http.Handler
	__operationPostFile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostFile(_fileService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostFile)

	// Registered routing GET /file/{filename}
	var __operationGetFileFilename http.Handler
	__operationGetFileFilename = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetFileFilename(_fileService, w, r)
	})
	subrouter.Methods("GET").Path("/{filename}").Handler(__operationGetFileFilename)

	return router
}

// RouteItemService is routing for ItemService
func RouteItemService(router *mux.Router, _itemService *githubComWzshimingGenExamplesServiceItem.ItemService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/item").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing GET /item
	var __operationGetItem http.Handler
	__operationGetItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetItem(_itemService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetItem)

	// Registered routing POST /item
	var __operationPostItem http.Handler
	__operationPostItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostItem(_itemService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostItem)

	// Registered routing PUT /item/{item_id}
	var __operationPutItemItemID http.Handler
	__operationPutItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutItemItemID(_itemService, w, r)
	})
	subrouter.Methods("PUT").Path("/{item_id}").Handler(__operationPutItemItemID)

	// Registered routing GET /item/{item_id}
	var __operationGetItemItemID http.Handler
	__operationGetItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetItemItemID(_itemService, w, r)
	})
	subrouter.Methods("GET").Path("/{item_id}").Handler(__operationGetItemItemID)

	// Registered routing DELETE /item/{item_id}
	var __operationDeleteItemItemID http.Handler
	__operationDeleteItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteItemItemID(_itemService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{item_id}").Handler(__operationDeleteItemItemID)

	return router
}

// RouteMiddService is routing for MiddService
func RouteMiddService(router *mux.Router, _middService *githubComWzshimingGenExamplesServiceMiddle.MiddService, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	subrouter := router.PathPrefix("/midd").Subrouter()
	if len(fs) != 0 {
		subrouter.Use(fs...)
	}

	// Registered routing GET /midd
	var __operationGetMidd http.Handler
	__operationGetMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetMidd(_middService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetMidd)

	// Registered routing POST /midd
	var __operationPostMidd http.Handler
	__operationPostMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostMidd(_middService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostMidd)

	// Registered routing PUT /midd/{midd_id}
	var __operationPutMiddMiddID http.Handler
	__operationPutMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutMiddMiddID(_middService, w, r)
	})
	subrouter.Methods("PUT").Path("/{midd_id}").Handler(__operationPutMiddMiddID)

	// Registered routing GET /midd/{midd_id}
	var __operationGetMiddMiddID http.Handler
	__operationGetMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetMiddMiddID(_middService, w, r)
	})
	subrouter.Methods("GET").Path("/{midd_id}").Handler(__operationGetMiddMiddID)

	// Registered routing DELETE /midd/{midd_id}
	var __operationDeleteMiddMiddID http.Handler
	__operationDeleteMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteMiddMiddID(_middService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{midd_id}").Handler(__operationDeleteMiddMiddID)

	return router
}

// _requestBodyAuth Parsing the body for of auth
func _requestBodyAuth(w http.ResponseWriter, r *http.Request) (_auth *githubComWzshimingGenExamplesServiceAuth.Auth, err error) {

	var __auth []byte
	__auth, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__auth, &_auth)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathAuthID Parsing the path for of auth_id
func _requestPathAuthID(w http.ResponseWriter, r *http.Request) (_authID int, err error) {

	var _raw_authID = mux.Vars(r)["auth_id"]

	if __authID, err := strconv.ParseInt(_raw_authID, 0, 0); err == nil {
		_authID = int(__authID)
	}

	return
}

// _requestBodyFile Parsing the body for of file
func _requestBodyFile(w http.ResponseWriter, r *http.Request) (_file io.Reader, err error) {

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
		http.Error(w, err.Error(), 400)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_file = __file

	return
}

// _requestPathFilename Parsing the path for of filename
func _requestPathFilename(w http.ResponseWriter, r *http.Request) (_filename string, err error) {

	var _raw_filename = mux.Vars(r)["filename"]
	_filename = string(_raw_filename)

	return
}

// _requestBodyItem Parsing the body for of item
func _requestBodyItem(w http.ResponseWriter, r *http.Request) (_item *githubComWzshimingGenExamplesServiceItem.Item, err error) {

	var __item []byte
	__item, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__item, &_item)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathItemID Parsing the path for of item_id
func _requestPathItemID(w http.ResponseWriter, r *http.Request) (_itemID int, err error) {

	var _raw_itemID = mux.Vars(r)["item_id"]

	if __itemID, err := strconv.ParseInt(_raw_itemID, 0, 0); err == nil {
		_itemID = int(__itemID)
	}

	return
}

// _requestQueryLimit Parsing the query for of limit
func _requestQueryLimit(w http.ResponseWriter, r *http.Request) (_limit int, err error) {

	var _raw_limit = r.URL.Query().Get("limit")

	if __limit, err := strconv.ParseInt(_raw_limit, 0, 0); err == nil {
		_limit = int(__limit)
	}

	return
}

// _requestBodyMidd Parsing the body for of midd
func _requestBodyMidd(w http.ResponseWriter, r *http.Request) (_midd *githubComWzshimingGenExamplesServiceMiddle.Midd, err error) {

	var __midd []byte
	__midd, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__midd, &_midd)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathMiddID Parsing the path for of midd_id
func _requestPathMiddID(w http.ResponseWriter, r *http.Request) (_middID int, err error) {

	var _raw_middID = mux.Vars(r)["midd_id"]

	if __middID, err := strconv.ParseInt(_raw_middID, 0, 0); err == nil {
		_middID = int(__middID)
	}

	return
}

// _requestQueryOffset Parsing the query for of offset
func _requestQueryOffset(w http.ResponseWriter, r *http.Request) (_offset int, err error) {

	var _raw_offset = r.URL.Query().Get("offset")

	if __offset, err := strconv.ParseInt(_raw_offset, 0, 0); err == nil {
		_offset = int(__offset)
	}

	return
}

// _requestHeaderToken Parsing the header for of token
func _requestHeaderToken(w http.ResponseWriter, r *http.Request) (_token string, err error) {

	var _raw_token = r.Header.Get("token")
	_token = string(_raw_token)

	return
}

// _requestHeaderXToken Parsing the header for of x-token
func _requestHeaderXToken(w http.ResponseWriter, r *http.Request) (_xToken string, err error) {

	var _raw_xToken = r.Header.Get("x-token")
	_xToken = string(_raw_xToken)

	return
}

// _middlewareSessionMiddelSession Is the middleware of MiddelSession
func _middlewareSessionMiddelSession(w http.ResponseWriter, r *http.Request) (_session *githubComWzshimingGenExamplesServiceMiddle.Session, err error) {
	var _xToken string

	// Parsing x-token.
	_xToken, err = _requestHeaderXToken(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddelSession.
	_session, err = githubComWzshimingGenExamplesServiceMiddle.MiddelSession(_xToken)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _middlewareTokenMiddelToken Is the middleware of MiddelToken
func _middlewareTokenMiddelToken(w http.ResponseWriter, r *http.Request) (_token_1 string, err error) {
	var _session_1 *githubComWzshimingGenExamplesServiceMiddle.Session

	// Permission middlewares call MiddelSession.
	_session_1, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddelToken.
	_token_1, err = githubComWzshimingGenExamplesServiceMiddle.MiddelToken(_session_1)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _securityAPIKeyVerify Is the security of Verify
func _securityAPIKeyVerify(w http.ResponseWriter, r *http.Request) (_userID uint64, err error) {
	var _token string

	// Parsing token.
	_token, err = _requestHeaderToken(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Verify.
	_userID, err = githubComWzshimingGenExamplesServiceAuth.Verify(_token)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _operationPutAuthAuthID Is the route of Update
func _operationPutAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _userID_1 uint64
	var _authID int
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth

	// Permission verification call Verify.
	_userID_1, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_authID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Parsing auth.
	_auth, err = _requestBodyAuth(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Update.
	err = s.Update(_userID_1, _authID, _auth)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetAuth Is the route of List
func _operationGetAuth(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _userID_1 uint64
	var _offset int
	var _limit int
	var _auths []*githubComWzshimingGenExamplesServiceAuth.AuthWithID

	// Permission verification call Verify.
	_userID_1, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing offset.
	_offset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_limit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.List.
	_auths, err = s.List(_userID_1, _offset, _limit)

	// Response code 200 OK for auths.
	if _auths != nil {
		var __auths []byte
		__auths, err = json.Marshal(_auths)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auths)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __auths []byte
	__auths, err = json.Marshal(_auths)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auths)

	return
}

// _operationGetAuthAuthID Is the route of Get
func _operationGetAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _userID_1 uint64
	var _authID int
	var _auth_1 *githubComWzshimingGenExamplesServiceAuth.AuthWithID

	// Permission verification call Verify.
	_userID_1, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_authID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Get.
	_auth_1, err = s.Get(_userID_1, _authID)

	// Response code 200 OK for auth.
	if _auth_1 != nil {
		var __auth_1 []byte
		__auth_1, err = json.Marshal(_auth_1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__auth_1)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __auth_1 []byte
	__auth_1, err = json.Marshal(_auth_1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__auth_1)

	return
}

// _operationDeleteAuthAuthID Is the route of Delete
func _operationDeleteAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _userID_1 uint64
	var _authID int

	// Permission verification call Verify.
	_userID_1, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_authID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Delete.
	err = s.Delete(_userID_1, _authID)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostAuth Is the route of Create
func _operationPostAuth(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _userID_1 uint64
	var _auth *githubComWzshimingGenExamplesServiceAuth.Auth

	// Permission verification call Verify.
	_userID_1, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth.
	_auth, err = _requestBodyAuth(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Create.
	err = s.Create(_userID_1, _auth)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostFile Is the route of Upload
func _operationPostFile(s *githubComWzshimingGenExamplesServiceFile.FileService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _file io.Reader
	var _filename_1 string

	// Parsing file.
	_file, err = _requestBodyFile(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Upload.
	_filename_1, err = s.Upload(_file)

	// Response code 200 OK for filename.
	if _filename_1 != "" {
		var __filename_1 []byte
		__filename_1, err = json.Marshal(_filename_1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__filename_1)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __filename_1 []byte
	__filename_1, err = json.Marshal(_filename_1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__filename_1)

	return
}

// _operationGetFileFilename Is the route of Get
func _operationGetFileFilename(s *githubComWzshimingGenExamplesServiceFile.FileService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _filename string
	var _file_1 []byte

	// Parsing filename.
	_filename, err = _requestPathFilename(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Get.
	_file_1, err = s.Get(_filename)

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
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __file_1 []byte
	__file_1 = _file_1

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(__file_1)

	return
}

// _operationPutItemItemID Is the route of Update
func _operationPutItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _itemID int
	var _item *githubComWzshimingGenExamplesServiceItem.Item

	// Parsing item_id.
	_itemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Parsing item.
	_item, err = _requestBodyItem(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Update.
	err = s.Update(_itemID, _item)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetItem Is the route of List
func _operationGetItem(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _offset int
	var _limit int
	var _items []*githubComWzshimingGenExamplesServiceItem.ItemWithID

	// Parsing offset.
	_offset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_limit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.List.
	_items, err = s.List(_offset, _limit)

	// Response code 200 OK for items.
	if _items != nil {
		var __items []byte
		__items, err = json.Marshal(_items)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__items)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __items []byte
	__items, err = json.Marshal(_items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__items)

	return
}

// _operationGetItemItemID Is the route of Get
func _operationGetItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _itemID int
	var _item_1 *githubComWzshimingGenExamplesServiceItem.ItemWithID

	// Parsing item_id.
	_itemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Get.
	_item_1, err = s.Get(_itemID)

	// Response code 200 OK for item.
	if _item_1 != nil {
		var __item_1 []byte
		__item_1, err = json.Marshal(_item_1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__item_1)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __item_1 []byte
	__item_1, err = json.Marshal(_item_1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__item_1)

	return
}

// _operationDeleteItemItemID Is the route of Delete
func _operationDeleteItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _itemID int

	// Parsing item_id.
	_itemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Delete.
	err = s.Delete(_itemID)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostItem Is the route of Create
func _operationPostItem(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _item *githubComWzshimingGenExamplesServiceItem.Item

	// Parsing item.
	_item, err = _requestBodyItem(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Create.
	err = s.Create(_item)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPutMiddMiddID Is the route of Update
func _operationPutMiddMiddID(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _session_1 *githubComWzshimingGenExamplesServiceMiddle.Session
	var _middID int
	var _midd *githubComWzshimingGenExamplesServiceMiddle.Midd

	// Permission middlewares call MiddelSession.
	_session_1, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_middID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Parsing midd.
	_midd, err = _requestBodyMidd(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Update.
	err = s.Update(_session_1, _middID, _midd)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationGetMidd Is the route of List
func _operationGetMidd(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _token_2 string
	var _offset int
	var _limit int
	var _midds []*githubComWzshimingGenExamplesServiceMiddle.MiddWithID

	// Permission middlewares call MiddelToken.
	_token_2, err = _middlewareTokenMiddelToken(w, r)
	if err != nil {
		return
	}

	// Parsing offset.
	_offset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_limit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.List.
	_midds, err = s.List(_token_2, _offset, _limit)

	// Response code 200 OK for midds.
	if _midds != nil {
		var __midds []byte
		__midds, err = json.Marshal(_midds)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midds)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __midds []byte
	__midds, err = json.Marshal(_midds)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midds)

	return
}

// _operationGetMiddMiddID Is the route of Get
func _operationGetMiddMiddID(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _token_2 string
	var _middID int
	var _midd_1 *githubComWzshimingGenExamplesServiceMiddle.MiddWithID

	// Permission middlewares call MiddelToken.
	_token_2, err = _middlewareTokenMiddelToken(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_middID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Get.
	_midd_1, err = s.Get(_token_2, _middID)

	// Response code 200 OK for midd.
	if _midd_1 != nil {
		var __midd_1 []byte
		__midd_1, err = json.Marshal(_midd_1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__midd_1)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __midd_1 []byte
	__midd_1, err = json.Marshal(_midd_1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__midd_1)

	return
}

// _operationDeleteMiddMiddID Is the route of Delete
func _operationDeleteMiddMiddID(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _session_1 *githubComWzshimingGenExamplesServiceMiddle.Session
	var _middID int

	// Permission middlewares call MiddelSession.
	_session_1, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_middID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Delete.
	err = s.Delete(_session_1, _middID)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}

// _operationPostMidd Is the route of Create
func _operationPostMidd(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {

	var err error
	var _session_1 *githubComWzshimingGenExamplesServiceMiddle.Session
	var _midd *githubComWzshimingGenExamplesServiceMiddle.Midd

	// Permission middlewares call MiddelSession.
	_session_1, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd.
	_midd, err = _requestBodyMidd(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Create.
	err = s.Create(_session_1, _midd)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("null"))

	return
}
