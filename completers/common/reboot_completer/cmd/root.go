package cmd

import (
	"github.com/rsteube/carapace-bin/completers/halt_completer/cmd"
)

/**
Description for go:generate
	Use: "reboot",
	Short: "reboot the machine",
	Long: "https://linux.die.net/man/8/reboot",
*/

func Execute() error {
	return cmd.ExecuteReboot()
}
