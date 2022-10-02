package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var autocleanCmd = &cobra.Command{
	Use:   "autoclean",
	Short: "Like clean, but only removes package files that can no longer be downloaded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocleanCmd).Standalone()

	rootCmd.AddCommand(autocleanCmd)
}
