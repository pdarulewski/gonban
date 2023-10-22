package markdown

type Group struct {
	Name string
	Card []string
}

type Column struct {
	Name   string
	Groups []*Group
}

type Board struct {
	Columns []*Column
}

func (b *Board) ParseContent(content []string) {
	var currentColumn *Column
	var currentGroup *Group

	for _, line := range content {
		switch {

		case len(line) == 0:
			continue

		case line[0] == '#' && line[1] == ' ':
			name := line[2:]
			currentColumn = &Column{Name: name}
			b.Columns = append(b.Columns, currentColumn)

		case line[0] == '#' && line[1] == '#':
			name := line[3:]
			currentGroup = &Group{Name: name}
			currentColumn.Groups = append(currentColumn.Groups, currentGroup)

		case line[0] == '-':
			name := line[2:]
			currentGroup.Card = append(currentGroup.Card, name)
		}
	}
}

func (b *Board) PrintBoard() {
	for _, column := range b.Columns {
		println(column.Name)
		for _, group := range column.Groups {
			println("  " + group.Name)
			for _, card := range group.Card {
				println("    " + card)
			}
		}
	}
}
