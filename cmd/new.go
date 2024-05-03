/*
Copyright © 2024 Just-Goo
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Just-Goo/study-pal/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new studypal note",
	Long:  `Creates a new studypal note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNote()
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
	// ensure input is not empty
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
		fmt.Printf("prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("input: %s\n", result)
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1 // this keeps the prompt open until the user selects an item with a valid index

	var result string
	var err error

	// keep the prompt open if user has not selected an item
	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 { // if user added a new item to the items list, append the item
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("input: %s\n", result)
	return result
}

func createNote() {
	wordPromptContent := promptContent{
		"Please provide a word",
		"What word would you like to make a note of? : ",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition",
		fmt.Sprintf("What's the definition of %s : ", word),
	}

	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belong to?", word),
	}

	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
