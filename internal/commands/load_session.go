// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package commands

import (
  "fmt"

  "github.com/sugan0tech/tmuxcraft/internal/config"
)

func LoadSession(sessionName string) {
  // Logic to load a session from YAML configuration
  session, err := config.LoadSessionConfig(sessionName)
  if err != nil {
    fmt.Println("Error loading session:", err)
    return
  }

  fmt.Println("Session loaded:", session)
}

