package cmd

import (
	"github.com/rsteube/carapace-bin/completers/code_completer/cmd"
)

/**
Description for go:generate
	Use: "code-insiders",
	Short: "Visual Studio Code Insiders",
	Long: "https://code.visualstudio.com/insiders/",
*/

func Execute() error {
	return cmd.ExecuteInsiders()
}
