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
