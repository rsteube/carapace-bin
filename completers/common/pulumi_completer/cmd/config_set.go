package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	config_setCmd.PersistentFlags().Bool("path", false, "The key contains a path to a property in a map or list to set")
	config_setCmd.PersistentFlags().Bool("plaintext", false, "Save the value as plaintext (unencrypted)")
	config_setCmd.PersistentFlags().Bool("secret", false, "Encrypt the value instead of storing it in plaintext")
	configCmd.AddCommand(config_setCmd)
}
