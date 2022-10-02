package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs an Architect target with an optional custom builder configuration defined in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}

			splitted := strings.Split(c.Args[0], ":")
			project := splitted[0]
			target := ""
			if len(splitted) > 1 {
				target = splitted[1]
			}
			return action.ActionBuilderConfigurations(project, target).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionProjects().Invoke(c).Suffix(":").ToA()
			case 1:
				return action.ActionTargets(c.Parts[0])
			case 2:
				return action.ActionBuilderConfigurations(c.Parts[0], c.Parts[1])
			default:
				return carapace.ActionValues()
			}
		}),
	)

}
