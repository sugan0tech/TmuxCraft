// tmuxcraft
// Copyright (c) 2024 sugan0tech
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/sugan0tech/tmuxcraft/internal/config"
)

var (
	ROOT          string
	SESSION_NAME  string
	CURRENT_WINDOW string
	CURRENT_WINDOW_ID int
	CURRENT_PANE  int
  GENERATE_SHELL bool
  SHELL_WRITE_PATH string
  SHELL_OUT string
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
  if GENERATE_SHELL {
    PushCommand(cmdStr)
  }

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bin/sh"
	}

  cmd := exec.Command(shell, "-c", cmdStr)
  fmt.Println(cmdStr)

  if strings.Contains(cmdStr, "attach-session"){
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
      log.Fatalf("Failed to execute command: %v", err)
    }

  }
  if strings.Contains(cmdStr, "send-keys") {
    time.Sleep(100*time.Millisecond)
    cmdSend := exec.Command(shell, "-c", cmdStr)
    output, err := cmdSend.CombinedOutput()
    if err != nil {
      fmt.Printf("Failed to execute send keys command: %v\nOutput: %s\n", err, string(output))
      return
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

func TNewSession(session config.SessionConfig, flag bool, path string) {
	cmdStr := fmt.Sprintf("tmux new-session -d -s %s -c %s", session.SessionName, session.Path)
  if flag {
    Init()
    GENERATE_SHELL = flag
    if path != "" {
      SHELL_WRITE_PATH, _ = toAbsolutePath(path)
    }
    cwd, err := os.Getwd()
    if err != nil {
      log.Println("Error in getting current working directory: ", err)
    }
    SHELL_WRITE_PATH = cwd
  }
	ExecCommand(cmdStr)
}

func toAbsolutePath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(homeDir, path[1:])
	}

	return filepath.Abs(path)
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
	cmdStr := fmt.Sprintf("tmux -u attach-session -t %s", SESSION_NAME)
  if GENERATE_SHELL {
    GenerateShell(SESSION_NAME, SHELL_WRITE_PATH)
  }
	ExecCommand(cmdStr)
}

func TRunCommand(command string, window string, pane int) {

  args := strings.Split(command, " ")
  cmdStr := fmt.Sprintf("tmux send-keys -t %s:%s.%d %s C-m", SESSION_NAME, window, pane, fmt.Sprintf("\"%s\"", strings.Join(args, " ")))
  ExecCommand(cmdStr)
}


