package markdown

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type Reader struct {
	Path    string
	Content []string
}

func (r *Reader) Read() {
	if _, err := os.Stat(r.Path); errors.Is(err, os.ErrNotExist) {
		return
	}

	f, err := os.Open(r.Path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		r.Content = append(r.Content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
