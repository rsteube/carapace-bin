package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "resynchronize the package index files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	rootCmd.AddCommand(updateCmd)
}
