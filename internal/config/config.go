// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package config

import (
  "fmt"
  "os"
  "gopkg.in/yaml.v3"
)

type Pane struct {
  Command    string `yaml:"command"`
  Percentage int    `yaml:"percentage"`
}

type Window struct {
  Name  string `yaml:"name"`
  Panes []Pane `yaml:"panes"`
}

type SessionConfig struct {
  SessionName string   `yaml:"session_name"`
  Windows     []Window `yaml:"windows"`
}

func LoadSessionConfig(sessionName string) (*SessionConfig, error) {
  data, err := os.ReadFile(fmt.Sprintf("layouts/%s.yaml", sessionName))
  if err != nil {
    return nil, err
  }

  var config SessionConfig
  err = yaml.Unmarshal(data, &config)
  if err != nil {
    return nil, err
  }

  return &config, nil
}

