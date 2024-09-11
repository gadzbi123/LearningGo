package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	val := map[string]interface{}{"length": 5, "size": 10, "stupid": true}
	var writer strings.Builder
	jsonEncoder := json.NewEncoder(&writer)
	jsonEncoder.Encode(val)
	stringDecoded := writer.String()
	fmt.Println(stringDecoded)
	result := make([]byte, len(stringDecoded)*3/2)
	base64.StdEncoding.Strict().Encode(result, []byte(stringDecoded))
	fmt.Println(string(result))
}
