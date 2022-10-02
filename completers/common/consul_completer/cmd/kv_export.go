package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kv_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports a tree from the KV store as JSON",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_exportCmd).Standalone()
	addClientFlags(kv_exportCmd)
	addServerFlags(kv_exportCmd)
	kv_exportCmd.Flags().String("namespace", "", "Specifies the namespace to query.")

	kvCmd.AddCommand(kv_exportCmd)

	// TODO namespace completion
}
