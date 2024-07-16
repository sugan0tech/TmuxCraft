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
	"strings"
)

var (
  CONTENTS string
  BUILDER strings.Builder
)

func Init(){
  BUILDER.WriteString("#!/bin/bash\n")
}

func PushCommand(command string){
  if strings.Contains(command, "send-keys") {
    BUILDER.WriteString("sleep 0.1\n")
  }
  BUILDER.WriteString(command + "\n")
}

func GenerateShell(sessionName string, path string) {
  scriptPath := fmt.Sprintf("tmuxcraft_%s.sh", sessionName)
  storePath := filepath.Join(path, scriptPath)
  err := os.WriteFile(storePath, []byte(BUILDER.String()), 0755)
  if err != nil {
    fmt.Println("Error writing shell script:", err)
    return
  }

  fmt.Printf("Shell script generated: %s\n", scriptPath)
}

