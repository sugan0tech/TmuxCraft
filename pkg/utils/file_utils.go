package utils

import (
  "fmt"
  "strings"

  "github.com/sugan0tech/tmuxcraft/internal/config"
)

func GenerateTmuxScript(session *config.SessionConfig) string {
  var scriptBuilder strings.Builder

  scriptBuilder.WriteString("#!/bin/bash\n")
  scriptBuilder.WriteString(fmt.Sprintf("tmux new-session -d -s %s\n", session.SessionName))

  for _, window := range session.Windows {
    scriptBuilder.WriteString(fmt.Sprintf("tmux new-window -t %s -n %s\n", session.SessionName, window.Name))
    for _, pane := range window.Panes {
      scriptBuilder.WriteString(fmt.Sprintf("tmux split-window -t %s:%s -p %d -h\n", session.SessionName, window.Name, pane.Percentage))
      scriptBuilder.WriteString(fmt.Sprintf("tmux send-keys -t %s:%s '%s' C-m\n", session.SessionName, window.Name, pane.Command))
    }
  }

  scriptBuilder.WriteString(fmt.Sprintf("tmux attach -t %s\n", session.SessionName))

  return scriptBuilder.String()
}

