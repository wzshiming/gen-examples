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

//
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

//
type Session struct{}

// Verify #security:"apiKey"#
func Verify(_varToken string /* #in:"header"# */) {
	Client = Client.
		SetHeader("token", fmt.Sprint(_varToken))

}

var Client = requests.NewClient().NewRequest()

// Update the Auth #route:"PUT /{auth_id}"#
func (AuthService) Update(_varAuthID int /* #name:"auth_id"# */, _varAuth *Auth) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_varAuthID)).
		SetJSON(_varAuth).
		Put("/auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// List of the Auth #route:"GET /"#
func (AuthService) List(_varOffset int, _varLimit int) (_varAuths []*AuthWithID, _varErr error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_varOffset)).
		SetQuery("limit", fmt.Sprint(_varLimit)).
		Get("/auth")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varAuths)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Auth #route:"GET /{auth_id}"#
func (AuthService) Get(_varAuthID int /* #name:"auth_id"# */) (_varAuth *AuthWithID, _varErr error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_varAuthID)).
		Get("/auth/{auth_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varAuth)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Auth #route:"DELETE /{auth_id}"#
func (AuthService) Delete(_varAuthID int /* #name:"auth_id"# */) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("auth_id", fmt.Sprint(_varAuthID)).
		Delete("/auth/{auth_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// Create a Auth #route:"POST /"#
func (AuthService) Create(_varAuth *Auth) (_varErr error) {
	resp, err := Client.Clone().
		SetJSON(_varAuth).
		Post("/auth")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// Upload a file #route:"POST /"#
func (FileService) Upload(_varFile io.Reader) (_varFilename string, _varErr error) {
	resp, err := Client.Clone().
		SetBody(_varFile).
		Post("/file")
	if err != nil {
		return "", err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varFilename)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return "", err
	}

	return
}

// Get a file #route:"GET /{filename}"#
func (FileService) Get(_varFilename string) (_varFile []byte /* #content:"application/octet-stream"# */, _varErr error) {
	resp, err := Client.Clone().
		SetPath("filename", fmt.Sprint(_varFilename)).
		Get("/file/{filename}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Update the Item #route:"PUT /{item_id}"#
func (ItemService) Update(_varItemID int /* #name:"item_id"# */, _varItem *Item) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_varItemID)).
		SetJSON(_varItem).
		Put("/item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// List of the Item #route:"GET /"#
func (ItemService) List(_varOffset int, _varLimit int) (_varItems []*ItemWithID, _varErr error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_varOffset)).
		SetQuery("limit", fmt.Sprint(_varLimit)).
		Get("/item")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varItems)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Item #route:"GET /{item_id}"#
func (ItemService) Get(_varItemID int /* #name:"item_id"# */) (_varItem *ItemWithID, _varErr error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_varItemID)).
		Get("/item/{item_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varItem)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Item #route:"DELETE /{item_id}"#
func (ItemService) Delete(_varItemID int /* #name:"item_id"# */) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("item_id", fmt.Sprint(_varItemID)).
		Delete("/item/{item_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// Create a Item #route:"POST /"#
func (ItemService) Create(_varItem *Item) (_varErr error) {
	resp, err := Client.Clone().
		SetJSON(_varItem).
		Post("/item")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// Update the Midd #route:"PUT /{midd_id}"#
func (MiddService) Update(_varXToken string /* #in:"header" name:"x-token"# */, _varMiddID int /* #name:"midd_id"# */, _varMidd *Midd) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_varMiddID)).
		SetJSON(_varMidd).
		Put("/midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// List of the Midd #route:"GET /"#
func (MiddService) List(_varXToken string /* #in:"header" name:"x-token"# */, _varOffset int, _varLimit int) (_varMidds []*MiddWithID, _varErr error) {
	resp, err := Client.Clone().
		SetQuery("offset", fmt.Sprint(_varOffset)).
		SetQuery("limit", fmt.Sprint(_varLimit)).
		Get("/midd")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varMidds)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Get the Midd #route:"GET /{midd_id}"#
func (MiddService) Get(_varXToken string /* #in:"header" name:"x-token"# */, _varMiddID int /* #name:"midd_id"# */) (_varMidd *MiddWithID, _varErr error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_varMiddID)).
		Get("/midd/{midd_id}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_varMidd)
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// Delete the Midd #route:"DELETE /{midd_id}"#
func (MiddService) Delete(_varXToken string /* #in:"header" name:"x-token"# */, _varMiddID int /* #name:"midd_id"# */) (_varErr error) {
	resp, err := Client.Clone().
		SetPath("midd_id", fmt.Sprint(_varMiddID)).
		Delete("/midd/{midd_id}")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// Create a Midd #route:"POST /"#
func (MiddService) Create(_varXToken string /* #in:"header" name:"x-token"# */, _varMidd *Midd) (_varErr error) {
	resp, err := Client.Clone().
		SetJSON(_varMidd).
		Post("/midd")
	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_varErr = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}
