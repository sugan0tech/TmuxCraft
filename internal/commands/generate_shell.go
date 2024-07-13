package commands

import (
	"fmt"
	"os"

	"github.com/sugan0tech/tmuxcraft/internal/config"
	"github.com/sugan0tech/tmuxcraft/pkg/utils"
)

func GenerateShell(sessionName string) {
    session, err := config.LoadSessionConfig(sessionName)
    if err != nil {
        fmt.Println("Error loading session:", err)
        return
    }

    scriptContent := utils.GenerateTmuxScript(session)
    scriptPath := fmt.Sprintf("tmuxcraft_%s.sh", sessionName)
    err = os.WriteFile(scriptPath, []byte(scriptContent), 0755)
    if err != nil {
        fmt.Println("Error writing shell script:", err)
        return
    }

    fmt.Printf("Shell script generated: %s\n", scriptPath)
}

