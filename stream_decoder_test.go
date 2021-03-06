package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("data/Customer.json") // get json file
	decoder := json.NewDecoder(reader)         // read data

	// like json.unmarshal
	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamDecoder2(t *testing.T) {
	reader, _ := os.Open("data/Product.json") // get json file
	decoder := json.NewDecoder(reader)        // read data

	// like json.unmarshal
	product := &Product{}
	decoder.Decode(product)

	fmt.Println(product)
}
