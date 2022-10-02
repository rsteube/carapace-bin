package cmd

import (
	"github.com/spf13/cobra"
)

var stripspaceCmd = &cobra.Command{
	Use:   "stripspace",
	Short: "Remove unnecessary whitespace",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	stripspaceCmd.Flags().BoolP("comment-lines", "c", false, "prepend comment character and space to each line")
	stripspaceCmd.Flags().BoolP("strip-comments", "s", false, "skip and remove all lines starting with comment character")
	rootCmd.AddCommand(stripspaceCmd)
}
