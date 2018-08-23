package main

import (
  "github.com/mostafah/fsync"
  "os"
  "fmt"
)

func main() {
  args := os.Args
  submain(args)
}

func submain(args []string) {
  if len(args) != 3 {
      fmt.Println("Usage: gobackup <folder1> <folder2>")
      os.Exit(1)
    }
  folder1 := args[1]
  folder2 := args[2]
  fmt.Println("Syncing: "+ folder1 +" to: " + folder2)
  err := fsync.Sync(folder2, folder1)
  if err != nil {
    fmt.Println("Error copying folder", err)
  }
  fmt.Println("Sync Complete")
}
