package cmd

import (
	"github.com/spf13/cobra"
)

var rev_listCmd = &cobra.Command{
	Use:   "rev-list",
	Short: "Lists commit objects in reverse chronological order",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rev_listCmd.Flags().String("abbrev", "", "| --no-abbrev")
	rev_listCmd.Flags().Bool("abbrev-commit", false, "")
	rev_listCmd.Flags().Bool("all", false, "")
	rev_listCmd.Flags().Bool("bisect", false, "")
	rev_listCmd.Flags().Bool("bisect-all", false, "")
	rev_listCmd.Flags().Bool("bisect-vars", false, "")
	rev_listCmd.Flags().Bool("branches", false, "")
	rev_listCmd.Flags().Bool("children", false, "")
	rev_listCmd.Flags().Bool("count", false, "")
	rev_listCmd.Flags().Bool("date-order", false, "")
	rev_listCmd.Flags().String("header", "", "--pretty")
	rev_listCmd.Flags().Bool("left-right", false, "")
	rev_listCmd.Flags().String("max-age", "", "")
	rev_listCmd.Flags().String("max-count", "", "")
	rev_listCmd.Flags().String("max-parents", "", "")
	rev_listCmd.Flags().String("min-age", "", "")
	rev_listCmd.Flags().String("min-parents", "", "")
	rev_listCmd.Flags().Bool("no-max-parents", false, "")
	rev_listCmd.Flags().Bool("no-merges", false, "")
	rev_listCmd.Flags().Bool("no-min-parents", false, "")
	rev_listCmd.Flags().Bool("no-object-names", false, "")
	rev_listCmd.Flags().Bool("object-names", false, "")
	rev_listCmd.Flags().String("objects", "", "--objects-edge")
	rev_listCmd.Flags().Bool("parents", false, "")
	rev_listCmd.Flags().Bool("quiet", false, "")
	rev_listCmd.Flags().Bool("remotes", false, "")
	rev_listCmd.Flags().Bool("remove-empty", false, "")
	rev_listCmd.Flags().Bool("reverse", false, "")
	rev_listCmd.Flags().Bool("sparse", false, "")
	rev_listCmd.Flags().Bool("stdin", false, "")
	rev_listCmd.Flags().Bool("tags", false, "")
	rev_listCmd.Flags().Bool("topo-order", false, "")
	rev_listCmd.Flags().Bool("unpacked", false, "")
	rootCmd.AddCommand(rev_listCmd)
}
