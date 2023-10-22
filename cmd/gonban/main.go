package main

import (
	"github.com/pdarulewski/gonban/internal/markdown"
)

func main() {
	gonbanPath := markdown.FindFile()
	r := markdown.Reader{
		Path: gonbanPath,
	}
	r.Read()

	// TODO: add bubble tea
	b := markdown.Board{}
	b.ParseContent(r.Content)
	b.PrintBoard()

	w := markdown.Writer{
		Path: gonbanPath,
	}
	w.Write(&b)
}
