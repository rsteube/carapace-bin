package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mrouteCmd = &cobra.Command{
	Use:   "mroute",
	Short: "multicast routing cache entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mrouteCmd).Standalone()

	rootCmd.AddCommand(mrouteCmd)
}
