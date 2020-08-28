package prompts

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/manifoldco/promptui"
)

func NamePrompt() string {

	validate := func(input string) error {
		for _, letter := range input {
			if unicode.IsSymbol(letter) {
				return errors.New("Invalid name format")
			}
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Runner Name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
