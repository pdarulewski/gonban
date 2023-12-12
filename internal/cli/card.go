package cli

import (
	"strings"
)

var lineLength = 30

type Card struct {
	title       string
	description string
}

func (c Card) Title() string {
	if len(c.title) > lineLength {
		return splitString(c.title, lineLength)
	}
	return c.title
}

func (c Card) Description() string {
	if len(c.description) > lineLength {
		return splitString(c.description, lineLength)
	}
	return c.description
}

func (c Card) FilterValue() string { return c.title + c.description }

func splitString(input string, maxLength int) string {
	words := strings.Split(input, " ")
	var result string

	currentLine := ""
	for _, word := range words {
		if len(currentLine)+len(word) <= maxLength {
			currentLine += word + " "
		} else {
			result += strings.TrimSpace(currentLine) + "\n"
			currentLine = word + " "
		}
	}

	result += strings.TrimSpace(currentLine)
	return result
}

// func ItemStyle(c *Column) (s list.DefaultItemStyles) {
// 	normalTitle := lipgloss.NewStyle().
// 		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
// 		Padding(0, 0, 0, 2)
//
// 	normalDescription := s.NormalTitle.Copy().
// 		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})
//
// 	selectedTitle := lipgloss.NewStyle().
// 		Border(lipgloss.NormalBorder(), false, false, false, true).
// 		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
// 		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
// 		Padding(0, 0, 0, 1)
//
// 	selectedDescription := s.SelectedTitle.Copy().
// 		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})
//
// 	dimmedTitle := lipgloss.NewStyle().
// 		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
// 		Padding(0, 0, 0, 2)
//
// 	dimmedDescription := s.DimmedTitle.Copy().
// 		Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})
//
// 	filterMatch := lipgloss.NewStyle().Underline(true)
//
// 	if c.isFocused {
// 		s.NormalTitle = normalTitle
// 		s.NormalDesc = normalDescription
// 		s.SelectedTitle = selectedTitle
// 		s.SelectedDesc = selectedDescription
// 		s.DimmedTitle = dimmedTitle
// 		s.DimmedDesc = dimmedDescription
// 		s.FilterMatch = filterMatch
// 	} else {
// 		s.NormalTitle = normalTitle
// 		s.NormalDesc = normalDescription
// 		s.SelectedTitle = normalTitle
// 		s.SelectedDesc = normalDescription
// 		s.DimmedTitle = dimmedTitle
// 		s.DimmedDesc = dimmedDescription
// 		s.FilterMatch = filterMatch
// 	}
//
// 	return s
// }
