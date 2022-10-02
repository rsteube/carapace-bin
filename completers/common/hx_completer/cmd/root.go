package cmd

import (
	"github.com/rsteube/carapace-bin/completers/helix_completer/cmd"
)

/**
Description for go:generate
	Use: "hx",
	Short: "A post-modern text editor",
	Long:  "https://helix-editor.com/",
*/

func Execute() error {
	return cmd.ExecuteHx()
}
