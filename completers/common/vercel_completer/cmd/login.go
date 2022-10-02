package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into your account or creates a new one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	rootCmd.AddCommand(loginCmd)
}
