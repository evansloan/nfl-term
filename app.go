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

	a.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlO {
			a.Layout.PlayerStats.SwitchToPage("opage")
			a.Layout.PlayerStats.SetActive("opage")
		} else if event.Key() == tcell.KeyCtrlD {
			a.Layout.PlayerStats.SwitchToPage("dpage")
			a.Layout.PlayerStats.SetActive("dpage")
		} else if event.Key() == tcell.KeyTab {
			table := a.Layout.PlayerStats.ActivePage.SelectNext()
			a.SetFocus(table)
		} else if event.Key() == tcell.KeyEsc {
			a.SetFocus(a.Layout.GameList)
		}
		return event
	})

	a.Layout.GameList.AddItem("Quit", "", 'q', func() {
		a.Stop()
	})

	return a
}

func (a *App) RunApp() {
	a.SetRoot(a.Layout.Outer, true).SetFocus(a.Layout.GameList).Run()
}
