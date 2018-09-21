package main

import (
  "testing"
  "github.com/njfix6/tunnel/pkg/folder"
  "github.com/njfix6/tunnel/pkg/file"
  "github.com/stretchr/testify/assert"
)

func TestNoCommand(t *testing.T) {
  args := []string{}
  err := submain(args)
  if err == nil {
    t.Error("Error was not thrown")
  }
}



func TestMCommand(t *testing.T) {

  folder1 := "test_examples/test1"
  folder2 := "test_examples/temp_test"
  cleanUp()
  folder.Create("test_examples/temp_test")
  args := []string{"test", "jobName", folder1, folder2}
  submain(args)
  exists := file.Exists("test_examples/temp_test/test.txt")
  assert.Equal(t, exists, true)
  cleanUp()
}

func cleanUp(){
  folder.Delete("test_examples/temp_test/test.txt")
  folder.Delete("test_examples/temp_test/test2.txt")
  folder.Delete("test_examples/temp_test")
}
