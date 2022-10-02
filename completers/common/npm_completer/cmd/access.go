package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Set access level on published packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessCmd).Standalone()
	rootCmd.AddCommand(accessCmd)
}
