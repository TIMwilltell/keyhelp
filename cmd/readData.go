/**/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Function to load shortcuts from the JSON file
func loadShortcutsFromFile(filePath string) ([]Application, error) {
	var apps []Application

	// Checking if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Returning an empty list if the file does not exist
		return apps, nil
	}

	// Reading the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshalling the JSON data into the apps slice
	if err := json.Unmarshal(data, &apps); err != nil {
		return nil, err
	}

	return apps, nil
}
