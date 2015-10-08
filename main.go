package main

import (
  "log"
  "github.com/nsf/termbox-go"
	"github.com/dreae/pwm/draw"
	"github.com/dreae/pwm/screens"
)

func redraw(screen screens.Screen, event termbox.Event) {
	terminal := draw.TerminalWindow()

	terminal.Print(0, 0, termbox.ColorDefault, termbox.ColorDefault, screen.GetTitle())
	terminal.Fill(0, 1, terminal.Width, 1, termbox.Cell{Ch: '─'})
	screen.Draw(event)
	termbox.Flush()
}

func main() {
  err := termbox.Init()
  if err != nil {
    log.Fatal(err)
  }
  defer termbox.Close()

  terminal := draw.TerminalWindow()
  screenWindow := draw.NewWindow(0, 2, terminal.Width, terminal.Height - 2)

  screenList := struct {
    Database screens.Screen
    LoadDatabase screens.Screen
  }{
    screens.Database(screenWindow),
    screens.Load(screenWindow),
  }

	screen := screenList.Database
	redraw(screen, termbox.Event{})
	for {
    ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				return
			case '1':
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				screen = screenList.Database
			case '2':
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				screen = screenList.LoadDatabase
			}
		}
		redraw(screen, ev)
	}
}
