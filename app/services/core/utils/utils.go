package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Utility function to load JSON files
func LoadJSONFile(filename string, v interface{}) error {
	// Get the absolute path
	basePath, err := os.Getwd()
	if err != nil {
		return err
	}

	// Construct full file path
	fullPath := fmt.Sprintf("%s/app/services/domain/config/data/%s", basePath, filename)

	// Read file
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}

	// Unmarshal JSON
	return json.Unmarshal(data, v)
}