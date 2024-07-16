package commands

import (
	"fmt"
	"log"

	"github.com/sugan0tech/tmuxcraft/internal/config"
)

func LoadSession(templatePath string) {
  session, err := config.LoadSessionConfig(templatePath)
  if err != nil {
    log.Println(err)
    log.Fatal("unable to read the config" + templatePath)
  }

  // Just for development
  for ind, val := range session.Windows{
    fmt.Println(fmt.Sprintf("window %d", ind))
    fmt.Println(val)
  }

  ProjectRoot(session.Path)
  SessionName(session.SessionName)
  TNewSession(*session)

	for windInd, window := range session.Windows {
    if windInd == 0 {
      TRenameWindow(window.Name, windInd + 1)
      SetCurrentWindow(window.Name, windInd + 1)
      // window based command to be executed in first pane
      if window.Command != "" {
        TRunCommand(window.Command, window.Name, 1)
      }
    } else {
		  TNewWindow(window, windInd + 1)
    }
		for ind, pane := range window.Panes {
      switch pane.Split {
      case "h":
				TSplitHorizontal(pane, ind)
      case "v":
				TSplitVertical(pane, ind)
      default:
        fmt.Println("Unspecified Split")
        
      }
      // if default window, need to ignore first pane
      if windInd == 0 {
        TRunCommand(pane.Command, window.Name, ind + 2)
      }else {

        TRunCommand(pane.Command, window.Name, ind + 1)
      }
    }
	}

	TAttachToSession()
}

