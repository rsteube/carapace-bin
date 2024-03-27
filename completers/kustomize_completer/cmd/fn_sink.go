package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fn_sinkCmd = &cobra.Command{
	Use:   "sink",
	Short: "[Alpha] Implement a Sink by writing input to a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fn_sinkCmd).Standalone()
	fnCmd.AddCommand(fn_sinkCmd)
}
