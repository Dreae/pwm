package screens

import (
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
)

type LoadScreen struct {
  Title string
}

func Load() Screen {
  return &LoadScreen{
    Title: "Load Database",
  }
}

func (scr *LoadScreen) Draw(w *draw.Window) {
  w.SetCell(0, 0, 'L', termbox.ColorDefault, termbox.ColorDefault)
}

func (scr *LoadScreen) GetTitle() string {
  return scr.Title
}
