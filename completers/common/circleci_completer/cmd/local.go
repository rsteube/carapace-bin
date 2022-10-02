package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Debug jobs on the local machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localCmd).Standalone()
	rootCmd.AddCommand(localCmd)
}
