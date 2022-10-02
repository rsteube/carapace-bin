package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "update packages to use new APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	rootCmd.AddCommand(fixCmd)
}
