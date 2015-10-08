package database

import (
  "encoding/json"
)

type Type int
const (
  Entry_Folder Type = iota
  Entry_Account
)

type Entry interface {
  Type() Type
  GetName() string
}

type Folder struct {
  Name string
  Folders []*Folder
  Accounts []*Account
  Parent *Folder
}

type Account struct {
  Name string
  URL string
  AccountName string
  Password string
  Parent *Folder
}

func Load(blob []byte) *Folder {
  var rootFolder Folder
  json.Unmarshal(blob, &rootFolder)
  return &rootFolder
}

func (folder *Folder) Count() int {
  return len(folder.Folders) + len(folder.Accounts)
}

func (folder *Folder) Items() []Entry {
  items := make([]Entry, folder.Count())
  for i := range folder.Folders {
    items[i] = folder.Folders[i]
  }
  for i := range folder.Accounts {
    items[i] = folder.Accounts[i]
  }

  return items
}

func (_ *Folder) Type() Type {
  return Entry_Folder
}

func (_ *Account) Type() Type {
  return Entry_Account
}

func (folder *Folder) GetName() string {
  return folder.Name
}

func (account *Account) GetName() string {
  return account.Name
}
