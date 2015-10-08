package screens

import (
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
)

var items = [2]string{"Load Database","Create New Database"}

type LoadScreen struct {
  Title string
  Window *draw.Window
  Selected int
}

func Load(w *draw.Window) Screen {
  return &LoadScreen{
    Title: "Load Database",
    Window: w,
    Selected: 0,
  }
}

func (scr *LoadScreen) Draw(event termbox.Event) {
  switch event.Key {
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
    scr.Window.Print(0, y, itemCol, itemCol, items[i])
    y += 1
  }
}

func (scr *LoadScreen) GetTitle() string {
  return scr.Title
}
