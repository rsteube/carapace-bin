package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a headless debug server with a terminal client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).Standalone()
	rootCmd.AddCommand(connectCmd)
}
