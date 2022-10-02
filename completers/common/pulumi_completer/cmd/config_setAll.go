package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setAllCmd = &cobra.Command{
	Use:   "set-all",
	Short: "Set multiple configuration values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setAllCmd).Standalone()
	config_setAllCmd.PersistentFlags().Bool("path", false, "Parse the keys as paths in a map or list rather than raw strings")
	config_setAllCmd.PersistentFlags().StringArray("plaintext", []string{}, "Marks a value as plaintext (unencrypted)")
	config_setAllCmd.PersistentFlags().StringArray("secret", []string{}, "Marks a value as secret to be encrypted")
	configCmd.AddCommand(config_setAllCmd)
}
