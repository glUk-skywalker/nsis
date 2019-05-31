package nsis

import (
	"fmt"
	"os/exec"
)

type Installer struct {
	path string
}

func (i Installer) InstallTo(path string) error {
	cmd := exec.Command(i.path, "/S", "/NCRC", fmt.Sprintf("/D=%s", path))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("installer running failed: %s", err)
	}
	return nil
}
