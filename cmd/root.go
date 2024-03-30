package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"threads/pkg/api"
)

var rootCmd = &cobra.Command{
	Use:   "titlecli",
	Short: "TitleCLI sends a title to an API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter the title:")
		var title string
		fmt.Scanln(&title)
		title = strings.TrimSpace(title)
		err := api.SendTitle(title)
		if err != nil {
			fmt.Println("Error sending title:", err)
			return
		}
		fmt.Println("Title sent successfully.")
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
