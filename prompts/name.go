package prompts

import (
	"errors"
	"fmt"
	"os"
	"unicode"

	"github.com/manifoldco/promptui"
)

func NamePrompt() string {

	validate := func(input string) error {
		if input == "" {
			return promptui.ErrAbort
		}

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
		if err == promptui.ErrInterrupt || err == promptui.ErrEOF || err == promptui.ErrAbort {
			os.Exit(-1)
		}
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
