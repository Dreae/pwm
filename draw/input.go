package draw

import (
  "bytes"
  "github.com/nsf/termbox-go"
)

func GetString(buffer *bytes.Buffer, cb chan string) {
  for {
    ev := termbox.PollEvent()
    switch ev.Key {
    case termbox.KeyEnter:
      close(cb)
      return
    case termbox.KeyBackspace:
      if buffer.Len() > 0 {
        buffer.Truncate(buffer.Len() - 1)
      }
    default:
      buffer.WriteRune(ev.Ch)
    }
    cb<-buffer.String()
  }
}
