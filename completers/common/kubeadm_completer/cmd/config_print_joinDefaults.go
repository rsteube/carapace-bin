package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_print_joinDefaultsCmd = &cobra.Command{
	Use:   "join-defaults",
	Short: "Print default join configuration, that can be used for 'kubeadm join'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_print_joinDefaultsCmd).Standalone()
	config_print_joinDefaultsCmd.Flags().StringSlice("component-configs", []string{}, "A comma-separated list for component config API objects to print the default values for. Available values: [KubeProxyConfiguration KubeletConfiguration]. If this flag is not set, no component configs will be printed.")
	config_printCmd.AddCommand(config_print_joinDefaultsCmd)

	carapace.Gen(config_print_joinDefaultsCmd).FlagCompletion(carapace.ActionMap{
		"component-configs": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("KubeProxyConfiguration", "KubeletConfiguration").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
