package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into pub.dev",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_loginCmd).Standalone()

	pub_loginCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_loginCmd)
}
