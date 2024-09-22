/*
Copyright © 2024 MARIA NIRANJAN <s19835@sci.pdn.ac.lk>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/s19835/interactive-cli-app-go/data"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize command to run",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
