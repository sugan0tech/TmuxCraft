// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package main

import (
	"flag"

	"github.com/sugan0tech/tmuxcraft/internal/commands"
)

func main() {
	loadSession := flag.String("load-session", "", "Load the specified session layout.")
	loadSessionAlias := flag.String("ls", "", "Load the specified session layout (alias).")
	list := flag.Bool("list", false, "List all session layouts.")
	listAlias := flag.Bool("l", false, "List all session layouts (alias).")
	newSession := flag.String("new-session", "", "Create new session layout and open it with $EDITOR.")
	newSessionAlias := flag.String("ns", "", "Create new session layout and open it with $EDITOR (alias).")
	editSession := flag.String("edit-session", "", "Edit specified session layout with $EDITOR.")
	editSessionAlias := flag.String("es", "", "Edit specified session layout with $EDITOR (alias).")
	commandsFlag := flag.Bool("commands", false, "List all tmuxcraft commands.")
	commandsFlagAlias := flag.Bool("c", false, "List all tmuxcraft commands (alias).")
	version := flag.Bool("version", false, "Print Tmuxcraft version.")
	versionAlias := flag.Bool("v", false, "Print Tmuxcraft version (alias).")
	help := flag.Bool("help", false, "Show this message.")
	helpAlias := flag.Bool("h", false, "Show this message (alias).")
	generateShell := flag.Bool("generate-shell", false, "Generate a shell script for the specified session layout.")
	generateShellAlias := flag.Bool("gs", false, "Generate a shell script for the specified session layout (alias).")
	output := flag.String("O", "", "An output string flag. for sh file generation")

	flag.Parse()

	switch {
	case *loadSession != "", *loadSessionAlias != "":
		session := *loadSession
		if session == "" {
			session = *loadSessionAlias
		}
		genShell := *generateShell || *generateShellAlias
		commands.LoadSession(session, genShell, *output)
	case *list, *listAlias:
		commands.ListSessions()
	case *newSession != "", *newSessionAlias != "":
		session := *newSession
		if session == "" {
			session = *newSessionAlias
		}
		commands.NewSession(session)
	case *editSession != "", *editSessionAlias != "":
		session := *editSession
		if session == "" {
			session = *editSessionAlias
		}
		commands.EditSession(session)
	case *commandsFlag, *commandsFlagAlias:
		commands.ListCommands()
	case *version, *versionAlias:
		commands.Version()
	case *help, *helpAlias:
		commands.Help()
	default:
		commands.Help()
	}
}
