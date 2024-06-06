package dir

import (
	"os/exec"

	pathHelper "github.com/mikeunge/go/pkg/path-helper"
)

type Dir struct {
	Path string
}

func New(path string) Dir {
	return Dir{
		Path: pathHelper.SanitizePath(path),
	}
}

func (d *Dir) ChangeDir() error {
	cmd := exec.Command("bash", "-c", "cd", d.Path)
	return cmd.Run()
}
