package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Show changes required by the current configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(planCmd).Standalone()

	planCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	planCmd.Flags().BoolS("destroy", "destroy", false, "Select the \"destroy\" planning mode.")
	planCmd.Flags().BoolS("detailed-exitcode", "detailed-exitcode", false, "Return detailed exit codes when the command exits.")
	planCmd.Flags().StringS("input", "input", "", "Ask for input for variables if not directly set.")
	planCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	planCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	planCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	planCmd.Flags().StringS("out", "out", "", "Write a plan file to the given path.")
	planCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of concurrent operations.")
	planCmd.Flags().StringS("refresh", "refresh", "", "Skip checking for external changes to remote objects while creating the plan.")
	planCmd.Flags().BoolS("refresh-only", "refresh-only", false, "Select the \"refresh only\" planning mode.")
	planCmd.Flags().StringS("replace", "replace", "", "Force replacement of a particular resource instance using its resource address.")
	planCmd.Flags().StringS("state", "state", "", "A legacy option used for the local backend only.")
	planCmd.Flags().StringS("target", "target", "", "Limit the planning operation to only the given module, resource.")
	planCmd.Flags().StringS("var", "var", "", "Set a value for one of the input variables in the root module of the configuration.")
	planCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(planCmd)

	planCmd.Flag("refresh").NoOptDefVal = " "
	planCmd.Flag("replace").NoOptDefVal = " "
	planCmd.Flag("target").NoOptDefVal = " "
	planCmd.Flag("var-file").NoOptDefVal = " "
	planCmd.Flag("input").NoOptDefVal = " "
	planCmd.Flag("lock").NoOptDefVal = " "
	planCmd.Flag("lock-timeout").NoOptDefVal = " "
	planCmd.Flag("out").NoOptDefVal = " "
	planCmd.Flag("parallelism").NoOptDefVal = " "
	planCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(planCmd).FlagCompletion(carapace.ActionMap{
		"lock":    action.ActionBool(),
		"out":     carapace.ActionFiles(),
		"refresh": action.ActionBool(),
		"replace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources().Invoke(c).ToMultiPartsA(".")
		}),
		"state": carapace.ActionFiles(),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources().Invoke(c).ToMultiPartsA(".")
		}),
		"var-file": carapace.ActionFiles(),
	})
}
