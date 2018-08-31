package main

import (
  "io/ioutil"
  "encoding/json"
  "github.com/njfix6/tunnel/pkg/file"
)

type Job struct {
    Name string
    Src string
    Dst string
}

type Config struct {
  Jobs []Job `json:"jobs"`
}


func readConfig(path string) Config {
  bytes, _ := file.ReadBytes(path)
  var config Config
  json.Unmarshal(bytes, &config)

  return config
}

func updateJob(job Job, config Config) Config {
  return config
}

func writeConfig(path string, config Config) error {
  configJson, _ := json.Marshal(config)
  err := ioutil.WriteFile("path", configJson, 0644)
  if err != nil {
    return err
  }
  return nil
}
