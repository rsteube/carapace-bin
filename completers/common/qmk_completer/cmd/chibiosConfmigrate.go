package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var chibiosConfmigrateCmd = &cobra.Command{
	Use:   "chibios-confmigrate",
	Short: "Generates a migrated ChibiOS configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chibiosConfmigrateCmd).Standalone()

	chibiosConfmigrateCmd.Flags().BoolP("delete", "d", false, "migration will delete the input file")
	chibiosConfmigrateCmd.Flags().BoolP("force", "f", false, "Re-migrates an already migrated file")
	chibiosConfmigrateCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	chibiosConfmigrateCmd.Flags().StringP("input", "i", "", "Specify input config file.")
	chibiosConfmigrateCmd.Flags().BoolP("overwrite", "o", false, "Overwrites the input file during migration.")
	chibiosConfmigrateCmd.Flags().StringP("reference", "r", "", "Specify the reference file to compare against")
	rootCmd.AddCommand(chibiosConfmigrateCmd)

	carapace.Gen(chibiosConfmigrateCmd).FlagCompletion(carapace.ActionMap{
		"input":     carapace.ActionFiles(),
		"reference": carapace.ActionFiles(),
	})
}
