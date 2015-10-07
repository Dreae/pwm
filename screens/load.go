package screens

import (
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
)

var items = [2]string{"Load Database","Create New Database"}

type LoadScreen struct {
  Title string
  Selected int
}

func Load() Screen {
  return &LoadScreen{
    Title: "Load Database",
    Selected: 0,
  }
}

func (scr *LoadScreen) Draw(w *draw.Window, key termbox.Key) {
  switch key {
  case termbox.KeyArrowDown:
    scr.Selected = scr.Selected + 1
    if scr.Selected == len(items) {
      scr.Selected = 0
    }
  case termbox.KeyArrowUp:
    scr.Selected = scr.Selected - 1
    if scr.Selected < 0 {
      scr.Selected = len(items) - 1
    }
  }

  col := termbox.ColorDefault
  y := 0
  for i := range items {
    itemCol := col
    if i == scr.Selected {
      itemCol = col | termbox.AttrReverse
    }
    w.Print(0, y, itemCol, itemCol, items[i])
    y += 1
  }
}

func (scr *LoadScreen) GetTitle() string {
  return scr.Title
}
