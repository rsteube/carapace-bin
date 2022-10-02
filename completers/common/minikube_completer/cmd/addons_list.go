package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var addons_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available minikube addons as well as their current statuses (enabled/disabled)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	addons_listCmd.Flags().StringP("output", "o", "list", "minikube addons list --output OUTPUT. json, list")
	addonsCmd.AddCommand(addons_listCmd)

	carapace.Gen(addons_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "list"),
	})
}
