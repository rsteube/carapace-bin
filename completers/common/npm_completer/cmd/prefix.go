package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Display prefix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prefixCmd).Standalone()
	prefixCmd.Flags().BoolP("global", "g", false, "operate globally")

	rootCmd.AddCommand(prefixCmd)
}
