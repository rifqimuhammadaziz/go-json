package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json tag (output tag)
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P00001",
		Name:     "iPhone 13 Pro Max 256GB",
		ImageURL: "htttps://www.apple.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"Id":"P00001","Name":"iPhone 13 Pro Max 256GB","image_url":"htttps://www.apple.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
