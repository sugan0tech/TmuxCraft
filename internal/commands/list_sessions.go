package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListSessions() {
    files, err := os.ReadDir("layouts")
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

