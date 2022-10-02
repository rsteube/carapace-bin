package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link local directory to a Vercel Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
