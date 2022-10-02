package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var switchRootCmd = &cobra.Command{
	Use:   "switch-root",
	Short: "Change to a different root file system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchRootCmd).Standalone()

	rootCmd.AddCommand(switchRootCmd)
}
