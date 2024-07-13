// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package commands

import (
  "fmt"
  "os"
  "path/filepath"
)

func ListSessions() {
  files, err := os.ReadDir("/home/sugan/.config/tmuxcraft/layouts")
  if err != nil {
    fmt.Println("Error listing sessions:", err)
    return
  }

  fmt.Println("Available sessions:")
  for _, file := range files {
    if filepath.Ext(file.Name()) == ".yaml" {
      fmt.Println(file.Name())
    }
  }
}

