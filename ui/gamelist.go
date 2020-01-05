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
	gameList := &GameList{
		List: tview.NewList(),
	}
	gameList.ShowSecondaryText(false)
	gameList.SetBorder(true).
		SetTitle("Games")

	return gameList
}
