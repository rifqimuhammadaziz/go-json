package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	// json data
	jsonString := `{"FirstName":"Rifqi","MiddleName":"Muhammad","LastName":"Aziz","Age":22,"Married":false}`

	// convert json data to byte array
	jsonBytes := []byte(jsonString)

	// create object from struct
	customer := &Customer{}

	// read data json to go type data
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
}
