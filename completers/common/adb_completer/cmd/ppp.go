package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pppCmd = &cobra.Command{
	Use:   "ppp",
	Short: "run PPP over USB",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pppCmd).Standalone()

	rootCmd.AddCommand(pppCmd)
}
