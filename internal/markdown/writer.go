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

	writeString(f, "# "+b.Title+"\n\n")
	for _, line := range b.Columns {
		writeString(f, "## "+line.Title+"\n\n")
		for _, group := range line.Groups {
			writeString(f, "### "+group.Name+"\n\n")
			for _, card := range group.Card {
				writeString(f, "- "+card.Name+"\n\n")
				writeString(f, "    "+card.Description+"\n\n")
			}
			writeString(f, "\n")
		}
	}
}
