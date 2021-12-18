package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("data/Customer.json")
	decoder := json.NewDecoder(reader)

	// like json.unmarshal
	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
