package cmd

import (
	"fmt"

	"github.com/s19835/interactive-cli-app-go/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all added data",
	Long:  `See all data that you have added to your table in formated view`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
