package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "faas-cli",
	Short: "Manage your OpenFaaS functions from the command line",
	Long:  "https://github.com/openfaas/faas-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().String("filter", "", "Wildcard to match with function names in YAML file")
	rootCmd.PersistentFlags().String("regex", "", "Regex to match with function names in YAML file")
	rootCmd.PersistentFlags().StringP("yaml", "f", "stack.yml", "Path to YAML file describing function(s)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"yaml": carapace.ActionFiles(),
	})
}
