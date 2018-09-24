package main

import (
  "io/ioutil"
  "encoding/json"
  "github.com/njfix6/tunnel/pkg/file"
  "errors"
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
  for i := range config.Jobs {
    if config.Jobs[i].Name == job.Name {
      config.Jobs[i].Src = job.Src
      config.Jobs[i].Dst = job.Dst
      return config
    }
  }
  config.Jobs = append(config.Jobs, job)
  return config
}

func writeConfig(path string, config Config) error {
  configJson, _ := json.Marshal(config)
  err := ioutil.WriteFile(path, configJson, 0644)
  if err != nil {
    return err
  }
  return nil
}

func initConfig(path string) error {
  jobs := []Job{}

  config := Config { Jobs : jobs }
  return writeConfig(path, config)
}

func getJob(config Config, name string) (Job, error) {
  for i := range config.Jobs {
    if config.Jobs[i].Name == name {
      return config.Jobs[i], nil
    }
  }
  return Job{} , errors.New("Job not found")
}
