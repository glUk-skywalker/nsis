package nsis

import (
	"fmt"
	"os/exec"
)

type NSIS struct {
	installerPath string
}

func (i NSIS) InstallTo(path string) error {
	cmd := exec.Command(i.installerPath, "/S", "/NCRC", fmt.Sprintf("/D=%s", path))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("installer running failed: %s", err)
	}
	return nil
}
