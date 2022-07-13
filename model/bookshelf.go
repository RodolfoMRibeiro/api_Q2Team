package model

import "encoding/json"

type Bookshelf struct {
	Book     int `json:"book" gorm:"primary_key; auto_increment:false"`
	Category int `json:"category" gorm:"primary_key; auto_increment:false"`
}

func (b *Bookshelf) SetBookshelf(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}
