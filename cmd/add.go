/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [application] [action] [shortcut]",
	Short: "add a keyboard shortcut",
	Long: `Add a new keyboard shortcut for an application 
	with the specified action and shortcut key combination.
	Example: keyhelp add "Visual Studio Code" "Save File" "Ctrl+S"`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		application := args[0]
		action := args[1]
		shortcutKey := args [2]

		filePath := "data/test.json" //eventually set this via a config file

		// Load existing shortcuts
		var apps []Application
		if fileData, err := os.ReadFile(filePath); err == nil {
			if err := json.Unmarshal(fileData, &apps); err != nil {
				log.Fatalf("Error reading JSON from file: %v", err)
			}
		} // If file doesn't exist or read fails, start with an empty list

		// Check if the application and shortcut already exist
		foundApp := false
		for i, app := range apps {
			if app.Name == application {
				foundApp = true
				for _, sc := range app.Shortcuts {
					if sc.Action == action && sc.Keys == shortcutKey {
						fmt.Println("Shortcut already exists.")
						return
					}
				}
				// Add shortcut to existing application
				apps[i].Shortcuts = append(apps[i].Shortcuts, Shortcut{Action: action, Keys: shortcutKey})
			}
		}

		if !foundApp {
			// Add new application with the shortcut
			apps = append(apps, Application{
				Name: application,
				Shortcuts: []Shortcut{{Action: action, Keys: shortcutKey}},
			})
		}

		// Serialize the updated apps slice to JSON
		bytes, err := json.MarshalIndent(apps, "", "    ")
		if err != nil {
			log.Fatalf("Error serializing JSON: %v", err)
		}

		// Write the JSON data back to the file
		if err := os.WriteFile(filePath, bytes, 0644); err != nil {
			log.Fatalf("Error writing JSON to file: %v", err)
		}

		fmt.Println("Shortcut added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

type Shortcut struct {
	Name        string `json:"name"`
	Action      string `json:"action"` 
	Keys        string `json:"keys"`
}

type Application struct {
	// ID          int `json:"id"`
	Name        string `json:"name"`
	Shortcuts   []Shortcut `json:"shortcuts"`
} 
