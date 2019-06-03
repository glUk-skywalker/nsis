package nsis

import (
	"fmt"
	"os/exec"
)

type Uninstaller struct {
	path string
}

func (u Uninstaller) Uninstall() error {
	cmd := exec.Command(u.path, "/S")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("uninstalling failed: %s", err)
	}
	return nil
}
