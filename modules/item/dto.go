package item

import "mime/multipart"

type CreateItemPayload struct {
	Nama    string           `json:"nama"`
	SellIn  int              `json:"sellIn"`
	Images  []multipart.File `json:"ImagesUrl"`
	Quality int              `json:"quality"`
}
