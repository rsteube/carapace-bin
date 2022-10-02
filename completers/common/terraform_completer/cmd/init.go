package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prepare your working directory for other commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringS("backend", "backend", "", "Configure the backend for this configuration.")
	initCmd.Flags().StringSliceS("backend-config", "backend-config", []string{}, "Backend configuration")
	initCmd.Flags().BoolS("force-copy", "force-copy", false, "Suppress prompts about copying state data.")
	initCmd.Flags().StringS("from-module", "from-module", "", "Copy the contents of the given module into the target directory before initialization.")
	initCmd.Flags().StringS("get", "get", "", "Download any modules for this configuration.")
	initCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	initCmd.Flags().StringS("input", "input", "", "Ask for input if necessary.")
	initCmd.Flags().StringS("lockfile", "lockfile", "", "Set a dependency lockfile mode.")
	initCmd.Flags().BoolS("migrate-state", "migrate-state", false, "Reconfigure the backend, and attempt to migrate any existing state.")
	initCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	initCmd.Flags().BoolS("plugin-dir", "plugin-dir", false, "Directory containing plugin binaries.")
	initCmd.Flags().BoolS("reconfigure", "reconfigure", false, "Reconfigure the backend, ignoring any saved configuration.")
	initCmd.Flags().StringS("upgrade", "upgrade", "", "ignore previously-downloaded objects and install the latest version allowed")
	rootCmd.AddCommand(initCmd)

	initCmd.Flag("backend").NoOptDefVal = " "
	initCmd.Flag("backend-config").NoOptDefVal = " "
	initCmd.Flag("from-module").NoOptDefVal = " "
	initCmd.Flag("get").NoOptDefVal = " "
	initCmd.Flag("input").NoOptDefVal = " "
	initCmd.Flag("lockfile").NoOptDefVal = " "
	initCmd.Flag("upgrade").NoOptDefVal = " "

	// TODO module completion
	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"backend":        action.ActionBool(),
		"backend-config": carapace.ActionFiles(),
		"get":            action.ActionBool(),
		"input":          action.ActionBool(),
		"lockfile":       carapace.ActionValues("readonly"),
		"upgrade":        action.ActionBool(),
	})
}
