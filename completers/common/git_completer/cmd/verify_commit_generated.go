package cmd

import (
	"github.com/spf13/cobra"
)

var verify_commitCmd = &cobra.Command{
	Use:   "verify-commit",
	Short: "Check the GPG signature of commits",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	verify_commitCmd.Flags().Bool("raw", false, "print raw gpg status output")
	verify_commitCmd.Flags().BoolP("verbose", "v", false, "print commit contents")
	rootCmd.AddCommand(verify_commitCmd)
}
