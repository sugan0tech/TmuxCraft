package commands

import (
    "fmt"
    "os"
    "os/exec"
)

func EditSession(sessionName string) {
    sessionFile := fmt.Sprintf("layouts/%s.yaml", sessionName)
    if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
        fmt.Println("Session file does not exist:", sessionFile)
        return
    }

    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "vi"
    }

    cmd := exec.Command(editor, sessionFile)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        fmt.Println("Error opening editor:", err)
    }
}

