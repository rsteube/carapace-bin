package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing project or group variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_updateCmd).Standalone()
	variable_updateCmd.Flags().StringP("group", "g", "", "Set variable for a group")
	variable_updateCmd.Flags().BoolP("masked", "m", false, "Whether the variable is masked")
	variable_updateCmd.Flags().BoolP("protected", "p", false, "Whether the variable is protected")
	variable_updateCmd.Flags().StringP("scope", "s", "*", "The environment_scope of the variable. All (*), or specific environments")
	variable_updateCmd.Flags().StringP("type", "t", "env_var", "The type of a variable: {env_var|file}")
	variable_updateCmd.Flags().StringP("value", "v", "", "The value of a variable")
	variableCmd.AddCommand(variable_updateCmd)

	carapace.Gen(variable_updateCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_updateCmd),
		"scope": carapace.Batch(
			carapace.ActionValues("*"),
			action.ActionEnvironments(variable_updateCmd),
		).ToA(),
		"type": carapace.ActionValues("env_var", "file"),
	})

	carapace.Gen(variable_updateCmd).PositionalCompletion(
		action.ActionVariables(variable_updateCmd),
	)
}
