package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var l2tpCmd = &cobra.Command{
	Use:   "l2tp",
	Short: "tunnel ethernet over IP (L2TPv3)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(l2tpCmd).Standalone()

	rootCmd.AddCommand(l2tpCmd)
}
