package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ayatmaulana/tada-runner-generator/lib"
	"github.com/ayatmaulana/tada-runner-generator/prompts"
	struts "github.com/ayatmaulana/tada-runner-generator/structs"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"

	"github.com/gobuffalo/packr"
)

var (
	box        packr.Box
	currentDir string

	installDep bool
	copyDir    bool
)

func main() {

	var cmdAdd = &cobra.Command{
		Use:              "add [runner name]",
		Short:            "Add new runner",
		Long:             "",
		Args:             cobra.MinimumNArgs(1),
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "" {
				fmt.Println("Invalid Arguments")
				os.Exit(-1)
			}
			exec(args[0])
		},
	}
	cmdAdd.PersistentFlags().BoolVarP(&installDep, "install-dep", "i", true, "Install Depedencies")
	cmdAdd.PersistentFlags().BoolVarP(&copyDir, "copy-folder", "c", true, "Copy app/config/lib/locales folder ?")

	var cmdInteractiveMode = &cobra.Command{
		Use:   "interactive",
		Short: "Enter to interactive mode",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			interactiveMode()
		},
	}

	var rootCmd = &cobra.Command{Use: "tada-runner-generator"}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdInteractiveMode)
	rootCmd.Execute()
}

func init() {
	box = packr.NewBox("./template/runner")
	currentDir, _ = os.Getwd()
	jsonFile, err := os.Open(currentDir + "/package.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	if result["name"] != "runners" {
		fmt.Println("Please run in your runner directory")
		os.Exit(-1)
	}
}

func interactiveMode() {
	name := prompts.NamePrompt()
	copyDir = prompts.CopyFolderPrompt()
	installDep = prompts.InstallDepPrompt()

	exec(name)
}

func exec(name string) {
	TopicName := strcase.ToCamel(name)
	folderName := strcase.ToKebab(name)
	titleName := strcase.ToScreamingDelimited(name, '.', ' ', true)

	data := &struts.RunnerGenerator{
		TopicName:  TopicName,
		FolderName: folderName,
		TitleName:  titleName,
		CurrentDir: currentDir,

		InstallDep: installDep,
		CopyDir:    copyDir,
		PackrBox:   box,
	}

	lib.NewCreateRunner(data)
}
