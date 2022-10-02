package cmd

import (
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Add or inspect object notes",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	notesCmd.Flags().String("ref", "", "use notes from <notes-ref>")
	rootCmd.AddCommand(notesCmd)
}
