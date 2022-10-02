package cmd

import (
	"github.com/spf13/cobra"
)

var check_ignoreCmd = &cobra.Command{
	Use:   "check-ignore",
	Short: "Debug gitignore / exclude files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	check_ignoreCmd.Flags().Bool("no-index", false, "ignore index when checking")
	check_ignoreCmd.Flags().BoolP("non-matching", "n", false, "show non-matching input paths")
	check_ignoreCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	check_ignoreCmd.Flags().Bool("stdin", false, "read file names from stdin")
	check_ignoreCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	check_ignoreCmd.Flags().BoolS("z", "z", false, "terminate input and output records by a NUL character")
	rootCmd.AddCommand(check_ignoreCmd)
}
