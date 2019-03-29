// Code generated; DO NOT EDIT.
// file ./client/client_gen.go

package client

import (
	json "encoding/json"
	fmt "fmt"
	requests "github.com/wzshiming/requests"
	io "io"
	http "net/http"
)

// Auth is a auth
type Auth struct {
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// AuthService #path:"/auth/"#
type AuthService struct{}

// AuthWithID is Auth with ID
type AuthWithID struct {
	ID   int `json:"auth_id,string"` //
	Auth     //
}

// FileService #path:"/file/"#
type FileService struct {
	BasePath string //
}

// Item is a item
type Item struct {
	Name    string // Name is the name of item
	Message string // Message is the message of item
}

// ItemService #path:"/item/"#
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

// MiddService #path:"/midd/"#
type MiddService struct{}

// MiddWithID is Midd with ID
type MiddWithID struct {
	ID   int `json:"midd_id,string"` //
	Midd     //
}

type Session struct{}

// Verify #security:"apiKey"#
func Verify(_token string /* #in:"header"# */) {
	Client = Client.
		SetHeader("token", fmt.Sprint(_token))

}

var Client = requests.NewClient().NewRequest()

// Update the Auth #route:"PUT /{auth_id}"#
func (AuthService) Update(_authID int /* #name:"auth_id"# */, _auth *Auth) (err error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_authID)).
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

// List of the Auth #route:"GET /"#
func (AuthService) List(_offset int, _limit int) (_auths []*AuthWithID, err error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_offset)).
		SetQuery("limit", fmt.Sprint(_limit)).
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

// Get the Auth #route:"GET /{auth_id}"#
func (AuthService) Get(_authID int /* #name:"auth_id"# */) (_auth *AuthWithID, err error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_authID)).
		Get("auth/{auth_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_auth)
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

// Delete the Auth #route:"DELETE /{auth_id}"#
func (AuthService) Delete(_authID int /* #name:"auth_id"# */) (err error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_authID)).
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

// Create a Auth #route:"POST /"#
func (AuthService) Create(_auth *Auth) (err error) {
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

// Upload a file #route:"POST /"#
func (FileService) Upload(_file io.Reader) (_filename string, err error) {
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

// Get a file #route:"GET /{filename}"#
func (FileService) Get(_filename string) (_file []byte /* #content:"application/octet-stream"# */, err error) {
	resp, err := Client.Clone().
		SetPath("filename", fmt.Sprint(_filename)).
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

// Update the Item #route:"PUT /{item_id}"#
func (ItemService) Update(_itemID int /* #name:"item_id"# */, _item *Item) (err error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_itemID)).
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

// List of the Item #route:"GET /"#
func (ItemService) List(_offset int, _limit int) (_items []*ItemWithID, err error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_offset)).
		SetQuery("limit", fmt.Sprint(_limit)).
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

// Get the Item #route:"GET /{item_id}"#
func (ItemService) Get(_itemID int /* #name:"item_id"# */) (_item *ItemWithID, err error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_itemID)).
		Get("item/{item_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_item)
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

// Delete the Item #route:"DELETE /{item_id}"#
func (ItemService) Delete(_itemID int /* #name:"item_id"# */) (err error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_itemID)).
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

// Create a Item #route:"POST /"#
func (ItemService) Create(_item *Item) (err error) {
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

// Update the Midd #route:"PUT /{midd_id}"#
func (MiddService) Update(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */, _midd *Midd) (err error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_middID)).
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

// List of the Midd #route:"GET /"#
func (MiddService) List(_xToken string /* #in:"header" name:"x-token"# */, _offset int, _limit int) (_midds []*MiddWithID, err error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_offset)).
		SetQuery("limit", fmt.Sprint(_limit)).
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

// Get the Midd #route:"GET /{midd_id}"#
func (MiddService) Get(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (_midd *MiddWithID, err error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_middID)).
		Get("midd/{midd_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_midd)
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

// Delete the Midd #route:"DELETE /{midd_id}"#
func (MiddService) Delete(_xToken string /* #in:"header" name:"x-token"# */, _middID int /* #name:"midd_id"# */) (err error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_middID)).
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

// Create a Midd #route:"POST /"#
func (MiddService) Create(_xToken string /* #in:"header" name:"x-token"# */, _midd *Midd) (err error) {
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
