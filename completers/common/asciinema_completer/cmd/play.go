package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Replay terminal recording",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	playCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	playCmd.Flags().StringP("idle-time-limit", "i", "", "limit idle time during playback to given number of seconds")
	playCmd.Flags().StringP("speed", "s", "", "playback speedup (can be fractional)")
	rootCmd.AddCommand(playCmd)

	carapace.Gen(playCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
