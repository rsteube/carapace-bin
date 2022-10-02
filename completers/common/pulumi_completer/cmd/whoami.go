package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display the current logged-in user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()
	rootCmd.AddCommand(whoamiCmd)
}
