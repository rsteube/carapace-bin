package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of the Pulumi service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()
	logoutCmd.PersistentFlags().Bool("all", false, "Logout of all backends")
	logoutCmd.PersistentFlags().StringP("cloud-url", "c", "", "A cloud URL to log out of (defaults to current cloud)")
	logoutCmd.PersistentFlags().BoolP("local", "l", false, "Log out of using local mode")
	rootCmd.AddCommand(logoutCmd)
}
