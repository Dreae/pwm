package screens

import (
  "os"
  "fmt"
  "io/ioutil"
  "github.com/nsf/termbox-go"
  "github.com/dreae/pwm/draw"
  "github.com/atotto/clipboard"
  "github.com/dreae/pwm/database"
)

type DatabaseScreen struct {
  Title string
  Window *draw.Window
  Root *database.Folder
  Path []Node
  Current *database.Folder
  Selected int
  ParentColumn *draw.Window
  CurrentColumn *draw.Window
  NextColumn *draw.Window
}

type Node struct {
  Selected *database.Folder
  SelectedIndex int
}

func Database(w *draw.Window) Screen {
  dbFile, err := os.Open("database/test.json")
  if err != nil {
    panic(err)
  }
  blob, err := ioutil.ReadAll(dbFile)
  if err != nil {
    panic(err)
  }

  return &DatabaseScreen {
    Title: "Password Database",
    Window: w,
    Root: database.Load(blob),
    ParentColumn: w.NewWindow(0, 0, 18, w.Height),
    CurrentColumn: w.NewWindow(19, 0, 32, w.Height),
    NextColumn: w.NewWindow(52, 0, 32, w.Height),
  }
}

func (scr *DatabaseScreen) Draw(event termbox.Event) {
  switch event.Key {
  case termbox.KeyArrowUp:
    scr.Selected = scr.Selected - 1
  case termbox.KeyArrowDown:
    scr.Selected = scr.Selected + 1
  case termbox.KeyArrowRight:
    item := scr.Current.Items()[scr.Selected]
    if item.Type() == database.Entry_Account {
      account := item.(*database.Account)
      clipboard.WriteAll(account.Password)
    } else {
      node := Node{
        Selected: scr.Current,
        SelectedIndex: scr.Selected,
      }

      scr.Path = append(scr.Path, node)
      scr.Current = scr.Current.Folders[scr.Selected]
      scr.Selected = 0
    }
  case termbox.KeyArrowLeft:
    if len(scr.Path) != 0 {
      var node Node
      node, scr.Path = scr.Path[len(scr.Path)-1], scr.Path[:len(scr.Path)-1]
      scr.Current = node.Selected
      scr.Selected = node.SelectedIndex
    }
  }
  
  switch event.Ch {
  case 'c':
    item := scr.Current.Items()[scr.Selected]
    if item.Type() == database.Entry_Account {
      account := item.(*database.Account)
      clipboard.WriteAll(account.Password)
    }
  }

  if scr.Root == nil {
    scr.Window.Print(0, 0, termbox.ColorRed, termbox.ColorDefault, "No database loaded")
  } else {
    if len(scr.Path) == 0 {
      scr.Current = scr.Root
      scr.ParentColumn.Clear()
    } else {
      node := scr.Path[len(scr.Path) - 1]
      scr.renderFolder(scr.ParentColumn, node.Selected, node.SelectedIndex)
    }
    count := scr.Current.Count()
    if scr.Selected == count {
      scr.Selected = 0
    } else if scr.Selected < 0 {
      scr.Selected = count - 1
    }

    scr.renderFolder(scr.CurrentColumn, scr.Current, scr.Selected)
    item := scr.Current.Items()[scr.Selected]
    if item.Type() == database.Entry_Folder {
      scr.renderFolder(scr.NextColumn, item.(*database.Folder), -1)
    } else {
      scr.renderAccount(scr.NextColumn, item.(*database.Account))
    }
  }
}

func (scr *DatabaseScreen) GetTitle() string {
  return scr.Title
}

func (scr *DatabaseScreen) renderFolder(w *draw.Window, folder *database.Folder, highlight int) {
  w.Clear()
  y := 0
  items := folder.Items()

  // TODO: Paginate
  for i := range items {
    fg := termbox.ColorDefault
    bg := termbox.ColorDefault
    if items[i].Type() == database.Entry_Folder {
      fg = termbox.ColorCyan
    }
    if highlight != -1 && y == highlight {
      fg = fg | termbox.AttrReverse
      bg = bg | termbox.AttrReverse
      w.Fill(0, y, w.Width, 1, termbox.Cell{Bg: bg, Fg: fg})
    } else {
      w.Fill(0, y, w.Width, 1, termbox.Cell{Bg: bg, Fg: fg})
    }
    w.Print(0, y, fg, bg, items[i].GetName())

    y = y + 1
  }
}

func (scr *DatabaseScreen) renderAccount(w *draw.Window, account *database.Account) {
  w.Print(0, 0, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("Username: %s", account.AccountName))
  w.Print(0, 1, termbox.ColorDefault, termbox.ColorDefault, fmt.Sprintf("URL: %s", account.URL))
  w.Print(0, 2, termbox.ColorDefault, termbox.ColorDefault, "Password: ••••••••")
}
