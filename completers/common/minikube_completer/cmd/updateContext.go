package cmd

import (
	"github.com/spf13/cobra"
)

var updateContextCmd = &cobra.Command{
	Use:   "update-context",
	Short: "Update kubeconfig in case of an IP or port change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(updateContextCmd)
}
