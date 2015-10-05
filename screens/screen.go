package screens

import (
  "github.com/dreae/pwm/draw"
)

type Screen interface {
  Draw(*draw.Window)
  GetTitle() string
}
