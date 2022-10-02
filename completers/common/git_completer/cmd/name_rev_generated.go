package cmd

import (
	"github.com/spf13/cobra"
)

var name_revCmd = &cobra.Command{
	Use:   "name-rev",
	Short: "Find symbolic names for given revs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	name_revCmd.Flags().Bool("all", false, "list all commits reachable from all refs")
	name_revCmd.Flags().Bool("always", false, "show abbreviated commit object as fallback")
	name_revCmd.Flags().String("exclude", "", "ignore refs matching <pattern>")
	name_revCmd.Flags().Bool("name-only", false, "print only names (no SHA-1)")
	name_revCmd.Flags().String("refs", "", "only use refs matching <pattern>")
	name_revCmd.Flags().Bool("stdin", false, "read from stdin")
	name_revCmd.Flags().Bool("tags", false, "only use tags to name the commits")
	name_revCmd.Flags().Bool("undefined", false, "allow to print `undefined` names (default)")
	rootCmd.AddCommand(name_revCmd)
}
