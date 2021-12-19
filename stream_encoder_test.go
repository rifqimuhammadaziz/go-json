package gojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	// create output file
	writer, _ := os.Create("data/OutputCustomer.json")

	// create encoder to encode json to writer
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Rifqi",
		MiddleName: "Muhammad",
		LastName:   "Aziz",
	}

	// encode to json
	encoder.Encode(customer)
}

func TestStreamEncoder2(t *testing.T) {
	// create output file
	writer, _ := os.Create("data/OutputProduct.json")

	// create encoder to encode json to writer
	encoder := json.NewEncoder(writer)

	product := Product{
		Id:       "ProductID0001",
		Name:     "Apple Macbook Pro M1 Pro",
		ImageURL: "https://www.apple.com",
	}

	// encode to json
	encoder.Encode(product)
}
