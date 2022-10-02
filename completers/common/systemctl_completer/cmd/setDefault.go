package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set the default target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setDefaultCmd).Standalone()

	rootCmd.AddCommand(setDefaultCmd)
}
