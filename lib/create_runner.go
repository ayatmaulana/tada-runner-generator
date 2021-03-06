package lib

import (
	"fmt"
	"os"
	"os/exec"

	mystructs "github.com/ayatmaulana/tada-runner-generator/structs"
	"github.com/fatih/structs"
	"github.com/valyala/fasttemplate"
)

type CreateRunner struct {
	runnerGeneratorData *mystructs.RunnerGenerator
	targetFolder        string
}

func NewCreateRunner(runnerGeneratorData *mystructs.RunnerGenerator) {
	createRunner := &CreateRunner{
		runnerGeneratorData: runnerGeneratorData,
	}

	createRunner.makeDir()
	createRunner.parsingAndCopyFromTemplate()

	if createRunner.runnerGeneratorData.InstallDep {
		createRunner.runNpmInstall()
	}

	if createRunner.runnerGeneratorData.CopyDir {
		createRunner.copyDir()
	}

	fmt.Println("[+] Well Done ✅")

}

func RunCMD(path string, args []string) {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (this *CreateRunner) runNpmInstall() {
	fmt.Println("\n\n[+] Installing NPM package ")
	args := []string{"i", this.targetFolder}
	RunCMD("npm", args)
}

func (this *CreateRunner) copyDir() {
	appDir := this.targetFolder + "/../../app"
	configDir := this.targetFolder + "/../../config"
	localesDir := this.targetFolder + "/../../locales"
	libDir := this.targetFolder + "/../../lib"

	fmt.Println("\n\n[+] Copying folder app, config, lib, locales to your runner folder")

	args := []string{"-R", appDir, configDir, localesDir, libDir, this.targetFolder}
	RunCMD("cp", args)
}

func (this *CreateRunner) makeDir() {
	targetFolder := this.runnerGeneratorData.CurrentDir + "/services/" + this.runnerGeneratorData.FolderName
	err := os.MkdirAll(targetFolder, 0777)

	this.targetFolder = targetFolder

	if err != nil {
		// do some stuff
	}
}

func (this *CreateRunner) parseTemplate(filename string) string {
	file, err := this.runnerGeneratorData.PackrBox.FindString(filename)

	if err != nil {
		// do stuff
	}

	t := fasttemplate.New(file, "{{", "}}")
	s := t.ExecuteString(structs.Map(this.runnerGeneratorData))

	return s
}

func (this *CreateRunner) writeFile(filename, content string, rootPath bool) {
	var path string
	if rootPath {
		path = this.runnerGeneratorData.CurrentDir + "/" + filename
	} else {
		path = this.targetFolder + "/" + filename
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (this *CreateRunner) parsingAndCopyFromTemplate() {
	// parse server.js
	serverJSFile := this.parseTemplate("server.js.template")
	this.writeFile("server.js", serverJSFile, false)

	// parse runner.js
	runnerJSFile := this.parseTemplate("runner.js.template")
	this.writeFile("runner.js", runnerJSFile, false)

	// parse package.json
	packageJSONFile := this.parseTemplate("package.json.template")
	this.writeFile("package.json", packageJSONFile, false)

	// parse Dockerfile
	dockerFile := this.parseTemplate("DockerFile.template")
	this.writeFile("Dockerfile", dockerFile, false)

	// parse CHANGELOG.md
	changeLog := this.parseTemplate("CHANGELOG.md.template")
	this.writeFile("CHANGELOG.md", changeLog, false)

	// parse deployment.yaml
	deploymentYaml := this.parseTemplate("deployment/deployment.yaml.template")
	this.writeFile("deployment/template/"+this.runnerGeneratorData.FolderName+".yaml", deploymentYaml, true)

	// parse .gitlab.ci
	gitlabCIFile := this.parseTemplate("deployment/.gitlab-ci.yaml.template")
	f, err := os.OpenFile(this.runnerGeneratorData.CurrentDir+"/.gitlab-ci.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(gitlabCIFile); err != nil {
		fmt.Println(err)
	}
}
