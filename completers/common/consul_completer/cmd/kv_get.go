package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var kv_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves or lists data from the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_getCmd).Standalone()
	addClientFlags(kv_getCmd)
	addServerFlags(kv_getCmd)

	kv_getCmd.Flags().Bool("base64", false, "Base64 encode the value.")
	kv_getCmd.Flags().Bool("detailed", false, "Provide additional metadata about the key in addition to the value.")
	kv_getCmd.Flags().Bool("keys", false, "List keys which start with the given prefix, but not their values.")
	kv_getCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	kv_getCmd.Flags().Bool("recurse", false, "Recursively look at all keys prefixed with the given path.")
	kv_getCmd.Flags().String("separator", "", "String to use as a separator between keys.")
	kvCmd.AddCommand(kv_getCmd)

	// TODO namespace completion

	carapace.Gen(kv_getCmd).PositionalCompletion(
		action.ActionKeys(),
	)
}
