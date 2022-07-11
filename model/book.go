package model

import "encoding/json"

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (b *Book) SetBook(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), b)
}

func (b *Book) FindById(id int) {
	DB.Where("id = ?", id).First(b)
}