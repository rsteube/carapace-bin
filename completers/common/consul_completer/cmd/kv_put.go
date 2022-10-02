package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var kv_putCmd = &cobra.Command{
	Use:   "put",
	Short: "Sets or updates data in the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_putCmd).Standalone()
	addClientFlags(kv_putCmd)
	addServerFlags(kv_putCmd)

	kv_putCmd.Flags().Bool("acquire", false, "Obtain a lock on the key.")
	kv_putCmd.Flags().Bool("base64", false, "Treat the data as base 64 encoded.")
	kv_putCmd.Flags().Bool("cas", false, "Perform a Check-And-Set operation.")
	kv_putCmd.Flags().String("flags", "", "Unsigned integer value to assign to this key-value pair.")
	kv_putCmd.Flags().String("modify-index", "", "Unsigned integer representing the ModifyIndex of the key.")
	kv_putCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	kv_putCmd.Flags().Bool("release", false, "Forfeit the lock on the key at the given path.")
	kv_putCmd.Flags().String("session", "", "User-defined identifer for this session as a string.")
	kvCmd.AddCommand(kv_putCmd)

	// TODO flag completion

	carapace.Gen(kv_putCmd).PositionalCompletion(
		action.ActionKeys(),
	)
}
