package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addons_openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the addon w/ADDON_NAME within minikube (example: minikube addons open dashboard). For a list of available addons use: minikube addons list ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	addons_openCmd.PersistentFlags().String("format", "http://{{.IP}}:{{.Port}}", "Format to output addons URL in.  This format will be applied to each url individually and they will be printed one at a time.")
	addons_openCmd.Flags().Bool("https", false, "Open the addons URL with https instead of http")
	addons_openCmd.Flags().Int("interval", 1, "The time interval for each check that wait performs in seconds")
	addons_openCmd.Flags().Bool("url", false, "Display the Kubernetes addons URL in the CLI instead of opening it in the default browser")
	addons_openCmd.Flags().Int("wait", 2, "Amount of time to wait for service in seconds")
	addonsCmd.AddCommand(addons_openCmd)

	carapace.Gen(addons_openCmd).PositionalCompletion(
		action.ActionAddons(),
	)
}
