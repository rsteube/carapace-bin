package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "List or switch Flutter channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(channelCmd).Standalone()

	channelCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(channelCmd)

	carapace.Gen(channelCmd).PositionalCompletion(
		action.ActionChannels(),
	)
}
