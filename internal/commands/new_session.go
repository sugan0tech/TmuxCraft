// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func NewSession(sessionName string) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println(sessionName)
    log.Fatal("error getting home directory:", err)
  }

  sessionFile := filepath.Join(homeDir, ".config", "tmuxcraft", "layouts", fmt.Sprintf("%s.yaml", sessionName))
  file, err := os.Create(sessionFile)

  if err != nil {
    fmt.Println("Error creating session file:", err)
    return
  }
  file.Close()

  editor := os.Getenv("EDITOR")

  if editor == "" {
    editor = "vi"
  }

  cmd := exec.Command(editor, sessionFile)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err = cmd.Run()
  if err != nil {
    fmt.Println("Error opening editor:", err)
  }
}

