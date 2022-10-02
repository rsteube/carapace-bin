package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tcipCmd = &cobra.Command{
	Use:   "tcip",
	Short: "restart adbd listening on TCP on PORT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcipCmd).Standalone()

	rootCmd.AddCommand(tcipCmd)
}
