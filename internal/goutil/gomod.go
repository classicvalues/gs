package goutil

import (
	"os"
	"os/exec"
)

func Mod(path string) ([]byte, error) {
	oldPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	defer func() {
		os.Chdir(oldPath)
	}()
	if err := os.Chdir(path); err != nil {
		return nil, err
	}
	return exec.Command(bin(), "mod", "download").CombinedOutput()
}