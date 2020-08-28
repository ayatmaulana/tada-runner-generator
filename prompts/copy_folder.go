package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func CopyFolderPrompt() bool {
	prompt := promptui.Prompt{
		Label:     "Copy app,config,lib,locales folders ?",
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
