package markdown

import (
	"log"
	"os"
)

type Writer struct {
	Path string
}

func writeString(f *os.File, s string) {
	_, err := f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Writer) Write(b *Board) {
	f, err := os.Create(r.Path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, line := range b.Columns {
		writeString(f, "# "+line.Name+"\n\n")
		for _, group := range line.Groups {
			writeString(f, "## "+group.Name+"\n")
			for _, card := range group.Card {
				writeString(f, "- "+card+"\n")
			}
			writeString(f, "\n")
		}
	}
}
