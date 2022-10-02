package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toit.pkg",
	Short: "The Toit package manager",
	Long:  "https://github.com/toitlang/toit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("project-root", "", "Specify the project root")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"project-root": carapace.ActionDirectories(),
	})
}
