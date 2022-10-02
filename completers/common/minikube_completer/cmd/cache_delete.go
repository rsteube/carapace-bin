package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an image from the local cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cacheCmd.AddCommand(cache_deleteCmd)

	carapace.Gen(cache_deleteCmd).PositionalCompletion(
		action.ActionCachedImages(),
	)
}
