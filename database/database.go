package database

type EntryType uint16
const (
  Entry_Folder EntryType = iota
  Entry_Account
)

type Entry struct {
  Type EntryType
  Name string
  Value interface{}
}

type Database struct {
  Entries map[string]Entry
}
