package draw

import (
  "github.com/nsf/termbox-go"
)

type Window struct {
  X int
  Y int
  Width int
  Height int
}

func (w *Window) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
  termbox.SetCell(x + w.X, y + w.Y, ch, fg, bg)
}

func (w *Window) Fill(x, y, w_, h int, cell termbox.Cell) {
  for ly := 0; ly < h; ly++ {
    for lx := 0; lx < w_; lx++ {
      w.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
    }
  }
}

func (w *Window) Print(x, y int, fg, bg termbox.Attribute, msg string) {
  for _, c := range msg {
    w.SetCell(x, y, c, fg, bg)
    x++
  }
}

func (w *Window) Clear() {
  col := termbox.ColorDefault
  w.Fill(0, 0, w.Width, w.Height, termbox.Cell{Ch: ' ', Fg: col, Bg: col})
}

func (w *Window) NewWindow(x, y, width, height int) *Window {
  return &Window{
    X: w.X + x,
    Y: w.Y + y,
    Width: width,
    Height: height,
  }
}

func NewWindow(x, y, width, height int) *Window {
  return &Window{
    X: x,
    Y: y,
    Width: width,
    Height: height,
  }
}

func TerminalWindow() *Window {
  w, h := termbox.Size()
  return &Window{
    X: 0,
    Y: 0,
    Width: w,
    Height: h,
  }
}
