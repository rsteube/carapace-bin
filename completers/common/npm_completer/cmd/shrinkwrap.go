package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var shrinkwrapCmd = &cobra.Command{
	Use:   "shrinkwrap",
	Short: "Lock down dependency versions for publication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shrinkwrapCmd).Standalone()

	rootCmd.AddCommand(shrinkwrapCmd)
}
