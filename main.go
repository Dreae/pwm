package main

import (
  "log"
  "github.com/nsf/termbox-go"
	"github.com/dreae/pwm/draw"
	"github.com/dreae/pwm/screens"
)

func redraw(screen screens.Screen, key termbox.Key) {
	terminal := draw.TerminalWindow()

	terminal.Print(0, 0, termbox.ColorDefault, termbox.ColorDefault, screen.GetTitle())
	terminal.Fill(0, 1, terminal.Width, 1, termbox.Cell{Ch: '─'})
	screen.Draw(draw.NewWindow(0, 2, terminal.Width, terminal.Height - 2), key)
	termbox.Flush()
}

func main() {
  err := termbox.Init()
  if err != nil {
    log.Fatal(err)
  }
  defer termbox.Close()

	screen := screens.Database()
	redraw(screen, termbox.KeyEsc)
	for {
    ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				return
			case '1':
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				screen = screens.Database()
			case '2':
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				screen = screens.Load()
			}
		}
		redraw(screen, ev.Key)
	}
}
