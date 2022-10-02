package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Move objects and refs by archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()
	bundleCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	rootCmd.AddCommand(bundleCmd)
}
