package screens

import (
  "os"
  "io/ioutil"
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
  "github.com/dreae/pwm/database"
)

type DatabaseScreen struct {
  Title string
  Root *database.Folder
  Parent *database.Folder
  Current *database.Folder
  Selected int
}

func Database() Screen {
  dbFile, err := os.Open("database/test.json")
  if err != nil {
    panic(err)
  }
  blob, err := ioutil.ReadAll(dbFile)
  if err != nil {
    panic(err)
  }

  return &DatabaseScreen {
    Title: "Password Database",
    Root: database.Load(blob),
  }
}

func (scr *DatabaseScreen) Draw(w *draw.Window, key termbox.Key) {
  if scr.Root == nil {
    w.Print(0, 0, termbox.ColorRed, termbox.ColorDefault, "No database loaded")
  } else {
    if scr.Parent == nil {
      scr.renderFolder(w, scr.Root)
    } else {
      scr.renderFolder(w, scr.Parent)
    }
  }
}

func (scr *DatabaseScreen) GetTitle() string {
  return scr.Title
}

func (scr *DatabaseScreen) renderFolder(w *draw.Window, folder *database.Folder) {

}
