package main

import (
	"github.com/pdarulewski/gonban/internal/cli"
	"github.com/pdarulewski/gonban/internal/markdown"
)

func main() {
	gonbanPath := markdown.FindFile()
	r := markdown.Reader{
		Path: gonbanPath,
	}
	r.Read()

	b := markdown.Board{}
	b.ParseContent(r.Content)
	board := cli.CreateBoard(&b)
	cli.Run(&board)

	w := markdown.Writer{
		Path: gonbanPath,
	}
	w.Write(&b)
}
