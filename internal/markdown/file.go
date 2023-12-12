package markdown

import (
	"errors"
	"os"
	"path/filepath"
)

func FindFile() string {
	gonbanPath, ok := os.LookupEnv("GONBAN_PATH")
	if ok {
		return gonbanPath
	}

	homeDirectory := os.Getenv("HOME")
	gonbanDirectory := filepath.Join(homeDirectory, ".gonban")
	err := os.MkdirAll(gonbanDirectory, 0777)
	if err != nil {
		panic(err)
	}
	gonbanPath = filepath.Join(gonbanDirectory, "gonban.md")

	if _, err := os.Stat(gonbanPath); errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile("filename.txt", []byte("Hello"), 0755)
		if err != nil {
			panic(err)
		}
	}
	return gonbanPath
}
