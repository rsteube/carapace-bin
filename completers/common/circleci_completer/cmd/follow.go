package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Attempt to follow the project for the current git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(followCmd).Standalone()
	rootCmd.AddCommand(followCmd)
}
