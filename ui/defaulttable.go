package ui

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// DefaultTable is the default tview.Table configuration
// for the application.
//
// Inherits from tview.Table
type DefaultTable struct {
	*tview.Table
}

// NewDefaultTable creates a new Table element and
// applies basic styling.
func NewDefaultTable(title string) *DefaultTable {
	g := &DefaultTable{
		Table: tview.NewTable(),
	}
	g.SetTitle(title).
		SetBorder(true)
	g.SetBorders(false)
	return g
}

// SetTextCell creates a basic cell that contains text within
// a table.
//
// Is the same as tview.Table.SetCellSimple() but returns *GenericTable
// instead of *tview.Table.
func (g *DefaultTable) SetTextCell(row int, col int, text string) *DefaultTable {
	g.SetCellSimple(row, col, text)
	return g
}

// SetHeaderCell creates a cell that has a darker color than a standard
// table cell.
// Is used to keep table headers/categories uniform.
func (g *DefaultTable) SetHeaderCell(row int, col int, text string) *DefaultTable {
	g.SetCell(row, col, &tview.TableCell{Text: text, Color: tcell.ColorLightGray})
	return g
}

// SetColorCell creates a cell that contains text of a specific color
// within a table.
func (g *DefaultTable) SetColorCell(row int, col int, text string, color tcell.Color) *DefaultTable {
	g.SetCell(row, col, &tview.TableCell{Text: text, Color: color})
	return g
}

// SetIntCell converts integers into text that can be placed within
// a cell within a table.
func (g *DefaultTable) SetIntCell(row int, col int, text int) *DefaultTable {
	return g.SetTextCell(row, col, strconv.Itoa(text))
}

// SetCustomCell creates a customizable cell that contains text within
// a table.
func (g *DefaultTable) SetCustomCell(row, col int, text string, color tcell.Color, align int) *DefaultTable {
	cell := &tview.TableCell{Text: text, Color: color, Align: align}
	g.SetCell(row, col, cell)
	return g
}

// SetCustomIntCell creates a customizable cell that receives ints as
// its text value.
func (g *DefaultTable) SetCustomIntCell(row, col, text int, color tcell.Color, align int) *DefaultTable {
	return g.SetCustomCell(row, col, strconv.Itoa(text), color, align)
}

func (g *DefaultTable) SetHeader(headerText ...string) *DefaultTable {
	for i, text := range headerText {
		g.SetHeaderCell(0, i, text)
	}
	return g
}
