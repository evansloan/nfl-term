package main

import (
	"github.com/rivo/tview"

	"github.com/evansloan/nfl-term/ui"
)

type App struct {
	*tview.Application
	Layout *ui.Layout
}

func NewApp() *App {
	a := &App{
		Application: tview.NewApplication(),
		Layout:      ui.NewLayout(),
	}
	a.Layout.GameList.AddItem("Quit", "", 'q', func() {
		a.Stop()
	})
	return a
}

func (a *App) RunApp() {
	a.SetRoot(a.Layout.Outer, true).SetFocus(a.Layout.GameList).Run()
}
