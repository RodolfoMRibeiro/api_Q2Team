package model

import "encoding/json"

type Bookshelf struct {
	Book     int `json:"book"`
	Category int `json:"category"`
}

func (b *Bookshelf) SetBookshelf(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}
