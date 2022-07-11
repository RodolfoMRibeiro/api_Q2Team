package model

import "encoding/json"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) SetCategory(jsonElement string) {
	json.Unmarshal([]byte(jsonElement), c)
}
