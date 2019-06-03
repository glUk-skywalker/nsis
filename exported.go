package nsis

import "fmt"

func InitInstaller(path string) (Installer, error) {
	var installer Installer

	if !FileExists(path) {
		return installer, fmt.Errorf("file doesn't exist: %s", path)
	}

	installer.path = path

	return installer, nil
}

func InitUninstaller(path string) (Uninstaller, error) {
	var uninstaller Uninstaller

	if !FileExists(path) {
		return uninstaller, fmt.Errorf("file doesn't exist: %s", path)
	}

	uninstaller.path = path

	return uninstaller, nil
}
