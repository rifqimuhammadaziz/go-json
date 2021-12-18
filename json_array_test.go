package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Rifqi",
		MiddleName: "Muhammad",
		LastName:   "Aziz",
		Age:        22,
		Married:    false,
		Hobbies:    []string{"Play Dota", "Learning new tech", "Code Golang"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rifqi","MiddleName":"Muhammad","LastName":"Aziz","Age":22,"Married":false,"Hobbies":["Play Dota","Learning new tech","Code Golang"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Rifqi",
		Addresses: []Address{
			{
				Street:     "Tegal",
				Country:    "Indonesia",
				PostalCode: "505050",
			},
			{
				Street:     "Semarang",
				Country:    "Indonesia",
				PostalCode: "54321",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rifqi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Tegal","Country":"Indonesia","PostalCode":"505050"},{"Street":"Semarang","Country":"Indonesia","PostalCode":"54321"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}
