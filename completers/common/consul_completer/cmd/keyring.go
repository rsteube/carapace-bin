package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var keyringCmd = &cobra.Command{
	Use:   "keyring",
	Short: "Manages gossip layer encryption keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyringCmd).Standalone()
	addClientFlags(keyringCmd)

	keyringCmd.Flags().String("install", "", "Install a new encryption key.")
	keyringCmd.Flags().Bool("list", false, "List all keys currently in use within the cluster.")
	keyringCmd.Flags().Bool("list-primary", false, "List all primary keys currently in use within the cluster.")
	keyringCmd.Flags().Bool("local-only", false, "Setting this to true will force the keyring query to only hit local servers (no WAN traffic).")
	keyringCmd.Flags().String("relay-factor", "", "Setting this to a non-zero value will cause nodes to relay their response.")
	keyringCmd.Flags().String("remove", "", "Remove the given key from the cluster.")
	keyringCmd.Flags().String("use", "", "Change the primary encryption key, which is used to encrypt messages.")
	rootCmd.AddCommand(keyringCmd)

	// TODO flag completion
}
