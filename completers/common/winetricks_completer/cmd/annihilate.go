package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var annihilateCmd = &cobra.Command{
	Use:   "annihilate",
	Short: "Delete ALL DATA AND APPLICATIONS INSIDE THIS WINEPREFIX",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(annihilateCmd).Standalone()

	rootCmd.AddCommand(annihilateCmd)
}
