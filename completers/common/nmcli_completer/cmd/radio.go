package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var radioCmd = &cobra.Command{
	Use:   "radio",
	Short: "NetworkManager radio switches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(radioCmd).Standalone()

	rootCmd.AddCommand(radioCmd)

	carapace.Gen(radioCmd).PositionalCompletion(
		carapace.ActionValues("all", "wifi", "wwan"),
		carapace.ActionValues("on", "off"),
	)
}
