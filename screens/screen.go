package screens

import (
  "github.com/dreae/pwm/draw"
  "github.com/nsf/termbox-go"
)

type Screen interface {
  Draw(*draw.Window, termbox.Key)
  GetTitle() string
}
