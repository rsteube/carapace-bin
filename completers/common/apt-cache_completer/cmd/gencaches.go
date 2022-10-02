package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gencachesCmd = &cobra.Command{
	Use:   "gencaches",
	Short: "gencaches creates APT's package cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gencachesCmd).Standalone()

	rootCmd.AddCommand(gencachesCmd)
}
