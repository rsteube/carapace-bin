package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var kv_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Removes data from the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_deleteCmd).Standalone()
	addClientFlags(kv_deleteCmd)
	addServerFlags(kv_deleteCmd)

	kv_deleteCmd.Flags().Bool("cas", false, "Perform a Check-And-Set operation.")
	kv_deleteCmd.Flags().String("modify-index", "", "Unsigned integer representing the ModifyIndex of the key.")
	kv_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	kv_deleteCmd.Flags().Bool("recurse", false, "Recursively delete all keys with the path.")
	kvCmd.AddCommand(kv_deleteCmd)

	// TODO namespace completion

	carapace.Gen(kv_deleteCmd).PositionalCompletion(
		action.ActionKeys(),
	)
}
