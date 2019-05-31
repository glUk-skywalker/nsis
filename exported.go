package nsis

import (
	"fmt"
	"os"
)

func Init(path string) (NSIS, error) {
	var installer NSIS

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return installer, fmt.Errorf("installer doesn't exist: %s", err)
	}

	installer.installerPath = path

	return installer, nil
}
