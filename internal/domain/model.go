package domain

import (
	"encoding/json"
	"log"
)

type Item struct {
	Name        string
	Description string
	Price       float32
	Stock       int
	Status      string
}

func (item *Item) ToJson() string {
	data, err := json.Marshal(item)
	if err != nil {
		log.Printf("Error converting to JSON: %v", err)
		return ""
	}
	return string(data)
}
