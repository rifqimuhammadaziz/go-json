package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	// encode to json
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// conversion array to string
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Xenosty") // string
	logJson(1)         // number
	logJson(true)      // bool
	logJson([]string{  // slice (output = array)
		"Rifqi",
		"Xenosty",
		"Xenostheord",
	})
}
