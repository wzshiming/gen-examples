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
	var _varAuthService githubComWzshimingGenExamplesServiceAuth.AuthService
	RouteAuthService(router, &_varAuthService)

	// FileService Define the method scope
	var _varFileService githubComWzshimingGenExamplesServiceFile.FileService
	RouteFileService(router, &_varFileService)

	// ItemService Define the method scope
	var _varItemService githubComWzshimingGenExamplesServiceItem.ItemService
	RouteItemService(router, &_varItemService)

	// MiddService Define the method scope
	var _varMiddService githubComWzshimingGenExamplesServiceMiddle.MiddService
	RouteMiddService(router, &_varMiddService)

	return router
}

// RouteAuthService is routing for AuthService
func RouteAuthService(router *mux.Router, _varAuthService *githubComWzshimingGenExamplesServiceAuth.AuthService, fs ...mux.MiddlewareFunc) *mux.Router {
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
		_operationGetAuth(_varAuthService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetAuth)

	// Registered routing POST /auth
	var __operationPostAuth http.Handler
	__operationPostAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostAuth(_varAuthService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostAuth)

	// Registered routing PUT /auth/{auth_id}
	var __operationPutAuthAuthID http.Handler
	__operationPutAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutAuthAuthID(_varAuthService, w, r)
	})
	subrouter.Methods("PUT").Path("/{auth_id}").Handler(__operationPutAuthAuthID)

	// Registered routing GET /auth/{auth_id}
	var __operationGetAuthAuthID http.Handler
	__operationGetAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetAuthAuthID(_varAuthService, w, r)
	})
	subrouter.Methods("GET").Path("/{auth_id}").Handler(__operationGetAuthAuthID)

	// Registered routing DELETE /auth/{auth_id}
	var __operationDeleteAuthAuthID http.Handler
	__operationDeleteAuthAuthID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteAuthAuthID(_varAuthService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{auth_id}").Handler(__operationDeleteAuthAuthID)

	return router
}

// RouteFileService is routing for FileService
func RouteFileService(router *mux.Router, _varFileService *githubComWzshimingGenExamplesServiceFile.FileService, fs ...mux.MiddlewareFunc) *mux.Router {
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
		_operationPostFile(_varFileService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostFile)

	// Registered routing GET /file/{filename}
	var __operationGetFileFilename http.Handler
	__operationGetFileFilename = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetFileFilename(_varFileService, w, r)
	})
	subrouter.Methods("GET").Path("/{filename}").Handler(__operationGetFileFilename)

	return router
}

// RouteItemService is routing for ItemService
func RouteItemService(router *mux.Router, _varItemService *githubComWzshimingGenExamplesServiceItem.ItemService, fs ...mux.MiddlewareFunc) *mux.Router {
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
		_operationGetItem(_varItemService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetItem)

	// Registered routing POST /item
	var __operationPostItem http.Handler
	__operationPostItem = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostItem(_varItemService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostItem)

	// Registered routing PUT /item/{item_id}
	var __operationPutItemItemID http.Handler
	__operationPutItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutItemItemID(_varItemService, w, r)
	})
	subrouter.Methods("PUT").Path("/{item_id}").Handler(__operationPutItemItemID)

	// Registered routing GET /item/{item_id}
	var __operationGetItemItemID http.Handler
	__operationGetItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetItemItemID(_varItemService, w, r)
	})
	subrouter.Methods("GET").Path("/{item_id}").Handler(__operationGetItemItemID)

	// Registered routing DELETE /item/{item_id}
	var __operationDeleteItemItemID http.Handler
	__operationDeleteItemItemID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteItemItemID(_varItemService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{item_id}").Handler(__operationDeleteItemItemID)

	return router
}

// RouteMiddService is routing for MiddService
func RouteMiddService(router *mux.Router, _varMiddService *githubComWzshimingGenExamplesServiceMiddle.MiddService, fs ...mux.MiddlewareFunc) *mux.Router {
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
		_operationGetMidd(_varMiddService, w, r)
	})
	subrouter.Methods("GET").Path("").Handler(__operationGetMidd)

	// Registered routing POST /midd
	var __operationPostMidd http.Handler
	__operationPostMidd = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPostMidd(_varMiddService, w, r)
	})
	subrouter.Methods("POST").Path("").Handler(__operationPostMidd)

	// Registered routing PUT /midd/{midd_id}
	var __operationPutMiddMiddID http.Handler
	__operationPutMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationPutMiddMiddID(_varMiddService, w, r)
	})
	subrouter.Methods("PUT").Path("/{midd_id}").Handler(__operationPutMiddMiddID)

	// Registered routing GET /midd/{midd_id}
	var __operationGetMiddMiddID http.Handler
	__operationGetMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetMiddMiddID(_varMiddService, w, r)
	})
	subrouter.Methods("GET").Path("/{midd_id}").Handler(__operationGetMiddMiddID)

	// Registered routing DELETE /midd/{midd_id}
	var __operationDeleteMiddMiddID http.Handler
	__operationDeleteMiddMiddID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationDeleteMiddMiddID(_varMiddService, w, r)
	})
	subrouter.Methods("DELETE").Path("/{midd_id}").Handler(__operationDeleteMiddMiddID)

	return router
}

// _requestBodyAuth Parsing the body for of auth
func _requestBodyAuth(w http.ResponseWriter, r *http.Request) (_varAuth *githubComWzshimingGenExamplesServiceAuth.Auth, err error) {

	var __varAuth []byte
	__varAuth, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__varAuth, &_varAuth)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathAuthID Parsing the path for of auth_id
func _requestPathAuthID(w http.ResponseWriter, r *http.Request) (_varAuthID int, err error) {

	var _raw_varAuthID = mux.Vars(r)["auth_id"]

	if __varAuthID, err := strconv.ParseInt(_raw_varAuthID, 0, 0); err == nil {
		_varAuthID = int(__varAuthID)
	}

	return
}

// _requestBodyFile Parsing the body for of file
func _requestBodyFile(w http.ResponseWriter, r *http.Request) (_varFile io.Reader, err error) {

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

	__varFile := bytes.NewBuffer(nil)
	_, err = io.Copy(__varFile, body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_varFile = __varFile

	return
}

// _requestPathFilename Parsing the path for of filename
func _requestPathFilename(w http.ResponseWriter, r *http.Request) (_varFilename string, err error) {

	var _raw_varFilename = mux.Vars(r)["filename"]
	_varFilename = string(_raw_varFilename)

	return
}

// _requestBodyItem Parsing the body for of item
func _requestBodyItem(w http.ResponseWriter, r *http.Request) (_varItem *githubComWzshimingGenExamplesServiceItem.Item, err error) {

	var __varItem []byte
	__varItem, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__varItem, &_varItem)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathItemID Parsing the path for of item_id
func _requestPathItemID(w http.ResponseWriter, r *http.Request) (_varItemID int, err error) {

	var _raw_varItemID = mux.Vars(r)["item_id"]

	if __varItemID, err := strconv.ParseInt(_raw_varItemID, 0, 0); err == nil {
		_varItemID = int(__varItemID)
	}

	return
}

// _requestQueryLimit Parsing the query for of limit
func _requestQueryLimit(w http.ResponseWriter, r *http.Request) (_varLimit int, err error) {

	var _raw_varLimit = r.URL.Query().Get("limit")

	if __varLimit, err := strconv.ParseInt(_raw_varLimit, 0, 0); err == nil {
		_varLimit = int(__varLimit)
	}

	return
}

// _requestBodyMidd Parsing the body for of midd
func _requestBodyMidd(w http.ResponseWriter, r *http.Request) (_varMidd *githubComWzshimingGenExamplesServiceMiddle.Midd, err error) {

	var __varMidd []byte
	__varMidd, err = ioutil.ReadAll(r.Body)
	if err == nil {
		r.Body.Close()
		err = json.Unmarshal(__varMidd, &_varMidd)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _requestPathMiddID Parsing the path for of midd_id
func _requestPathMiddID(w http.ResponseWriter, r *http.Request) (_varMiddID int, err error) {

	var _raw_varMiddID = mux.Vars(r)["midd_id"]

	if __varMiddID, err := strconv.ParseInt(_raw_varMiddID, 0, 0); err == nil {
		_varMiddID = int(__varMiddID)
	}

	return
}

// _requestQueryOffset Parsing the query for of offset
func _requestQueryOffset(w http.ResponseWriter, r *http.Request) (_varOffset int, err error) {

	var _raw_varOffset = r.URL.Query().Get("offset")

	if __varOffset, err := strconv.ParseInt(_raw_varOffset, 0, 0); err == nil {
		_varOffset = int(__varOffset)
	}

	return
}

// _requestHeaderToken Parsing the header for of token
func _requestHeaderToken(w http.ResponseWriter, r *http.Request) (_varToken string, err error) {

	var _raw_varToken = r.Header.Get("token")
	_varToken = string(_raw_varToken)

	return
}

// _requestHeaderXToken Parsing the header for of x-token
func _requestHeaderXToken(w http.ResponseWriter, r *http.Request) (_varXToken string, err error) {

	var _raw_varXToken = r.Header.Get("x-token")
	_varXToken = string(_raw_varXToken)

	return
}

// _middlewareSessionMiddelSession Is the middleware of MiddelSession
func _middlewareSessionMiddelSession(w http.ResponseWriter, r *http.Request) (_varSession *githubComWzshimingGenExamplesServiceMiddle.Session, err error) {
	var _varXToken string

	// Parsing x-token.
	_varXToken, err = _requestHeaderXToken(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddelSession.
	_varSession, err = githubComWzshimingGenExamplesServiceMiddle.MiddelSession(_varXToken)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _middlewareTokenMiddelToken Is the middleware of MiddelToken
func _middlewareTokenMiddelToken(w http.ResponseWriter, r *http.Request) (_varToken string, err error) {
	var _varSession *githubComWzshimingGenExamplesServiceMiddle.Session

	// Permission middlewares call MiddelSession.
	_varSession, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddelToken.
	_varToken, err = githubComWzshimingGenExamplesServiceMiddle.MiddelToken(_varSession)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _securityAPIKeyVerify Is the security of Verify
func _securityAPIKeyVerify(w http.ResponseWriter, r *http.Request) (_varUserID uint64, err error) {
	var _varToken string

	// Parsing token.
	_varToken, err = _requestHeaderToken(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth Verify.
	_varUserID, err = githubComWzshimingGenExamplesServiceAuth.Verify(_varToken)

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	return
}

// _operationPutAuthAuthID Is the route of Update
func _operationPutAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	var _varUserID uint64
	var _varAuthID int
	var _varAuth *githubComWzshimingGenExamplesServiceAuth.Auth
	var err error

	// Permission verification call Verify.
	_varUserID, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_varAuthID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Parsing auth.
	_varAuth, err = _requestBodyAuth(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Update.
	err = s.Update(_varUserID, _varAuthID, _varAuth)

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
	var _varUserID uint64
	var _varOffset int
	var _varLimit int
	var _varAuths []*githubComWzshimingGenExamplesServiceAuth.AuthWithID
	var err error

	// Permission verification call Verify.
	_varUserID, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing offset.
	_varOffset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_varLimit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.List.
	_varAuths, err = s.List(_varUserID, _varOffset, _varLimit)

	// Response code 200 OK for auths.
	if _varAuths != nil {
		var __varAuths []byte
		__varAuths, err = json.Marshal(_varAuths)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varAuths)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varAuths []byte
	__varAuths, err = json.Marshal(_varAuths)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varAuths)

	return
}

// _operationGetAuthAuthID Is the route of Get
func _operationGetAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	var _varUserID uint64
	var _varAuthID int
	var _varAuth *githubComWzshimingGenExamplesServiceAuth.AuthWithID
	var err error

	// Permission verification call Verify.
	_varUserID, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_varAuthID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Get.
	_varAuth, err = s.Get(_varUserID, _varAuthID)

	// Response code 200 OK for auth.
	if _varAuth != nil {
		var __varAuth []byte
		__varAuth, err = json.Marshal(_varAuth)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varAuth)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varAuth []byte
	__varAuth, err = json.Marshal(_varAuth)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varAuth)

	return
}

// _operationDeleteAuthAuthID Is the route of Delete
func _operationDeleteAuthAuthID(s *githubComWzshimingGenExamplesServiceAuth.AuthService, w http.ResponseWriter, r *http.Request) {
	var _varUserID uint64
	var _varAuthID int
	var err error

	// Permission verification call Verify.
	_varUserID, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth_id.
	_varAuthID, err = _requestPathAuthID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Delete.
	err = s.Delete(_varUserID, _varAuthID)

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
	var _varUserID uint64
	var _varAuth *githubComWzshimingGenExamplesServiceAuth.Auth
	var err error

	// Permission verification call Verify.
	_varUserID, err = _securityAPIKeyVerify(w, r)
	if err != nil {
		return
	}

	// Parsing auth.
	_varAuth, err = _requestBodyAuth(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/auth AuthService.Create.
	err = s.Create(_varUserID, _varAuth)

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
	var _varFile io.Reader
	var _varFilename string
	var err error

	// Parsing file.
	_varFile, err = _requestBodyFile(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Upload.
	_varFilename, err = s.Upload(_varFile)

	// Response code 200 OK for filename.
	if _varFilename != "" {
		var __varFilename []byte
		__varFilename, err = json.Marshal(_varFilename)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varFilename)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varFilename []byte
	__varFilename, err = json.Marshal(_varFilename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varFilename)

	return
}

// _operationGetFileFilename Is the route of Get
func _operationGetFileFilename(s *githubComWzshimingGenExamplesServiceFile.FileService, w http.ResponseWriter, r *http.Request) {
	var _varFilename string
	var _varFile []byte
	var err error

	// Parsing filename.
	_varFilename, err = _requestPathFilename(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/file FileService.Get.
	_varFile, err = s.Get(_varFilename)

	// Response code 200 OK for file.
	if _varFile != nil {
		var __varFile []byte
		__varFile = _varFile

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		w.Write(__varFile)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varFile []byte
	__varFile = _varFile

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(__varFile)

	return
}

// _operationPutItemItemID Is the route of Update
func _operationPutItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	var _varItemID int
	var _varItem *githubComWzshimingGenExamplesServiceItem.Item
	var err error

	// Parsing item_id.
	_varItemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Parsing item.
	_varItem, err = _requestBodyItem(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Update.
	err = s.Update(_varItemID, _varItem)

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
	var _varOffset int
	var _varLimit int
	var _varItems []*githubComWzshimingGenExamplesServiceItem.ItemWithID
	var err error

	// Parsing offset.
	_varOffset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_varLimit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.List.
	_varItems, err = s.List(_varOffset, _varLimit)

	// Response code 200 OK for items.
	if _varItems != nil {
		var __varItems []byte
		__varItems, err = json.Marshal(_varItems)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varItems)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varItems []byte
	__varItems, err = json.Marshal(_varItems)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varItems)

	return
}

// _operationGetItemItemID Is the route of Get
func _operationGetItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	var _varItemID int
	var _varItem *githubComWzshimingGenExamplesServiceItem.ItemWithID
	var err error

	// Parsing item_id.
	_varItemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Get.
	_varItem, err = s.Get(_varItemID)

	// Response code 200 OK for item.
	if _varItem != nil {
		var __varItem []byte
		__varItem, err = json.Marshal(_varItem)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varItem)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varItem []byte
	__varItem, err = json.Marshal(_varItem)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varItem)

	return
}

// _operationDeleteItemItemID Is the route of Delete
func _operationDeleteItemItemID(s *githubComWzshimingGenExamplesServiceItem.ItemService, w http.ResponseWriter, r *http.Request) {
	var _varItemID int
	var err error

	// Parsing item_id.
	_varItemID, err = _requestPathItemID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Delete.
	err = s.Delete(_varItemID)

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
	var _varItem *githubComWzshimingGenExamplesServiceItem.Item
	var err error

	// Parsing item.
	_varItem, err = _requestBodyItem(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/item ItemService.Create.
	err = s.Create(_varItem)

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
	var _varSession *githubComWzshimingGenExamplesServiceMiddle.Session
	var _varMiddID int
	var _varMidd *githubComWzshimingGenExamplesServiceMiddle.Midd
	var err error

	// Permission middlewares call MiddelSession.
	_varSession, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_varMiddID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Parsing midd.
	_varMidd, err = _requestBodyMidd(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Update.
	err = s.Update(_varSession, _varMiddID, _varMidd)

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
	var _varToken string
	var _varOffset int
	var _varLimit int
	var _varMidds []*githubComWzshimingGenExamplesServiceMiddle.MiddWithID
	var err error

	// Permission middlewares call MiddelToken.
	_varToken, err = _middlewareTokenMiddelToken(w, r)
	if err != nil {
		return
	}

	// Parsing offset.
	_varOffset, err = _requestQueryOffset(w, r)
	if err != nil {
		return
	}

	// Parsing limit.
	_varLimit, err = _requestQueryLimit(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.List.
	_varMidds, err = s.List(_varToken, _varOffset, _varLimit)

	// Response code 200 OK for midds.
	if _varMidds != nil {
		var __varMidds []byte
		__varMidds, err = json.Marshal(_varMidds)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varMidds)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varMidds []byte
	__varMidds, err = json.Marshal(_varMidds)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varMidds)

	return
}

// _operationGetMiddMiddID Is the route of Get
func _operationGetMiddMiddID(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {
	var _varToken string
	var _varMiddID int
	var _varMidd *githubComWzshimingGenExamplesServiceMiddle.MiddWithID
	var err error

	// Permission middlewares call MiddelToken.
	_varToken, err = _middlewareTokenMiddelToken(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_varMiddID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Get.
	_varMidd, err = s.Get(_varToken, _varMiddID)

	// Response code 200 OK for midd.
	if _varMidd != nil {
		var __varMidd []byte
		__varMidd, err = json.Marshal(_varMidd)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(__varMidd)
		return
	}

	// Response code 400 Bad Request for err.
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var __varMidd []byte
	__varMidd, err = json.Marshal(_varMidd)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(__varMidd)

	return
}

// _operationDeleteMiddMiddID Is the route of Delete
func _operationDeleteMiddMiddID(s *githubComWzshimingGenExamplesServiceMiddle.MiddService, w http.ResponseWriter, r *http.Request) {
	var _varSession *githubComWzshimingGenExamplesServiceMiddle.Session
	var _varMiddID int
	var err error

	// Permission middlewares call MiddelSession.
	_varSession, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd_id.
	_varMiddID, err = _requestPathMiddID(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Delete.
	err = s.Delete(_varSession, _varMiddID)

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
	var _varSession *githubComWzshimingGenExamplesServiceMiddle.Session
	var _varMidd *githubComWzshimingGenExamplesServiceMiddle.Midd
	var err error

	// Permission middlewares call MiddelSession.
	_varSession, err = _middlewareSessionMiddelSession(w, r)
	if err != nil {
		return
	}

	// Parsing midd.
	_varMidd, err = _requestBodyMidd(w, r)
	if err != nil {
		return
	}

	// Call github.com/wzshiming/gen-examples/service/middle MiddService.Create.
	err = s.Create(_varSession, _varMidd)

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
