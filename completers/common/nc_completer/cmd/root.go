package cmd

import (
	"github.com/rsteube/carapace-bin/completers/netcat_completer/cmd"
)

/**
Description for go:generate
	Use: "nc",
	Short: "simple Unix utility which reads and writes data across network connections",
	Long:  "https://nc110.sourceforge.io/",
*/

func Execute() error {
	return cmd.ExecuteNc()
}
