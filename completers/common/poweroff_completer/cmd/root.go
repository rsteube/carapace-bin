package cmd

import (
	"github.com/rsteube/carapace-bin/completers/halt_completer/cmd"
)

/**
Description for go:generate
	Use: "poweroff",
	Short: "poweroff the machine",
	Long: "https://linux.die.net/man/8/poweroff",
*/

func Execute() error {
	return cmd.ExecutePoweroff()
}
