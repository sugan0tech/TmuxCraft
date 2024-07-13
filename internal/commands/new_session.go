package commands

import (
    "fmt"
    "os"
    "os/exec"
)

func NewSession() {
    fmt.Println("Creating a new session layout...")

    sessionFile := "layouts/new_session.yaml"
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

