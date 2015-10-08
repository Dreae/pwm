package screens

import (
  "github.com/nsf/termbox-go"
)

type Screen interface {
  Draw(termbox.Event)
  GetTitle() string
}

var ScreenList map[string]Screen

func init() {
  ScreenList = make(map[string]Screen)
}
