package cmd

import (
	"github.com/spf13/cobra"
)

var completion_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish completion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completionCmd.AddCommand(completion_fishCmd)
}
