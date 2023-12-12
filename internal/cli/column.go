package cli

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pdarulewski/gonban/internal/markdown"
)

type Column struct {
	isFocused bool
	list      list.Model
	height    int
	width     int
}

func (c Column) Init() tea.Cmd {
	return nil
}

func (c Column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		style := c.getStyle()
		h, v := style.GetFrameSize()
		c.list.SetSize(msg.Width-h, msg.Height-v)
	}
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c Column) View() string {
	return c.getStyle().Render(c.list.View())
}

func (c *Column) getStyle() lipgloss.Style {
	if c.isFocused {
		return lipgloss.NewStyle().
			Padding(1, 2).
			PaddingRight(20).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Height(c.height).
			Width(c.width)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		PaddingRight(20).
		Border(lipgloss.HiddenBorder()).
		Height(c.height).
		Width(c.width)
}

func CreateColumn(c *markdown.Column) Column {
	cards := []list.Item{}

	for _, group := range c.Groups {
		for _, card := range group.Card {
			cards = append(cards, Card{title: card.Name, description: card.Description})
		}
	}

	column := Column{}

	columnList := list.New(cards, list.NewDefaultDelegate(), 0, 0)
	columnList.SetShowHelp(false)
	columnList.Title = c.Title

	column.list = columnList
	return column
}
