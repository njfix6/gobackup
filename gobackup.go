package main

import (
  "github.com/mostafah/fsync"
  "os"
  "fmt"
  "gopkg.in/cheggaaa/pb.v1"
  "github.com/njfix6/tunnel/pkg/folder"
  "math"
)

func main() {
  args := os.Args
  submain(args)
}




func sync (folder1 string, folder2 string) {
  err := fsync.Sync(folder2, folder1)
  if err != nil {
    fmt.Println("Error copying folder", err)
  } else {
    fmt.Println("Sync Complete")
  }
}


func submain(args []string) {
  if len(args) != 3 {
      fmt.Println("Usage: gobackup <folder1> <folder2>")
      os.Exit(1)
    }
  folder1 := args[1]
  folder2 := args[2]
  fmt.Println("Syncing: "+ folder1 +" to: " + folder2)

  //

  size1, _ := folder.Size(folder1)
  size2, _ := folder.Size(folder2)
  fmt.Println(size1)
  fmt.Println(size2)
  difference := math.Abs(float64(size1 - size2))
  progress := difference - math.Abs(float64((size1 - size2)))


  fmt.Println(difference)

  bar := pb.New(int(difference))

  bar.Start()

  go sync(folder1, folder2)
  
  for progress < difference {

    size2, _  := folder.Size(folder2)
    size1, _  := folder.Size(folder1)

    progress = difference - math.Abs(float64(size1 - size2))
    bar.Set(int(progress))
  }




  bar.Finish()



}
