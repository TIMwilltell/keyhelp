/**/
package cmd

import (
	"fmt"
	"strings"
	"log"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search and list keyboard shortcuts",
	Long: `Search and list keyboard shortcuts by various criteria:
	- List all shortcuts for a specific application
	- Search for shortcuts by action
	- List all available applications`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("app")
		action, _ := cmd.Flags().GetString("action")
		listApps, _ := cmd.Flags().GetBool("list-apps")

		if appName != "" {
			listShortcutsForApp(appName)
		} else if action != "" {
			searchShortcutsByAction(action)
		} else if listApps {
			listAllApplications()
		} else {
			fmt.Println("Please specify a search type: --app, --action, or --list-apps")
		}
	},
}


func init() {
	rootCmd.AddCommand(searchCmd)

	// Adding flags here
	searchCmd.Flags().StringP("app", "a", "", "List all shortcuts for an application")
	searchCmd.Flags().StringP("action", "c", "", "Search for shortcuts by action")
	searchCmd.Flags().Bool("list-apps", false, "List all available applications")
}

func loadJsonData() {
	// helper function to load the JSON data
}

func listShortcutsForApp(appName string) {
	apps, err := loadShortcutsFromFile("data/test.json")
	if err != nil {
		log.Fatalf("Failed to load shortcuts: %v", err)
	}

	found := false
	for _, app := range apps {
		if app.Name == appName {
			found = true
			fmt.Printf("Shortcuts for %s:\n", appName)
			for _, sc := range app.Shortcuts {
				fmt.Printf("- %s: %s\n", sc.Action, sc.Keys)
			}
		}
	}

	if !found {
		fmt.Println("Application not found.")
	}
}

func searchShortcutsByAction(action string) {
	apps, err := loadShortcutsFromFile("data/test.json")
	if err != nil {
		log.Fatalf("Failed to load shortcuts: %v", err)
	}

	found := false

	for _, app := range apps {
		for _, sc := range app.Shortcuts {
			if strings.EqualFold(sc.Action, action) {
				fmt.Printf("Found in %s: %s - %s\n", app.Name, sc.Action, sc.Keys)
				found = true
			}
		}
	}

	if !found {
		fmt.Println("No shortcuts found for the specified action.")
	}
}

func listAllApplications() {
	fmt.Println("List of Applications 🙂")

	apps, err := loadShortcutsFromFile("data/test.json")
	if err != nil {
		log.Fatalf("Failed to load applications: %v", err)
	}

	if len(apps) == 0 {
		fmt.Println("No applications found.")
		return
	}

	for _, app := range apps {
		fmt.Println("- ", app.Name)
	}
}
