package struts

import "github.com/gobuffalo/packr"

type RunnerGenerator struct {
	TitleName  string
	FolderName string
	TopicName  string

	CopyDir    bool
	InstallDep bool

	CurrentDir string
	PackrBox   packr.Box
}
