package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create or update infrastructure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval of plan before applying.")
	applyCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	applyCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	applyCmd.Flags().StringS("input", "input", "", "Ask for input for variables if not directly set.")
	applyCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	applyCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	applyCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	applyCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations.")
	applyCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	applyCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flag("backup").NoOptDefVal = " "
	applyCmd.Flag("lock-timeout").NoOptDefVal = " "
	applyCmd.Flag("input").NoOptDefVal = " "
	applyCmd.Flag("parallelism").NoOptDefVal = " "
	applyCmd.Flag("state").NoOptDefVal = " "
	applyCmd.Flag("state-out").NoOptDefVal = " "

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"input":     action.ActionBool(),
		"lock":      action.ActionBool(),
		"state":     carapace.ActionFiles(),
		"state-out": carapace.ActionFiles(),
	})
}
