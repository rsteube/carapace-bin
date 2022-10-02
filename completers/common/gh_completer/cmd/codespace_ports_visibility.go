package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_ports_visibilityCmd = &cobra.Command{
	Use:   "visibility",
	Short: "Change the visibility of the forwarded port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_ports_visibilityCmd).Standalone()
	codespace_portsCmd.AddCommand(codespace_ports_visibilityCmd)

	carapace.Gen(codespace_ports_visibilityCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				codespace := codespace_ports_visibilityCmd.Flag("codespace").Value.String()
				return action.ActionCodespacePorts(codespace).Invoke(c).Suffix(":").ToA()
			case 1:
				return carapace.ActionValues("public", "private", "org")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
