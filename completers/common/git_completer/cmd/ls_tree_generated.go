package cmd

import (
	"github.com/spf13/cobra"
)

var ls_treeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "List the contents of a tree object",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ls_treeCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	ls_treeCmd.Flags().BoolS("d", "d", false, "only show trees")
	ls_treeCmd.Flags().Bool("full-name", false, "use full path names")
	ls_treeCmd.Flags().Bool("full-tree", false, "list entire tree; not just current directory (implies --full-name)")
	ls_treeCmd.Flags().BoolP("long", "l", false, "include object size")
	ls_treeCmd.Flags().Bool("name-only", false, "list only filenames")
	ls_treeCmd.Flags().Bool("name-status", false, "list only filenames")
	ls_treeCmd.Flags().BoolS("r", "r", false, "recurse into subtrees")
	ls_treeCmd.Flags().BoolS("t", "t", false, "show trees when recursing")
	ls_treeCmd.Flags().BoolS("z", "z", false, "terminate entries with NUL byte")
	rootCmd.AddCommand(ls_treeCmd)
}
