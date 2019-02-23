// Code generated; Do not regenerate the overwrite after editing.

package auth

import (
	"errors"
)

// AuthWithID is Auth with ID
type AuthWithID struct {
	ID int `json:"auth_id,string"`
	Auth
}

// AuthService #path:"/auth/"#
type AuthService struct {
	datas []*AuthWithID
}

// NewAuthService Create a new AuthService
func NewAuthService() (*AuthService, error) {
	return &AuthService{}, nil
}

// Create a Auth #route:"POST /"#
func (s *AuthService) Create(userID uint64 /* #in:"security"# */, auth *Auth) (err error) {
	authID := len(s.datas) + 1
	data := &AuthWithID{
		ID:   authID,
		Auth: *auth,
	}
	s.datas = append(s.datas, data)
	return nil
}

// Update the Auth #route:"PUT /{auth_id}"#
func (s *AuthService) Update(userID uint64 /* #in:"security"# */, authID int /* #name:"auth_id"# */, auth *Auth) (err error) {
	if 0 >= authID || authID > len(s.datas) || s.datas[authID-1] == nil {
		return errors.New("id does not exist")
	}
	v := s.datas[authID-1]
	v.Auth = *auth
	return nil
}

// Delete the Auth #route:"DELETE /{auth_id}"#
func (s *AuthService) Delete(userID uint64 /* #in:"security"# */, authID int /* #name:"auth_id"# */) (err error) {
	if 0 >= authID || authID > len(s.datas) || s.datas[authID-1] == nil {
		return errors.New("id does not exist")
	}
	s.datas[authID-1] = nil
	return nil
}

// Get the Auth #route:"GET /{auth_id}"#
func (s *AuthService) Get(userID uint64 /* #in:"security"# */, authID int /* #name:"auth_id"# */) (auth *AuthWithID, err error) {
	if 0 >= authID || authID > len(s.datas) || s.datas[authID-1] == nil {
		return nil, errors.New("id does not exist")
	}
	return s.datas[authID-1], nil
}

// List of the Auth #route:"GET /"#
func (s *AuthService) List(userID uint64 /* #in:"security"# */, offset, limit int) (auths []*AuthWithID, err error) {
	off := 0
	lim := 0
	for _, v := range s.datas {
		if v != nil {
			if offset != 0 && off != offset {
				off++
				continue
			}
			if limit == 0 || lim != limit {
				lim++
				auths = append(auths, v)
				if lim == limit {
					break
				}
				continue
			}
		}
	}
	return auths, nil
}
