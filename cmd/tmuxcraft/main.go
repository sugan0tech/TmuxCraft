// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package main

import (
	"flag"
	"fmt"

	"github.com/sugan0tech/tmuxcraft/internal/commands"
)

func main() {
  loadSession := flag.String("load-session", "", "Load the specified session layout.")
  list := flag.Bool("list", false, "List all session layouts.")
  newSession := flag.String("new-session", "", "Create new session layout and open it with $EDITOR.")
  editSession := flag.String("edit-session", "", "Edit specified session layout with $EDITOR.")
  commandsFlag := flag.Bool("commands", false, "List all tmuxcraft commands.")
  version := flag.Bool("version", false, "Print Tmuxcraft version.")
  help := flag.Bool("help", false, "Show this message.")
  generateShell := flag.String("generate-shell", "", "Generate a shell script for the specified session layout.")

  flag.Parse()

  switch {
  case *loadSession != "":
    fmt.Println(loadSession)
    commands.LoadSession(*loadSession)
  case *list:
    commands.ListSessions()
  case *newSession != "":
    commands.NewSession(*newSession)
  case *editSession != "":
    commands.EditSession(*editSession)
  case *commandsFlag:
    commands.ListCommands()
  case *version:
    commands.Version()
  case *help:
    commands.Help()
  case *generateShell != "":
    commands.GenerateShell(*generateShell)
  default:
    commands.Help()
}
}

