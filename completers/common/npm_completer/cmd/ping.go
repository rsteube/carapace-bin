package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pingCmd).Standalone()

	rootCmd.AddCommand(pingCmd)
}
