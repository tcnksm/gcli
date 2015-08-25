package command

import (
	"os"
)

func TmpChdir(dir string) (func(), error) {

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = os.Chdir(dir)
	if err != nil {
		return nil, err
	}

	return func() {
		os.Chdir(currentDir)
	}, nil

}
