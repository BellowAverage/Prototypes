package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func MappingJsonParser(fileContent string) map[string]interface{} {

	var data map[string]interface{}

	// Operating on address directly
	err_parse := json.Unmarshal([]byte(fileContent), &data)

	if err_parse != nil {
		fmt.Printf("Error parsing JSON string - %s", err_parse)
	}

	return data
}

type Nested struct {
	IsIt        bool   `json:"isit"`
	Description string `json:"description"`
}

type exampleJson struct {
	Name    string `json:"name"`
	Numbers []int  `json:"numbers"`
	Nested  Nested `json:"nested"`
}

func main() {

	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)
	fmt.Println("Credit: Chris Wang, 10.12.2024, Go Assignment 3")
	fmt.Println()

	fmt.Println("Example JSON: ")
	fmt.Println("----------------------------")

	fmt.Println(`{
	"name": "example",
	"numbers": [
		1, 2, 3, 4
	],
	"nested": {
		"isit": true,
		"description": "a nested json"
	}
}`)

	fmt.Println()

	fileContent, err_io := os.ReadFile("./GoAssignment3/data/exampleJson.json")

	if err_io != nil {
		log.Fatalf("Error reading file - %s", err_io)
	}

	fmt.Println("Parsing JSON using Mapping: ")
	fmt.Println("----------------------------")

	data := MappingJsonParser(string(fileContent))

	// Print "name" value
	fmt.Printf("Name is %s\n", data["name"].(string))

	// Print "numbers" (handle as float64 array and print each element)
	numbers := data["numbers"].([]interface{})
	for i, num := range numbers {
		fmt.Printf("Number %d is %.0f\n", i, num.(float64)) // Convert to float64
	}

	// Print "nested" description value
	nestedDescription := data["nested"].(map[string]interface{})["description"].(string)
	fmt.Printf("Nested description is %s\n", nestedDescription)

	fmt.Println()
	fmt.Println("Parsing JSON using Structs: ")
	fmt.Println("----------------------------")

	var exampleJson_data exampleJson

	err_exampleJson := json.Unmarshal(fileContent, &exampleJson_data)

	if err_exampleJson != nil {
		log.Fatalf("Error parsing JSON - %s", err_exampleJson)
	}

	// Access and print the fields
	fmt.Printf("Name is %s\n", exampleJson_data.Name)
	fmt.Printf("Numbers are %v\n", exampleJson_data.Numbers)
	fmt.Printf("Nested description is %s\n", exampleJson_data.Nested.Description)

}
