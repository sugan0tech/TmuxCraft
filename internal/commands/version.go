package commands

import "fmt"

const version = "0.1.0"

func Version() {
    fmt.Println("Tmuxcraft version:", version)
}

