package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/sugan0tech/tmuxcraft/internal/config"
)

var (
	ROOT          string
	SESSION_NAME  string
	CURRENT_WINDOW string
	CURRENT_WINDOW_ID int
	CURRENT_PANE  int
)

func ProjectRoot(root string) {
	ROOT = root
}

func SessionName(name string) {
	SESSION_NAME = name
}

func SetCurrentWindow(name string, id int) {
	CURRENT_WINDOW = name
	CURRENT_WINDOW_ID = id
}

func SetCurrentPane(pane int) {
	CURRENT_PANE = pane
}

func ExecCommand(cmdStr string) {
  cmd := exec.Command("sh", "-c", cmdStr)
  if strings.Contains(cmdStr, "attach-session"){
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
      log.Fatalf("Failed to execute command: %v", err)
    }

  }else {
    err := cmd.Run()
    if err != nil {
      log.Fatalf("Failed to execute command: %v", err)
    }

  }
}

func SelectWindow(window string) {
	cmdStr := fmt.Sprintf("tmux select-window -t %s:%s", SESSION_NAME, window)
	ExecCommand(cmdStr)
}

func TNewSession(session config.SessionConfig) {
	cmdStr := fmt.Sprintf("tmux new-session -d -s %s -c %s", session.SessionName, session.Path)
	ExecCommand(cmdStr)
}

func TNewWindow(window config.Window, id int) {
	currRoot := ROOT
	if window.Path != "" {
		currRoot = window.Path
	}
	cmdStr := fmt.Sprintf("tmux new-window -t %s -c %s -n %s", SESSION_NAME, currRoot, window.Name)
	ExecCommand(cmdStr)

	if window.Command != "" {
		TRunCommand(window.Command, window.Name, 1)
	}
	// SetCurrentWindow(window.Name, id)
}

func TSplitVertical(pane config.Pane, index int) {
	cmdStr := fmt.Sprintf("tmux split-window -t %s:%s.%d -c %s -v -l %d", SESSION_NAME, CURRENT_WINDOW, index+1, ROOT, pane.Size)
	ExecCommand(cmdStr)
	SetCurrentPane(CURRENT_PANE + 1)
}

func TSplitHorizontal(pane config.Pane, index int) {
	CURRENT_PANE = index + 1
	cmdStr := fmt.Sprintf("tmux split-window -t %s:%s.%d -c %s -h -l %d", SESSION_NAME, CURRENT_WINDOW, index+1, ROOT, pane.Size)
	ExecCommand(cmdStr)
	SetCurrentPane(CURRENT_PANE + 1)
}

//  todo: to be used to rename the default window
func TRenameWindow(name string, id int) {
	cmdStr := fmt.Sprintf("tmux rename-window -t %s:%s %s", SESSION_NAME, CURRENT_WINDOW, name)
	ExecCommand(cmdStr)
}

func TAttachToSession() {
	cmdStr := fmt.Sprintf("tmux attach-session -t %s", SESSION_NAME)
	ExecCommand(cmdStr)
}

func TRunCommand(command string, window string, pane int) {
	args := strings.Split(command, " ")
	cmdStr := fmt.Sprintf("tmux send-keys -t %s:%s.%d %s C-m", SESSION_NAME, window, pane, fmt.Sprintf("\"%s\"", strings.Join(args, " ")))
	ExecCommand(cmdStr)
}


