package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy previously-created infrastructure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	// TODO copied flags from plan - documentation is insufficient regarding which are actually applicable to destroy as well

	destroyCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	destroyCmd.Flags().BoolS("destroy", "destroy", false, "Select the \"destroy\" planning mode.")
	destroyCmd.Flags().BoolS("detailed-exitcode", "detailed-exitcode", false, "Return detailed exit codes when the command exits.")
	destroyCmd.Flags().StringS("input", "input", "", "Ask for input for variables if not directly set.")
	destroyCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	destroyCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	destroyCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	destroyCmd.Flags().StringS("out", "out", "", "Write a plan file to the given path.")
	destroyCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of concurrent operations.")
	destroyCmd.Flags().StringS("refresh", "refresh", "", "Skip checking for external changes to remote objects while creating the plan.")
	destroyCmd.Flags().BoolS("refresh-only", "refresh-only", false, "Select the \"refresh only\" planning mode.")
	destroyCmd.Flags().StringS("replace", "replace", "", "Force replacement of a particular resource instance using its resource address.")
	destroyCmd.Flags().StringS("state", "state", "", "A legacy option used for the local backend only.")
	destroyCmd.Flags().StringS("target", "target", "", "Limit the planning operation to only the given module, resource.")
	destroyCmd.Flags().StringS("var", "var", "", "Set a value for one of the input variables in the root module of the configuration.")
	destroyCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(destroyCmd)

	destroyCmd.Flag("input").NoOptDefVal = " "
	destroyCmd.Flag("lock").NoOptDefVal = " "
	destroyCmd.Flag("lock-timeout").NoOptDefVal = " "
	destroyCmd.Flag("out").NoOptDefVal = " "
	destroyCmd.Flag("parallelism").NoOptDefVal = " "
	destroyCmd.Flag("refresh").NoOptDefVal = " "
	destroyCmd.Flag("replace").NoOptDefVal = " "
	destroyCmd.Flag("state").NoOptDefVal = " "
	destroyCmd.Flag("target").NoOptDefVal = " "
	destroyCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(destroyCmd).FlagCompletion(carapace.ActionMap{
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
