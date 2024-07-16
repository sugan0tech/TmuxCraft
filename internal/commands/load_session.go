package commands

import (
	"fmt"
	"log"

	"github.com/sugan0tech/tmuxcraft/internal/config"
)

// Session represents a tmux session configuration
type Session struct {
	SessionName string   `yaml:"session_name"`
	Path        string   `yaml:"path"`
	Windows     []Window `yaml:"windows"`
}

// Window represents a tmux window configuration
type Window struct {
	Name   string `yaml:"name"`
	Panes  []Pane `yaml:"panes"`
	VSplit bool   `yaml:"v_split"`
	HSplit bool   `yaml:"h_split"`
}

// Pane represents a tmux pane configuration
type Pane struct {
	Command    string `yaml:"command"`
	Percentage int    `yaml:"percentage"`
}

func LoadSession(templatePath string) {
  session, err := config.LoadSessionConfig(templatePath)
  if err != nil {
    log.Println(err)
    log.Fatal("unable to read the config" + templatePath)
  }

  fmt.Println(session)
  fmt.Println(session.SessionName)
  fmt.Println(session.Path)
  // fmt.Println(session.Windows)
  for ind, val := range session.Windows{
    fmt.Println(fmt.Sprintf("window %d", ind))
    fmt.Println(val)
  }
	// ProjectRoot(session.Path)
	// SessionName(session.SessionName)
	// NewSession()
	//
	// for _, window := range session.Windows {
	// 	NewWindow(window.Name)
	// 	for _, pane := range window.Panes {
	// 		RunCommand(pane.Command)
	// 		if window.VSplit {
	// 			SplitVertical(pane.Percentage)
	// 		} else if window.HSplit {
	// 			SplitHorizontal(pane.Percentage)
	// 		}
	// 	}
	// }
	//
	// AttachToSession()
}

