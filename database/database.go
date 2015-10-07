package database

import (
  "sort"
  "reflect"
  "encoding/json"
)

type Type int

type Folder struct {
  Folders map[string]Folder
  Accounts map[string]Account
}

type Account struct {
  URL string
  AccountName string
  Password string
}

func Load(blob []byte) *Folder {
  var rootFolder Folder
  json.Unmarshal(blob, &rootFolder)
  return &rootFolder
}

func (folder *Folder) FolderKeys() []string {
  return keys(folder.Folders)
}

func (folder *Folder) AccountKeys() []string {
  return keys(folder.Accounts)
}

func keys(obj interface{}) []string {
  v := reflect.ValueOf(obj)
  ks := v.MapKeys()

  kStrings := make([]string, len(ks))
  for i := range ks {
    kStrings[i] = ks[i].String()
  }

  sort.Strings(kStrings)

  return kStrings
}
