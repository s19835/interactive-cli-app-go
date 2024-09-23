/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note can be anything that you want to review",
	Long:  `Learn anything by adding note and reviewing them regulary, make it stick in your mind.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
