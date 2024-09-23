package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new studybuddy note",
	Long:  `Create a new note for studybuddy reviewing process`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	wordPrmptContent := promptContent{
		"Please provide a word",
		"What word would you like to make a note of?",
	}

	word := promptGetInput(wordPrmptContent)

	defPromptContent := promptContent{
		"Please provide a definition",
		fmt.Sprintf("What is the definition of %s?", word),
	}

	definition := promptGetInput(defPromptContent)
}
