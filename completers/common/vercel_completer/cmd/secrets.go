package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Manages your global Secrets, for use in Environment Variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsCmd).Standalone()

	rootCmd.AddCommand(secretsCmd)
}
