package commands

import "fmt"

func Help() {
    fmt.Println("Usage: tmuxcraft <command> [<args>]")
    fmt.Println("\nSome useful tmuxcraft commands are:")
    fmt.Println("  load-session (s)   Load the specified session layout.")
    fmt.Println("  list (l)           List all session layouts.")
    fmt.Println("  new-session (ns)   Create new session layout and open it with $EDITOR.")
    fmt.Println("  edit-session (es)  Edit specified session layout with $EDITOR.")
    fmt.Println("  commands           List all tmuxcraft commands.")
    fmt.Println("  version            Print Tmuxcraft version.")
    fmt.Println("  help               Show this message.")
}

