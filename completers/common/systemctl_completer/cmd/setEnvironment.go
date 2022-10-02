package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setEnvironmentCmd = &cobra.Command{
	Use:   "set-environment",
	Short: "Set one or more environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setEnvironmentCmd).Standalone()

	rootCmd.AddCommand(setEnvironmentCmd)

	carapace.Gen(setEnvironmentCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionEnvironmentVariables().Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
