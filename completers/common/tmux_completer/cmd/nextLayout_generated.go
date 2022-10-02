package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var nextLayoutCmd = &cobra.Command{
	Use:   "next-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextLayoutCmd).Standalone()

	nextLayoutCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(nextLayoutCmd)
}
