package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "describe all known top-level flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagsCmd).Standalone()

	rootCmd.AddCommand(flagsCmd)
}
