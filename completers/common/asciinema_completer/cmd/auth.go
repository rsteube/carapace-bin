package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage recordings on asciinema.org account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authCmd).Standalone()

	authCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(authCmd)
}
