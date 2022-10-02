package cmd

import (
	"github.com/spf13/cobra"
)

var verify_tagCmd = &cobra.Command{
	Use:   "verify-tag",
	Short: "Check the GPG signature of tags",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	verify_tagCmd.Flags().String("format", "", "format to use for the output")
	verify_tagCmd.Flags().Bool("raw", false, "print raw gpg status output")
	verify_tagCmd.Flags().BoolP("verbose", "v", false, "print tag contents")
	rootCmd.AddCommand(verify_tagCmd)
}
