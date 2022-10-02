package cmd

import (
	"github.com/spf13/cobra"
)

var sh_setupCmd = &cobra.Command{
	Use:   "sh-setup",
	Short: "Common Git shell script setup code",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(sh_setupCmd)
}
