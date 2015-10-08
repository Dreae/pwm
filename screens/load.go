package screens

import (
  "os"
  "fmt"
  "bytes"
  "encoding/json"
  "path/filepath"
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
  "github.com/dreae/pwm/database"
)

var items = [3]menuItem{
  menuItem{
    "Load Database",
    loadDraw,
  },
  menuItem{
    "Create New Database",
    defDraw,
  },
  menuItem{
    "Save Database",
    saveDraw,
  },
}

func defDraw(w *draw.Window, statusCh chan string) {
  var buffer bytes.Buffer
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    panic(err)
  }
  buffer.WriteString(dir)

  statusCh<-fmt.Sprintf("Choose a Filename: %s", buffer.String())
  cb := make(chan string)
  go func() {
    for {
      buf, ok := <-cb
      if !ok {
        break
      }
      statusCh<-fmt.Sprintf("Choose a Filename: %s", buf)
    }
  }()
  draw.GetString(&buffer, cb)
  statusCh<-fmt.Sprintf("Final String: %s", buffer.String())
}

func saveDraw(w *draw.Window, statusCh chan string) {
  var buffer bytes.Buffer
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    panic(err)
  }
  buffer.WriteString(dir)

  statusCh<-fmt.Sprintf("Choose a Filename: %s", buffer.String())
  for {
    ev := termbox.PollEvent()
    switch ev.Key {
    case termbox.KeyEnter:
      statusCh<-"Saved"
      return
    case termbox.KeyBackspace:
      if buffer.Len() > 0 {
        buffer.Truncate(buffer.Len() - 1)
      }
    default:
      buffer.WriteRune(ev.Ch)
    }
    statusCh<-fmt.Sprintf("Choose a Filename: %s", buffer.String())
  }
}

func loadDraw(w *draw.Window, statusCh chan string) {
  var buffer bytes.Buffer
  buffer.WriteString("~/.pwm/database")

  statusCh<-fmt.Sprintf("Choose a Filename: %s", buffer.String())
  cb := make(chan string)
  go func() {
    for {
      buf, ok := <-cb
      if !ok {
        break
      }
      statusCh<-fmt.Sprintf("Choose a Filename: %s", buf)
    }
  }()
  draw.GetString(&buffer, cb)

  defer func() {
    if r := recover(); r != nil {
      switch r.(type) {
      case *os.PathError:
        statusCh<-fmt.Sprintf("No such file: %s", buffer.String())
      default:
        statusCh<-fmt.Sprintf("Could not load database: %s", buffer.String())
      }
    }
  }()
  root := LoadFile(buffer.String())
  ScreenList["Database"] = Database(w, root, statusCh)
}

func LoadFile(file string) *database.Folder {
  dbFile, err := os.Open(file)
  defer func() {
    dbFile.Close()
  }()

  if err != nil {
    panic(err)
  }
  decode := json.NewDecoder(dbFile)
  var root database.Folder
  err = decode.Decode(&root)
  if err != nil {
    panic(err)
  }
  return &root
}

type menuItem struct {
  name string
  draw func(*draw.Window, chan string)
}

type LoadScreen struct {
  Title string
  Window *draw.Window
  Selected int
  StatusChannel chan string
}

func Load(w *draw.Window, statusCh chan string) Screen {
  return &LoadScreen{
    Title: "Load Database",
    Window: w,
    Selected: 0,
    StatusChannel: statusCh,
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
  case termbox.KeyEnter:
    items[scr.Selected].draw(scr.Window, scr.StatusChannel)
  }

  col := termbox.ColorDefault
  y := 0
  for i := range items {
    itemCol := col
    if i == scr.Selected {
      itemCol = col | termbox.AttrReverse
    }
    scr.Window.Print(0, y, itemCol, itemCol, items[i].name)
    y += 1
  }
}

func (scr *LoadScreen) GetTitle() string {
  return scr.Title
}
