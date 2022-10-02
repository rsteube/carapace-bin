package cmd

import (
	"github.com/spf13/cobra"
)

var interpret_trailersCmd = &cobra.Command{
	Use:   "interpret-trailers",
	Short: "Add or parse structured information in commit messages",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	interpret_trailersCmd.Flags().String("if-exists", "", "action if trailer already exists")
	interpret_trailersCmd.Flags().String("if-missing", "", "action if trailer is missing")
	interpret_trailersCmd.Flags().Bool("in-place", false, "edit files in place")
	interpret_trailersCmd.Flags().Bool("no-divider", false, "do not treat --- specially")
	interpret_trailersCmd.Flags().Bool("only-input", false, "do not apply config rules")
	interpret_trailersCmd.Flags().Bool("only-trailers", false, "output only the trailers")
	interpret_trailersCmd.Flags().Bool("parse", false, "set parsing options")
	interpret_trailersCmd.Flags().String("trailer", "", "trailer(s) to add")
	interpret_trailersCmd.Flags().Bool("trim-empty", false, "trim empty trailers")
	interpret_trailersCmd.Flags().Bool("unfold", false, "join whitespace-continued values")
	interpret_trailersCmd.Flags().String("where", "", "where to place the new trailer")
	rootCmd.AddCommand(interpret_trailersCmd)
}
