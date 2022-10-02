package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var networkingCmd = &cobra.Command{
	Use:   "networking",
	Short: "overall networking control",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkingCmd).Standalone()

	rootCmd.AddCommand(networkingCmd)

	carapace.Gen(networkingCmd).PositionalCompletion(
		carapace.ActionValues("on", "off", "connectivity"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if c.Args[0] == "connectivity" {
				return carapace.ActionValues("check")
			} else {
				return carapace.ActionValues()
			}
		}),
	)
}
