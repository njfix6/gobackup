package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/njfix6/tunnel/pkg/file"
)

func TestInit(t *testing.T) {
  jsonFile := "./test_examples/testjobs.json"
  err := initConfig(jsonFile)
  assert.Equal(t, err, nil)
  exists := file.Exists(jsonFile)
  assert.Equal(t, exists, true)
}

func TestUpdate(t *testing.T) {
  jsonFile := "./test_examples/testjobs.json"
  // Pass in wrong job
  wrongJob := Job{Name: "testJob", Src: "wrong", Dst: "wrong"}
  config := readConfig(jsonFile)
  config = updateJob(wrongJob, config)
  err := writeConfig(jsonFile, config)
  assert.Equal(t, err, nil)
  config = readConfig(jsonFile)
  assert.Equal(t, config.Jobs[0], wrongJob)


  // Check if it updates the job
  config = readConfig(jsonFile)
  rightjob := Job{Name: "testJob", Src: "test", Dst: "test"}
  config = updateJob(rightjob, config)
  err = writeConfig(jsonFile, config)
  assert.Equal(t, err, nil)
  config = readConfig(jsonFile)
  assert.Equal(t, rightjob, config.Jobs[0])

  file.Delete(jsonFile)
}


func TestGetJob(t *testing.T) {
  config := Config{Jobs: []Job{}}
  job := Job{Name: "testJob", Src: "test", Dst: "test"}
  config = updateJob(job, config)
  testJob, _ := getJob(config, "testJob")
  assert.Equal(t, job, testJob)
}
