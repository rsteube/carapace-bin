package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kustomize",
	Short: "Manages declarative configuration of Kubernetes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "help for kustomize")
	rootCmd.PersistentFlags().Bool("stack-trace", false, "print a stack-trace on error")
}
