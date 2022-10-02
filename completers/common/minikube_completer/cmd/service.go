package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Returns a URL to connect to a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	serviceCmd.PersistentFlags().String("format", "http://{{.IP}}:{{.Port}}", "Format to output service URL in. This format will be applied to each url individually and they will be printed one at a time.")
	serviceCmd.Flags().Bool("https", false, "Open the service URL with https instead of http (defaults to \"false\")")
	serviceCmd.Flags().Int("interval", 1, "The initial time interval for each check that wait performs in seconds")
	serviceCmd.Flags().StringP("namespace", "n", "default", "The service namespace")
	serviceCmd.Flags().Bool("url", false, "Display the Kubernetes service URL in the CLI instead of opening it in the default browser")
	serviceCmd.Flags().Int("wait", 2, "Amount of time to wait for a service in seconds")
	rootCmd.AddCommand(serviceCmd)

	carapace.Gen(serviceCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
