package ui

import (
	"github.com/rivo/tview"
)

// GameList represents a UI element that displays
// a selectable list of upcoming/ongoing/recently ended
// NFL games.
// Inherits from tview.List
type GameList struct {
	*tview.List
}

// NewGameList creates a new GameList UI element
func NewGameList() *GameList {
	g := &GameList{
		List: tview.NewList(),
	}
	g.ShowSecondaryText(false)
	g.SetBorder(true).
		SetTitle("Games")

	return g
}
