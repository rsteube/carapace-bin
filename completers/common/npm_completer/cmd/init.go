package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a package.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()
	initCmd.Flags().BoolP("force", "f", false, "remove various protections against unfortunate side effects")
	initCmd.Flags().Bool("if-present", false, "")
	initCmd.Flags().Bool("scope", false, "create scoped package")
	initCmd.Flags().BoolP("yes", "y", false, "automatically answer yes to any prompts")
	addWorkspaceFlags(initCmd)

	rootCmd.AddCommand(initCmd)
}
