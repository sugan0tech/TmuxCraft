package commands

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

var (
	ROOT         string
	SESSION_NAME string
	CURRENT_WINDOW string
	CURRENT_PANE  int
)

func ProjectRoot(root string) {
	ROOT = root
}

func SessionName(name string) {
	SESSION_NAME = name
}

func SetCurrentWindow(name string) {
	CURRENT_WINDOW = name
}

func SetCurrentPane(pane int) {
	CURRENT_PANE = pane
}

func SelectWindow(window string) {
	cmd := exec.Command("tmux", "select-window", "-t", fmt.Sprintf("%s:%s", SESSION_NAME, window))
	cmd.Run()
}

func NewSessionTest() {
	dir, err := filepath.Abs(ROOT)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	cmd := exec.Command("tmux", "new-session", "-d", "-s", SESSION_NAME, "-c", dir)
	cmd.Run()
}

func NewWindow(name string) {
	cmd := exec.Command("tmux", "new-window", "-t", SESSION_NAME, "-c", ROOT, "-n", name)
	cmd.Run()
	SetCurrentWindow(name)
}

func SplitVertical(size int) {
	cmd := exec.Command("tmux", "split-window", "-t", fmt.Sprintf("%s:%s.%d", SESSION_NAME, CURRENT_WINDOW, CURRENT_PANE), "-c", ROOT, "-v", "-l", fmt.Sprintf("%d", size))
	cmd.Run()
	SetCurrentPane(CURRENT_PANE + 1)
}

func SplitHorizontal(size int) {
	cmd := exec.Command("tmux", "split-window", "-t", fmt.Sprintf("%s:%s.%d", SESSION_NAME, CURRENT_WINDOW, CURRENT_PANE), "-c", ROOT, "-h", "-l", fmt.Sprintf("%d", size))
	cmd.Run()
	SetCurrentPane(CURRENT_PANE + 1)
}

func RenameWindow(name string) {
	cmd := exec.Command("tmux", "rename-window", "-t", fmt.Sprintf("%s:%s", SESSION_NAME, CURRENT_WINDOW), name)
	cmd.Run()
}

func AttachToSession() {
	cmd := exec.Command("tmux", "attach-session", "-t", SESSION_NAME)
	cmd.Run()
}

func RunCommand(command string) {
	cmd := exec.Command("tmux", "send-keys", "-t", fmt.Sprintf("%s:%s.%d", SESSION_NAME, CURRENT_WINDOW, CURRENT_PANE), command, "C-m")
	cmd.Run()
}

