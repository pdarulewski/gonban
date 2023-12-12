package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pdarulewski/gonban/internal/markdown"
)

var boardStyle = lipgloss.NewStyle()

type Board struct {
	focusedColumn int
	columns       []Column
}

func CreateBoard(b *markdown.Board) Board {
	board := Board{focusedColumn: 0}
	for _, column := range b.Columns {
		board.columns = append(board.columns, CreateColumn(column))
	}
	board.columns[0].isFocused = true
	return board
}

func (b Board) Init() tea.Cmd {
	return nil
}

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return b, tea.Quit

		case "l":
			return goRight(b), nil

		case "h":
			return goLeft(b), nil
		}

	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		for i := 0; i < len(b.columns); i++ {
			var res tea.Model
			res, cmd = b.columns[i].Update(msg)
			b.columns[i] = res.(Column)
			cmds = append(cmds, cmd)
		}
		return b, tea.Batch(cmds...)
	}

	res, cmd := b.columns[b.focusedColumn].Update(msg)
	if _, ok := res.(Column); ok {
		b.columns[b.focusedColumn] = res.(Column)
	} else {
		return res, cmd
	}
	return b, cmd
}

func (b Board) View() string {
	views := []string{}
	for _, column := range b.columns {
		views = append(views, column.View())
	}
	joined := lipgloss.JoinHorizontal(lipgloss.Left, views...)
	return boardStyle.Render(joined)
}

func goRight(b Board) Board {
	if b.focusedColumn < len(b.columns)-1 {
		b.columns[b.focusedColumn].isFocused = false
		b.focusedColumn++
		b.columns[b.focusedColumn].isFocused = true
	}
	return b
}

func goLeft(b Board) Board {
	if b.focusedColumn > 0 {
		b.columns[b.focusedColumn].isFocused = false
		b.focusedColumn--
		b.columns[b.focusedColumn].isFocused = true
	}
	return b
}
