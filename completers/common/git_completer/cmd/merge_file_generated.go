package cmd

import (
	"github.com/spf13/cobra"
)

var merge_fileCmd = &cobra.Command{
	Use:   "merge-file",
	Short: "Run a three-way file merge",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	merge_fileCmd.Flags().StringS("L", "L", "", "set labels for file1/orig-file/file2")
	merge_fileCmd.Flags().Bool("diff3", false, "use a diff3 based merge")
	merge_fileCmd.Flags().String("marker-size", "", "for conflicts, use this marker size")
	merge_fileCmd.Flags().Bool("ours", false, "for conflicts, use our version")
	merge_fileCmd.Flags().BoolP("quiet", "q", false, "do not warn about conflicts")
	merge_fileCmd.Flags().BoolP("stdout", "p", false, "send results to standard output")
	merge_fileCmd.Flags().Bool("theirs", false, "for conflicts, use their version")
	merge_fileCmd.Flags().Bool("union", false, "for conflicts, use a union version")
	rootCmd.AddCommand(merge_fileCmd)
}
