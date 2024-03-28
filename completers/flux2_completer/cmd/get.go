package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the resources and their status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()
	getCmd.PersistentFlags().BoolP("all-namespaces", "A", false, "list the requested object(s) across all namespaces")
	getCmd.PersistentFlags().Bool("no-header", false, "skip the header when printing the results")
	getCmd.PersistentFlags().String("status-selector", "", "specify the status condition name and the desired state to filter the get result, e.g. ready=false")
	getCmd.PersistentFlags().BoolP("watch", "w", false, "After listing/getting the requested object, watch for changes.")
	rootCmd.AddCommand(getCmd)
}
