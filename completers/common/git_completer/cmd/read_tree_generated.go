package cmd

import (
	"github.com/spf13/cobra"
)

var read_treeCmd = &cobra.Command{
	Use:   "read-tree",
	Short: "Reads tree information into the index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	read_treeCmd.Flags().Bool("aggressive", false, "3-way merge in presence of adds and removes")
	read_treeCmd.Flags().Bool("debug-unpack", false, "debug unpack-trees")
	read_treeCmd.Flags().BoolP("dry-run", "n", false, "don't update the index or the work tree")
	read_treeCmd.Flags().Bool("empty", false, "only empty the index")
	read_treeCmd.Flags().String("exclude-per-directory", "", "allow explicitly ignored files to be overwritten")
	read_treeCmd.Flags().BoolS("i", "i", false, "don't check the working tree after merging")
	read_treeCmd.Flags().String("index-output", "", "write resulting index to <file>")
	read_treeCmd.Flags().BoolS("m", "m", false, "perform a merge in addition to a read")
	read_treeCmd.Flags().Bool("no-sparse-checkout", false, "skip applying sparse checkout filter")
	read_treeCmd.Flags().String("prefix", "", "read the tree into the index under <subdirectory>/")
	read_treeCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	read_treeCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	read_treeCmd.Flags().Bool("reset", false, "same as -m, but discard unmerged entries")
	read_treeCmd.Flags().Bool("trivial", false, "3-way merge if no file level merging required")
	read_treeCmd.Flags().BoolS("u", "u", false, "update working tree with merge result")
	read_treeCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(read_treeCmd)
}
