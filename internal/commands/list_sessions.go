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

func ListSessions() {
  files, err := os.ReadDir("/home/sugan/.config/tmuxcraft/layouts")
  if err != nil {
    fmt.Println("Error listing sessions:", err)
    return
  }


  var layouts = []string{}

  for _, file := range files {
    if filepath.Ext(file.Name()) == ".yaml" {
      filename := strings.Split(file.Name(), ".")[0]
      layouts = append(layouts, filename)
    }
  }

  if len(layouts) > 0 {
    fmt.Println("Available sessions layouts:")

    for _, layout := range layouts {
      fmt.Println(layout)
    }
  } else {
    fmt.Println("No session layout found, add them under .config/tmuxcraft/layouts")
  }
}

