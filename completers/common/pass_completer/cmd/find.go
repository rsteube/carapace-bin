package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "list passwords that match pass-name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCmd).Standalone()

	rootCmd.AddCommand(findCmd)
}
