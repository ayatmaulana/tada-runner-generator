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

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./template/runner")
	currentDir, _ := os.Getwd()
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
		return
	}

	name := prompts.NamePrompt()
	copyDir := prompts.CopyFolderPrompt()
	installDep := prompts.InstallDepPrompt()

	fmt.Println(copyDir)
	fmt.Println(installDep)

	// fmt.Println(name)
	// fmt.Println(copyDir)
	// fmt.Println(installDep)

	TopicName := strcase.ToCamel(name)
	folderName := strcase.ToKebab(name)
	titleName := strcase.ToScreamingDelimited(name, '.', ' ', true)

	runnerGenerator := &struts.RunnerGenerator{
		TopicName:  TopicName,
		FolderName: folderName,
		TitleName:  titleName,
		CurrentDir: currentDir,

		InstallDep: installDep,
		CopyDir:    copyDir,
		PackrBox:   box,
	}

	lib.NewCreateRunner(runnerGenerator)

	// fmt.Println(channelName)
	// fmt.Println(folderName)
	// fmt.Println(titleName)

}
