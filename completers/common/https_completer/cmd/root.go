package cmd

import (
	"github.com/rsteube/carapace-bin/completers/http_completer/cmd"
)

/**
Description for go:generate
	Use: "https",
	Short: "command-line HTTP client for the API era",
	Long: "https://httpie.io/",
*/

func Execute() error {
	return cmd.ExecuteHttps()
}
