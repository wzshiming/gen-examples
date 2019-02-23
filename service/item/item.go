// Code generated; Do not regenerate the overwrite after editing.

package item

import (
	"errors"
)

// ItemWithID is Item with ID
type ItemWithID struct {
	ID int `json:"item_id,string"`
	Item
}

// ItemService #path:"/item/"#
type ItemService struct {
	datas []*ItemWithID
}

// NewItemService Create a new ItemService
func NewItemService() (*ItemService, error) {
	return &ItemService{}, nil
}

// Create a Item #route:"POST /"#
func (s *ItemService) Create(item *Item) (err error) {
	itemID := len(s.datas) + 1
	data := &ItemWithID{
		ID:   itemID,
		Item: *item,
	}
	s.datas = append(s.datas, data)
	return nil
}

// Update the Item #route:"PUT /{item_id}"#
func (s *ItemService) Update(itemID int /* #name:"item_id"# */, item *Item) (err error) {
	if 0 >= itemID || itemID > len(s.datas) || s.datas[itemID-1] == nil {
		return errors.New("id does not exist")
	}
	v := s.datas[itemID-1]
	v.Item = *item
	return nil
}

// Delete the Item #route:"DELETE /{item_id}"#
func (s *ItemService) Delete(itemID int /* #name:"item_id"# */) (err error) {
	if 0 >= itemID || itemID > len(s.datas) || s.datas[itemID-1] == nil {
		return errors.New("id does not exist")
	}
	s.datas[itemID-1] = nil
	return nil
}

// Get the Item #route:"GET /{item_id}"#
func (s *ItemService) Get(itemID int /* #name:"item_id"# */) (item *ItemWithID, err error) {
	if 0 >= itemID || itemID > len(s.datas) || s.datas[itemID-1] == nil {
		return nil, errors.New("id does not exist")
	}
	return s.datas[itemID-1], nil
}

// List of the Item #route:"GET /"#
func (s *ItemService) List(offset, limit int) (items []*ItemWithID, err error) {
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
				items = append(items, v)
				if lim == limit {
					break
				}
				continue
			}
		}
	}
	return items, nil
}
