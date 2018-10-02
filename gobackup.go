package main

import (
  "github.com/mostafah/fsync"
  "os"
  "os/user"
  "fmt"
  "gopkg.in/cheggaaa/pb.v1"
  "github.com/njfix6/tunnel/pkg/folder"
  "github.com/njfix6/tunnel/pkg/file"
  "math"
  "errors"
  "path/filepath"
)

func main() {
  usr, err := user.Current()
  if err != nil {
      fmt.Println( err )
  }
  d := usr.HomeDir + "/.gobackup"
  folder.Create(d)
  f := d + "/config.json"
  file.Create(f)
  args := os.Args
  submain(args, f)
}


func submain(args []string, path string)  error {
  a := len(args)
  if a == 2 {
    jobName := args[1]
    config := readConfig(path)
    job, err := getJob(config, jobName)
    if err != nil {
      return err
    }

    src := job.Src
    dst := job.Dst
    fmt.Println("SYNCING: "+ src +" -> " + dst)
    syncFolders(src, dst)
    return nil
  } else if a == 4{
    jobName := args[1]
    src := args[2]
    dst := args [3]
    config := readConfig(path)
    job := Job{Name: jobName, Src: src, Dst: dst}
    config = updateJob(job, config)
    err := writeConfig(path, config)
    if err != nil {
      return err
    }
    fmt.Println("SYNCING: "+ src +" -> " + dst)
    syncFolders(src, dst)
    return nil
  } else {
    fmt.Println("USAGE:")
    fmt.Println("'gobackup <job> <folder1> <folder2>' - to update or create a job")
    fmt.Println("OR")
    fmt.Println("'gobackup <job>' - to run a job already created")
    return errors.New("Wrong number of inputs")
  }

}

func syncFolders(folder1 string, folder2 string) {
    folder1, _ = filepath.Abs(folder1)
    folder2, _ = filepath.Abs(folder2)
    size1, _ := folder.Size(folder1)
    folder.Create(folder2)
    size2, _ := folder.Size(folder2)
    difference := math.Abs(float64(size1))
    progress := difference - math.Abs(float64((size1 - size2)))

    bar := pb.New(int(difference))
    bar.SetUnits(pb.U_BYTES)
    bar.Set(int(progress))
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
