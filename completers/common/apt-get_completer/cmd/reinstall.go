package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "reinstall is an alias for install --reinstall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	rootCmd.AddCommand(reinstallCmd)
}
