package markdown

import (
	"strings"
)

type Card struct {
	Name        string
	Description string
}

type Group struct {
	Name string
	Card []*Card
}

type Column struct {
	Title  string
	Groups []*Group
}

type Board struct {
	Title   string
	Columns []*Column
}

func (b *Board) ParseContent(content []string) {
	var currentColumn *Column
	var currentGroup *Group
	var currentCard *Card

	for _, line := range content {
		switch {

		case len(line) == 0:
			continue

		case strings.HasPrefix(line, "# "):
			b.Title = line[2:]

		case strings.HasPrefix(line, "## "):
			name := line[3:]
			currentColumn = &Column{Title: name}
			b.Columns = append(b.Columns, currentColumn)

		case strings.HasPrefix(line, "### "):
			name := line[4:]
			currentGroup = &Group{Name: name}
			currentColumn.Groups = append(currentColumn.Groups, currentGroup)

		case strings.HasPrefix(line, "- "):
			name := line[2:]
			currentCard = &Card{Name: name}
			currentGroup.Card = append(currentGroup.Card, currentCard)

		case strings.HasPrefix(line, "    "):
			currentCard.Description = line[4:]

		case strings.HasPrefix(line, "\t"):
			currentCard.Description = line[1:]
		}
	}
}

func (b *Board) PrintBoard() {
	println(b.Title)
	for _, column := range b.Columns {
		println(column.Title)
		for _, group := range column.Groups {
			println("  " + group.Name)
			for _, card := range group.Card {
				println("    " + card.Name)
				println("      " + card.Description)
			}
		}
	}
}
