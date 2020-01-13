package ui

import (
	"github.com/evansloan/nfl-term/api"
	"github.com/rivo/tview"
)

// statPage contains all the stat tables for a
// given stat type
type statPage struct {
	*tview.Grid
	Stats    []*PlayerStats
	name     string
	selected int
}

func newStatPage(stats []*PlayerStats, name string, cols ...int) *statPage {
	s := &statPage{
		Grid:     tview.NewGrid(),
		Stats:    stats,
		name:     name,
		selected: -1,
	}

	s.SetColumns(cols...)

	for i, item := range stats {
		item.SetSelectable(true, false)
		s.AddItem(item, 0, i, 3, 1, 0, 0, true)
	}

	return s
}

// NextCategory returns the next stat table to be
// focused within the app
func (s *statPage) NextCategory() *PlayerStats {
	if s.selected+1 == len(s.Stats) {
		s.selected = 0
	} else {
		s.selected = s.selected + 1
	}
	return s.Stats[s.selected]
}

// StatPages is a container that holds all statPage elements
type StatPages struct {
	*tview.Pages
	oPage      *statPage
	dPage      *statPage
	stPage     *statPage
	ActivePage *statPage
}

// NewStatPages creates statPage elements for all stat types
func NewStatPages() *StatPages {
	s := &StatPages{
		Pages: tview.NewPages(),
	}

	oStats := []*PlayerStats{
		NewPlayerStats("Passing", Passing),
		NewPlayerStats("Rushing", Rushing),
		NewPlayerStats("Receiving", Receiving),
	}
	dStats := []*PlayerStats{
		NewPlayerStats("Defense", Defense),
	}
	stStats := []*PlayerStats{
		NewPlayerStats("Kicking", Kicking),
		NewPlayerStats("Punting", Punting),
		NewPlayerStats("Kick return", KickRet),
		NewPlayerStats("Punt return", PuntRet),
	}
	s.oPage = newStatPage(oStats, "opage", 40, 40, 40)
	s.dPage = newStatPage(dStats, "dpage", 40, 0)
	s.stPage = newStatPage(stStats, "stpage", 30, 30, 30, 30)

	s.AddPage(s.oPage.name, s.oPage, true, true)
	s.AddPage(s.dPage.name, s.dPage, true, false)
	s.AddPage(s.stPage.name, s.stPage, true, false)
	s.ActivePage = s.oPage

	return s
}

// setActive sets a new active statPage and switches to
// it within the application
func (s *StatPages) setActive(page *statPage) {
	s.ActivePage = page
	s.SwitchToPage(page.name)
}

// NextPage switches to the next stat page depending on
// the current active stat page
func (s *StatPages) NextPage() {
	if s.ActivePage == s.oPage {
		s.setActive(s.dPage)
	} else if s.ActivePage == s.dPage {
		s.setActive(s.stPage)
	} else if s.ActivePage == s.stPage {
		s.setActive(s.oPage)
	}
}

// PrevPage switches to the previous stat page depending on
// the current active stat page
func (s *StatPages) PrevPage() {
	if s.ActivePage == s.oPage {
		s.setActive(s.stPage)
	} else if s.ActivePage == s.dPage {
		s.setActive(s.oPage)
	} else if s.ActivePage == s.stPage {
		s.setActive(s.dPage)
	}
}

// SetStats populates the stat pages with their
// respective stats
func (s *StatPages) SetStats(game *api.Game) {
	for _, stat := range s.oPage.Stats {
		stat.SetStats(game)
	}
	for _, stat := range s.dPage.Stats {
		stat.SetStats(game)
	}
	for _, stat := range s.stPage.Stats {
		stat.SetStats(game)
	}
}
