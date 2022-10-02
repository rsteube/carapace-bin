package cmd

import (
	"github.com/rsteube/carapace-bin/completers/aplay_completer/cmd"
)

/**
Description for go:generate
	Use: "arecord",
	Short: "command-line sound recorder and player for ALSA soundcard driver",
	Long: "https://linux.die.net/man/1/arecord",
*/

func Execute() error {
	return cmd.ExecuteArecord()
}
