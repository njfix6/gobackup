package main

import (
  "github.com/mostafah/fsync"
  "os"
  "fmt"
  "gopkg.in/cheggaaa/pb.v1"
  "github.com/njfix6/tunnel/pkg/folder"
  "math"
  "errors"
)

func main() {
  args := os.Args
  submain(args)
}


func submain(args []string)  error {
  a := len(args)
  fmt.Println(args)
  if a == 2 {
    return nil
  } else if a == 4{
    jobName := args[1]
    path := "test_examples/test_jobs.json"
    config := readConfig(path)
    job := Job{Name: jobName, Src: "test", Dst: "test"}
    config = updateJob(job, config)
    err := writeConfig(path, config)
    if err != nil {
      return err
    }
    folder1 := args[2]
    folder2 := args[3]
    fmt.Println("Syncing: "+ folder1 +" to: " + folder2)

    syncFolders(folder1, folder2)
    return nil
  } else {
    fmt.Println("Usage: gobackup <job> <folder1> <folder2>")
    fmt.Println("Usage: gobackup <job>")
    return errors.New("Wrong number of inputs")
  }

}

func syncFolders(folder1 string, folder2 string) {

    size1, _ := folder.Size(folder1)
    size2, _ := folder.Size(folder2)
    difference := math.Abs(float64(size1 - size2))
    progress := difference - math.Abs(float64((size1 - size2)))



    bar := pb.New(int(difference))
    bar.SetUnits(pb.U_BYTES)
    bar.Start()



    go func() {
      for progress < difference {

        size2, _  := folder.Size(folder2)
        size1, _  := folder.Size(folder1)

        progress = difference - math.Abs(float64(size1 - size2))
        bar.Set(int(progress))
      }
    }()

    sync(folder1, folder2)

    bar.Set(int(difference))
    bar.Finish()
}


func sync(folder1 string, folder2 string) {
  err := fsync.Sync(folder2, folder1)
  if err != nil {
    fmt.Println("Error copying folder", err)
    os.Exit(1)
  }
}
