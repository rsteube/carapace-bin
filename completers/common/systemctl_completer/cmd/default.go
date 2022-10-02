package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Enter system default mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultCmd).Standalone()

	rootCmd.AddCommand(defaultCmd)
}
