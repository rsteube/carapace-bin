package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_print_initDefaultsCmd = &cobra.Command{
	Use:   "init-defaults",
	Short: "Print default init configuration, that can be used for 'kubeadm init'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_print_initDefaultsCmd).Standalone()
	config_print_initDefaultsCmd.Flags().StringSlice("component-configs", []string{}, "A comma-separated list for component config API objects to print the default values for.")
	config_printCmd.AddCommand(config_print_initDefaultsCmd)

	carapace.Gen(config_print_initDefaultsCmd).FlagCompletion(carapace.ActionMap{
		"component-configs": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("KubeProxyConfiguration", "KubeletConfiguration").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
