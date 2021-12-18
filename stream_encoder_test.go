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
