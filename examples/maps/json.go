package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataToJSON := make(map[string]interface{})

	dataToJSON["key1"] = "string data"
	dataToJSON["key2"] = 55.55

	jsonString, _ := json.Marshal(dataToJSON)
	var jsonToMap map[string]interface{}
	json.Unmarshal(jsonString, &jsonToMap)

	fmt.Println(jsonToMap)
	str := jsonToMap["key1"].(string)
	num := jsonToMap["key2"].(float64)
	fmt.Println(str)
	fmt.Println(num)

}
