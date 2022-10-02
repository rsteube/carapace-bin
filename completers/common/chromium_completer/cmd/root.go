package cmd

import (
	"github.com/rsteube/carapace-bin/completers/google-chrome_completer/cmd"
)

/**
Description for go:generate
	Use: "chromium",
	Short: "chromium browser",
	Long: "https://www.chromium.org/Home",
*/

func Execute() error {
	return cmd.ExecuteAlias("chromium")
}
