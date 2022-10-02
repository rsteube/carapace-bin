package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "login",
	Short: "Log out of the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()
	logoutCmd.Flags().String("scope", "", "associate with scope")

	rootCmd.AddCommand(logoutCmd)
}
