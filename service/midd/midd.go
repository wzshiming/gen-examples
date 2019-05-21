// Code generated; Do not regenerate the overwrite after editing.

package midd

import (
	"errors"
)

// MiddWithID is Midd with ID
type MiddWithID struct {
	ID int `json:"midd_id"`
	Midd
}

// MiddService #path:"/midd/"#
type MiddService struct {
	datas []*MiddWithID
}

// NewMiddService Create a new MiddService
func NewMiddService() (*MiddService, error) {
	return &MiddService{}, nil
}

// Create a Midd #route:"POST /"#
func (s *MiddService) Create(
	session *Session, /* #in:"middleware"# */
	midd *Midd,
) (err error) {
	middID := len(s.datas) + 1
	data := &MiddWithID{
		ID:   middID,
		Midd: *midd,
	}
	s.datas = append(s.datas, data)
	return nil
}

// Update the Midd #route:"PUT /{midd_id}"#
func (s *MiddService) Update(
	session *Session, /* #in:"middleware"# */
	middID int /* #name:"midd_id"# */, midd *Midd,
) (err error) {
	if 0 >= middID || middID > len(s.datas) || s.datas[middID-1] == nil {
		return errors.New("id does not exist")
	}
	v := s.datas[middID-1]
	v.Midd = *midd
	return nil
}

// Delete the Midd #route:"DELETE /{midd_id}"#
func (s *MiddService) Delete(
	session *Session, /* #in:"middleware"# */
	middID int, /* #name:"midd_id"# */
) (err error) {
	if 0 >= middID || middID > len(s.datas) || s.datas[middID-1] == nil {
		return errors.New("id does not exist")
	}
	s.datas[middID-1] = nil
	return nil
}

// Get the Midd #route:"GET /{midd_id}"#
func (s *MiddService) Get(
	token string, /* #in:"middleware"# */
	middID int, /* #name:"midd_id"# */
) (midd *MiddWithID, err error) {
	if 0 >= middID || middID > len(s.datas) || s.datas[middID-1] == nil {
		return nil, errors.New("id does not exist")
	}
	return s.datas[middID-1], nil
}

// List of the Midd #route:"GET /"#
func (s *MiddService) List(
	token string, /* #in:"middleware"# */
	offset, limit int,
) (midds []*MiddWithID, err error) {
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
				midds = append(midds, v)
				if lim == limit {
					break
				}
				continue
			}
		}
	}
	return midds, nil
}
