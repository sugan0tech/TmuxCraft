// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT


package commands

import "fmt"

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("  load-session (ls)   Load the specified session layout.")
	fmt.Println("  list (l)            List all session layouts.")
	fmt.Println("  new-session (ns)    Create new session layout and open it with $EDITOR.")
	fmt.Println("  edit-session (es)   Edit specified session layout with $EDITOR.")
	fmt.Println("  generate-shell (gs) Generate a shell script for the specified session layout.")
	fmt.Println("  commands (c)        List all tmuxcraft commands.")
	fmt.Println("  version (v)         Print Tmuxcraft version.")
	fmt.Println("  help (h)            Show this message.")
	fmt.Println("  -O                  An optional string flag.")
}

