package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authMethodCmd = &cobra.Command{
	Use:   "auth-method",
	Short: "Auth method management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethodCmd).Standalone()

	rootCmd.AddCommand(authMethodCmd)
}
