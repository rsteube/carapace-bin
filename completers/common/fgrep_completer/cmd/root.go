package cmd

import (
	"github.com/rsteube/carapace-bin/completers/grep_completer/cmd"
)

/**
Description for go:generate
	Use: "fgrep",
	Short: "print lines that match patterns",
	Long: "https://en.wikipedia.org/wiki/Grep",
*/

func Execute() error {
	return cmd.ExecuteOverride("fgrep")
}
