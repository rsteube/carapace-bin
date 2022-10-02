package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "View or set your gatsby-cli configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optionsCmd).Standalone()

	rootCmd.AddCommand(optionsCmd)

	carapace.Gen(optionsCmd).PositionalCompletion(
		carapace.ActionValues("set"),
		carapace.ActionValues("pm", "package-manager"),
		carapace.ActionValues("npm", "yarn"),
	)
}
