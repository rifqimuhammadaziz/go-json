package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonString := `{"id":"P001", "name":"iPhone 11 Pro Max 256GB Gold", "price":500000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P002",
		"name":  "Lenovo Thinkpad",
		"price": 7500000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
