package screens

import (
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
  "github.com/dreae/pwm/database"
)

type DatabaseScreen struct {
  Title string
  Database *database.Database
  Parent *database.Entry
  Selected *database.Entry
  Current string
}

func Database() Screen {
  return &DatabaseScreen {
    Title: "Password Database",
    Database: nil,
  }
}

func (scr *DatabaseScreen) Draw(w *draw.Window) {
  if scr.Database == nil {
    w.Print(0, 0, termbox.ColorRed, termbox.ColorDefault, "No database loaded")
  } else {
    scr.renderDatabase(w, scr.Database)
  }
}

func (scr *DatabaseScreen) GetTitle() string {
  return scr.Title
}

func (scr *DatabaseScreen) renderDatabase(w *draw.Window, db *database.Database) {
  if scr.Parent != nil {
    col1 := w.NewWindow(0, 0, 18, w.Height)
    y := 0

    parentEntries := scr.Parent.Value.(map[string]database.Entry)
    for k := range parentEntries {
      if parentEntries[k].Type == database.Entry_Account {
        col := termbox.ColorDefault
        if scr.Parent.Name == k {
          col = termbox.ColorDefault | termbox.AttrReverse | termbox.AttrBold
        }
        col1.Print(0, y, col, col, k)
      }
    }
  }
}
