package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var autocleanCmd = &cobra.Command{
	Use:   "autoclean",
	Short: "Remove all outdated packages from .deb package cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocleanCmd).Standalone()

	rootCmd.AddCommand(autocleanCmd)
}
