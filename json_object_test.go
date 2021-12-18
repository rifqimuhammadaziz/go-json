package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Rifqi",
		MiddleName: "Muhammad",
		LastName:   "Aziz",
		Age:        22,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
