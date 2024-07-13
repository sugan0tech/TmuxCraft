package tests

import (
  "testing"

  "github.com/sugan0tech/tmuxcraft/internal/commands"
)

func TestLoadSession(t *testing.T) {
  commands.LoadSession("example_session")
}

func TestListSessions(t *testing.T) {
  commands.ListSessions()
}

func TestGenerateShell(t *testing.T) {
  commands.GenerateShell("example_session")
}

