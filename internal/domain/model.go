package domain

import (
	"encoding/json"
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
		return ""
	}
	return string(data)
}

func (item *Item) FromJson(data string) error {
	err := json.Unmarshal([]byte(data), item)
	if err != nil {
		return err
	}
	return nil
}
