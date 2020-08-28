package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func InstallDepPrompt() bool {
	prompt := promptui.Prompt{
		Label:     "do npm install ?",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	if result == "Y" || result == "y" {
		return true
	} else if result == "N" || result == "n" {
		return false
	} else {
		return false
	}
}
