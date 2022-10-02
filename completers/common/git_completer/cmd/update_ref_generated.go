package cmd

import (
	"github.com/spf13/cobra"
)

var update_refCmd = &cobra.Command{
	Use:   "update-ref",
	Short: "Update the object name stored in a ref safely",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	update_refCmd.Flags().Bool("create-reflog", false, "create a reflog")
	update_refCmd.Flags().BoolS("d", "d", false, "delete the reference")
	update_refCmd.Flags().StringS("m", "m", "", "reason of the update")
	update_refCmd.Flags().Bool("no-deref", false, "update <refname> not the one it points to")
	update_refCmd.Flags().Bool("stdin", false, "read updates from stdin")
	update_refCmd.Flags().BoolS("z", "z", false, "stdin has NUL-terminated arguments")
	rootCmd.AddCommand(update_refCmd)
}
