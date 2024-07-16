// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Pane struct {
  Command    string `yaml:"command"`
  Size      int    `yaml:"size"`
  Split string  `yaml:"split"`
}

type Window struct {
  Name  string `yaml:"name"`
  Path  string `yaml:"path"`
  Command  string `yaml:"command"`
  SelectPane int `yaml:"select_pane"`
  Panes []Pane `yaml:"panes"`
}

type SessionConfig struct {
  SessionName string   `yaml:"session_name"`
  Path        string   `yaml:"path"`
  Windows     []Window `yaml:"windows"`
}

func LoadSessionConfig(sessionName string) (*SessionConfig, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println(sessionName)
    log.Fatal("error getting home directory:", err)
  }

  filePath := filepath.Join(homeDir, ".config", "tmuxcraft", "layouts", fmt.Sprintf("%s.yaml", sessionName))
  data, err := os.ReadFile(filePath)
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
 
