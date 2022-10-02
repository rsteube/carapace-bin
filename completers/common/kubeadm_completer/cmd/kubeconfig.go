package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Kubeconfig file utilities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfigCmd).Standalone()
	rootCmd.AddCommand(kubeconfigCmd)
}
