package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/evansloan/nfl-term/api"
	"github.com/evansloan/nfl-term/ui"
)

type App struct {
	*tview.Application
	Layout *ui.Layout
}

func NewApp(games []*api.Game) *App {
	a := &App{
		Application: tview.NewApplication(),
		Layout:      ui.NewLayout(games),
	}

	a.Layout.GameList.AddItem("Quit", "", 'q', func() {
		a.Stop()
	})

	a.setKeyBinds()

	return a
}

func (a *App) RunApp() {
	a.SetRoot(a.Layout.Outer, true).SetFocus(a.Layout.GameList).Run()
}

func (a *App) setKeyBinds() {
	a.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlD {
			a.Layout.PlayerStats.NextPage()
		} else if event.Key() == tcell.KeyCtrlA {
			a.Layout.PlayerStats.PrevPage()
		} else if event.Key() == tcell.KeyTab {
			table := a.Layout.PlayerStats.ActivePage.NextCategory()
			a.SetFocus(table)
		} else if event.Key() == tcell.KeyEsc {
			a.SetFocus(a.Layout.GameList)
		}
		return event
	})
}
