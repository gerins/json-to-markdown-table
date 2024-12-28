package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type JSONObject map[string]any

// GenerateMarkdown generates markdown tables dynamically for any JSON structure
func GenerateMarkdown(inputJSON string) (string, error) {
	var inputData JSONObject
	if err := json.Unmarshal([]byte(inputJSON), &inputData); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	var markdown bytes.Buffer
	err := generateObjectTable(&markdown, inputData, "Main Structure")
	if err != nil {
		return "", err
	}

	return markdown.String(), nil
}

// Recursively generate a markdown table for a given map
func generateObjectTable(markdown *bytes.Buffer, data JSONObject, title string) error {
	// Write table title
	if title != "" {
		markdown.WriteString(fmt.Sprintf("### %s\n", title))
	}

	// Write table headers
	markdown.WriteString("| Parameter  | Required | Data Type | Description      | Example         |\n")
	markdown.WriteString("|------------|----------|-----------|------------------|-----------------|\n")

	// First pass: Process all scalar (non-object) fields
	for key, value := range data {
		dataType, example := getTypeAndExample(value)
		description := fmt.Sprintf("Field for '%s'", key) // Dynamically generate description
		required := inferRequired(value)                  // Infer if the field is required

		// Write current row
		markdown.WriteString(fmt.Sprintf("| %s        | %s        | %s      | %s              | %v           |\n", key, required, dataType, description, example))
	}

	// Second pass: Process nested objects
	for key, value := range data {
		if nestedMap, isNested := value.(map[string]any); isNested {
			childTitle := fmt.Sprintf("%s Structure", key)
			if err := generateObjectTable(markdown, nestedMap, childTitle); err != nil {
				return err
			}
		}

		if arrayObject, isArrayObject := value.([]any); isArrayObject {
			childTitle := fmt.Sprintf("%s Structure", key)
			if err := generateArrayTable(markdown, arrayObject, childTitle); err != nil {
				return err
			}
		}
	}

	// Add spacing between tables
	markdown.WriteString("\n")
	return nil
}

func generateArrayTable(markdown *bytes.Buffer, data []any, title string) error {
	// Check object with largest field
	var targetObject JSONObject
	for _, object := range data {
		if jsonObject, isJsonObject := object.(map[string]any); isJsonObject {
			if len(jsonObject) > len(targetObject) {
				targetObject = jsonObject
			}
		}
	}

	if targetObject == nil {
		return nil // Skip, not array of object
	}

	return generateObjectTable(markdown, targetObject, title)
}

// Determine the type and example value for a field
func getTypeAndExample(value any) (string, string) {
	dataType := "String" // Default type
	example := fmt.Sprintf("%v", value)

	switch v := value.(type) {
	case int, int64, float64:
		dataType = "Number"
		example = fmt.Sprintf("%v", v)
	case bool:
		dataType = "Boolean"
		example = fmt.Sprintf("%t", v)
	case map[string]any:
		dataType = "Object"
		example = "Refer to sub-structure"
	case []any:
		dataType = "Array Object"
		example = "Refer to list"

		if len(v) == 0 {
			dataType = "Array" // Set just array, because we cant identify the data type
		}

		for _, arrayValue := range value.([]any) {
			switch arrayValue.(type) {
			case string:
				dataType = "Array String"
			case bool:
				dataType = "Array Boolean"
			case int, int64, float64:
				dataType = "Array Number"
			}
		}
	}

	return dataType, example
}

// Infer if a field is required based on its value
func inferRequired(value any) string {
	// A very basic inference: consider fields with non-zero or non-empty values as "M" (Mandatory)
	// and empty/null values as "O" (Optional)
	if value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
		return "O"
	}
	return "M"
}

func main() {
	// Example JSON input
	inputJSON := `
	{
		"quiz": {
			"sport": {
				"q1": {
					"question": "Which one is correct team name in NBA?",
					"options": [
						"New York Bulls",
						"Los Angeles Kings",
						"Golden State Warriros",
						"Huston Rocket"
					],
					"answer": "Huston Rocket"
				}
			},
			"maths": {
				"q1": {
					"question": "5 + 7 = ?",
					"options": [
						"10",
						"11",
						"12",
						"13"
					],
					"answer": "12"
				},
				"q2": {
					"question": "12 - 8 = ?",
					"options": [
						"1",
						"2",
						"3",
						"4"
					],
					"answer": "4"
				}
			}
		}
	}`

	// Generate markdown
	markdown, err := GenerateMarkdown(inputJSON)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print markdown output
	fmt.Println(markdown)
}
