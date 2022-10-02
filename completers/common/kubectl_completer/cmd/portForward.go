package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var portForwardCmd = &cobra.Command{
	Use:   "port-forward",
	Short: "Forward one or more local ports to a pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portForwardCmd).Standalone()

	portForwardCmd.Flags().String("address", "", "Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When lo")
	portForwardCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	rootCmd.AddCommand(portForwardCmd)

	carapace.Gen(portForwardCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionPorts().Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
