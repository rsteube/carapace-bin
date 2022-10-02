package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	configCmd.Flags().BoolP("json", "j", false, "Emit output as JSON")
	configCmd.Flags().Bool("show-secrets", false, "Show secret values when listing config instead of displaying blinded values")
	configCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	rootCmd.AddCommand(configCmd)
}
