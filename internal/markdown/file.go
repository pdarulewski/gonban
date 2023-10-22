package markdown

import (
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
	err := os.MkdirAll(gonbanDirectory, 0666)
	if err != nil {
		panic(err)
	}
	gonbanPath = filepath.Join(gonbanDirectory, "gonban.md")
	return gonbanPath
}
