// Code generated; DO NOT EDIT.
// file ./client/client_gen.go

package client

import (
	json "encoding/json"
	fmt "fmt"
	io "io"
	http "net/http"
	strconv "strconv"

	requests "github.com/wzshiming/requests"
)

// Auth is a auth
type Auth struct {
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// AuthService
// #path:"/auth/"#
type AuthService struct{}

// AuthWithID is Auth with ID
type AuthWithID struct {
	ID   int `json:"auth_id,string"` //
	Auth     //
}

// FileService
// #path:"/file/"#
type FileService struct {
	BasePath string //
}

// Group
// #path:"/group"#
type Group struct {
	Item ItemService // #path:"/item"#
	File FileService // #path:"/file"#
	Auth AuthService // #path:"/auth"#
	Midd MiddService // #path:"/midd"#
}

// Item is a item
type Item struct {
	Name    string // Name is the name of item
	Message string // Message is the message of item
}

// ItemService
// #path:"/item/"#
type ItemService struct{}

// ItemWithID is Item with ID
type ItemWithID struct {
	ID   int `json:"item_id,string"` //
	Item     //
}

// Midd is a midd
type Midd struct {
	MiddID  int    `json:"midd_id"` // MiddID is the Midd ID
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// MiddService
// #path:"/midd/"#
type MiddService struct{}

// MiddWithID is Midd with ID
type MiddWithID struct {
	ID   int `json:"midd_id,string"` //
	Midd     //
}

type Session struct{}

// VerifyApiKey
// #security:"apiKey"#
func VerifyApiKey(_token string /* #in:"header"# #name:"token"# */) {
	var __token string
	__token = string(_token)

	Client = Client.
		SetHeader("token", __token)
}

// VerifyBasic
// #security:"basic"#
func VerifyBasic(username string, password string) {

	Client = Client.
		SetBasicAuth(username, password)

}

var Client = requests.NewClient().NewRequest()

// Update the Auth
// #route:"PUT /{auth_id}"#
func (AuthService) Update(_authID int /* #name:"auth_id"# */, _auth *Auth /* #name:"auth"# */) (err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		SetJSON(_auth).
		Put("auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Auth
// #route:"GET /"#
func (AuthService) List(_offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_auths []*AuthWithID /* #name:"auths"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("auth")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_auths)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Auth
// #route:"GET /{auth_id}"#
func (AuthService) Get(_authID int /* #name:"auth_id"# */) (_auth_1 *AuthWithID /* #name:"auth"# */, err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		Get("auth/{auth_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_auth_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Auth
// #route:"DELETE /{auth_id}"#
func (AuthService) Delete(_authID int /* #name:"auth_id"# */) (err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		Delete("auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Auth
// #route:"POST /"#
func (AuthService) Create(_auth *Auth /* #name:"auth"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_auth).
		Post("auth")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Upload a file
// #route:"POST /"#
func (FileService) Upload(_file io.Reader /* #name:"file"# */) (_filename string /* #name:"filename"# */, err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetBody(_file).
		Post("file")
	if err != nil {
		return "", err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_filename)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return "", err
	}

	return
}

// Get a file
// #route:"GET /{filename}"#
func (FileService) Get(_filename_1 string /* #name:"filename"# */) (_file_1 []byte /* #content:"application/octet-stream"# #name:"file"# */, err error /* #name:"err"# */) {
	var __filename_1 string
	__filename_1 = string(_filename_1)

	resp, err := Client.Clone().
		SetPath("filename", __filename_1).
		Get("file/{filename}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Update the Item
// #route:"PUT /{item_id}"#
func (Group) ItemUpdate(_itemID int /* #name:"item_id"# */, _item *Item /* #name:"item"# */) (err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		SetJSON(_item).
		Put("group/item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Item
// #route:"GET /"#
func (Group) ItemList(_offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_items []*ItemWithID /* #name:"items"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("group/item")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_items)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Item
// #route:"GET /{item_id}"#
func (Group) ItemGet(_itemID int /* #name:"item_id"# */) (_item_1 *ItemWithID /* #name:"item"# */, err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		Get("group/item/{item_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_item_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Item
// #route:"DELETE /{item_id}"#
func (Group) ItemDelete(_itemID int /* #name:"item_id"# */) (err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		Delete("group/item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Item
// #route:"POST /"#
func (Group) ItemCreate(_item *Item /* #name:"item"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_item).
		Post("group/item")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Upload a file
// #route:"POST /"#
func (Group) FileUpload(_file io.Reader /* #name:"file"# */) (_filename string /* #name:"filename"# */, err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetBody(_file).
		Post("group/file")
	if err != nil {
		return "", err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_filename)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return "", err
	}

	return
}

// Get a file
// #route:"GET /{filename}"#
func (Group) FileGet(_filename_1 string /* #name:"filename"# */) (_file_1 []byte /* #content:"application/octet-stream"# #name:"file"# */, err error /* #name:"err"# */) {
	var __filename_1 string
	__filename_1 = string(_filename_1)

	resp, err := Client.Clone().
		SetPath("filename", __filename_1).
		Get("group/file/{filename}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Update the Auth
// #route:"PUT /{auth_id}"#
func (Group) AuthUpdate(_authID int /* #name:"auth_id"# */, _auth *Auth /* #name:"auth"# */) (err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		SetJSON(_auth).
		Put("group/auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Auth
// #route:"GET /"#
func (Group) AuthList(_offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_auths []*AuthWithID /* #name:"auths"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("group/auth")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_auths)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Auth
// #route:"GET /{auth_id}"#
func (Group) AuthGet(_authID int /* #name:"auth_id"# */) (_auth_1 *AuthWithID /* #name:"auth"# */, err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		Get("group/auth/{auth_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_auth_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Auth
// #route:"DELETE /{auth_id}"#
func (Group) AuthDelete(_authID int /* #name:"auth_id"# */) (err error /* #name:"err"# */) {
	var __authID string

	__authID = strconv.FormatInt(int64(_authID), 10)

	resp, err := Client.Clone().
		SetPath("auth_id", __authID).
		Delete("group/auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Auth
// #route:"POST /"#
func (Group) AuthCreate(_auth *Auth /* #name:"auth"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_auth).
		Post("group/auth")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Update the Midd
// #route:"PUT /{midd_id}"#
func (Group) MiddUpdate(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */, _midd *Midd /* #name:"midd"# */) (err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		SetJSON(_midd).
		Put("group/midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Midd
// #route:"GET /"#
func (Group) MiddList(_xToken string /* #in:"header" name:"x-token"# */, _offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_midds []*MiddWithID /* #name:"midds"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("group/midd")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_midds)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Midd
// #route:"GET /{midd_id}"#
func (Group) MiddGet(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (_midd_1 *MiddWithID /* #name:"midd"# */, err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		Get("group/midd/{midd_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_midd_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Midd
// #route:"DELETE /{midd_id}"#
func (Group) MiddDelete(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		Delete("group/midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Midd
// #route:"POST /"#
func (Group) MiddCreate(_xToken string /* #in:"header" name:"x-token"# */, _midd *Midd /* #name:"midd"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_midd).
		Post("group/midd")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Update the Item
// #route:"PUT /{item_id}"#
func (ItemService) Update(_itemID int /* #name:"item_id"# */, _item *Item /* #name:"item"# */) (err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		SetJSON(_item).
		Put("item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Item
// #route:"GET /"#
func (ItemService) List(_offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_items []*ItemWithID /* #name:"items"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("item")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_items)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Item
// #route:"GET /{item_id}"#
func (ItemService) Get(_itemID int /* #name:"item_id"# */) (_item_1 *ItemWithID /* #name:"item"# */, err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		Get("item/{item_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_item_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Item
// #route:"DELETE /{item_id}"#
func (ItemService) Delete(_itemID int /* #name:"item_id"# */) (err error /* #name:"err"# */) {
	var __itemID string

	__itemID = strconv.FormatInt(int64(_itemID), 10)

	resp, err := Client.Clone().
		SetPath("item_id", __itemID).
		Delete("item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Item
// #route:"POST /"#
func (ItemService) Create(_item *Item /* #name:"item"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_item).
		Post("item")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Update the Midd
// #route:"PUT /{midd_id}"#
func (MiddService) Update(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */, _midd *Midd /* #name:"midd"# */) (err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		SetJSON(_midd).
		Put("midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// List of the Midd
// #route:"GET /"#
func (MiddService) List(_xToken string /* #in:"header" name:"x-token"# */, _offset int /* #name:"offset"# */, _limit int /* #name:"limit"# */) (_midds []*MiddWithID /* #name:"midds"# */, err error /* #name:"err"# */) {
	var __offset string
	var __limit string

	__offset = strconv.FormatInt(int64(_offset), 10)

	__limit = strconv.FormatInt(int64(_limit), 10)

	resp, err := Client.Clone().
		SetQuery("offset", __offset).
		SetQuery("limit", __limit).
		Get("midd")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_midds)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Midd
// #route:"GET /{midd_id}"#
func (MiddService) Get(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (_midd_1 *MiddWithID /* #name:"midd"# */, err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		Get("midd/{midd_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_midd_1)
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Midd
// #route:"DELETE /{midd_id}"#
func (MiddService) Delete(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (err error /* #name:"err"# */) {
	var __middID string

	__middID = strconv.FormatInt(int64(_middID), 10)

	resp, err := Client.Clone().
		SetPath("midd_id", __middID).
		Delete("midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}

// Create a Midd
// #route:"POST /"#
func (MiddService) Create(_xToken string /* #in:"header" name:"x-token"# */, _midd *Midd /* #name:"midd"# */) (err error /* #name:"err"# */) {

	resp, err := Client.Clone().
		SetJSON(_midd).
		Post("midd")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		err = fmt.Errorf(string(resp.Body()))
	default:
		if code >= 400 {
			err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
		}
	}

	if err != nil {
		return err
	}

	return
}
