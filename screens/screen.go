package screens

import (
  "github.com/nsf/termbox-go"
)

type Screen interface {
  Draw(termbox.Event)
  GetTitle() string
}
