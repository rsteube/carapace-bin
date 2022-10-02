package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of pub.dev",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_logoutCmd).Standalone()

	pub_logoutCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_logoutCmd)
}
