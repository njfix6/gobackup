package main

import (
  "testing"
  "github.com/njfix6/tunnel/pkg/folder"
  "github.com/njfix6/tunnel/pkg/file"
  "github.com/stretchr/testify/assert"
)

func TestNoCommand(t *testing.T) {
  args := []string{}
  err := submain(args, "test_examples/test_jobs.json")
  if err == nil {
    t.Error("Error was not thrown")
  }
}

func TestTooManyArgs(t *testing.T) {
  args := []string{"test", "test", "test", "test", "test"}
  err := submain(args, "test_examples/test_jobs.json")
  assert.NotEqual(t, err, nil)
}



func TestRun(t *testing.T) {

  // Create a job
  src := "test_examples/test_src"
  dst := "test_examples/test_dst"
  cleanUp()
  folder.Create(dst)
  args := []string{"test", "jobName", src, dst}

  configJson := "test_examples/test_jobs.json"
  err := submain(args, configJson)
  assert.Equal(t, err, nil)
  exists := file.Exists(dst + "/test.txt")
  assert.Equal(t, exists, true)

  job := Job{Name: "jobName", Src: src, Dst: dst}
  conf := readConfig(configJson)
  assert.Equal(t, conf.Jobs[0], job)


  // Run that job again
  args = []string{"test", "jobName"}
  err = submain(args, configJson)
  assert.Equal(t, err, nil)

  // cleanUp()
}


func cleanUp(){
  folder.Delete("test_examples/test_dst/test.txt")
  folder.Delete("test_examples/test_dst/test2.txt")
  folder.Delete("test_examples/test_dst")
  folder.Delete("test_examples/test_jobs.json")
}
