package helper

import (
	"os"
	"path/filepath"
)

func GetProjectRootDir() string {
	exe, _ := os.Executable()
	pjRootDir := filepath.Join(filepath.Dir(exe), "../..")
	return pjRootDir
}
