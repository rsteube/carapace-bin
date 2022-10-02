package cmd

import (
	"github.com/spf13/cobra"
)

var completion_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "generate autocompletion script for bash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completionCmd.AddCommand(completion_bashCmd)
}
