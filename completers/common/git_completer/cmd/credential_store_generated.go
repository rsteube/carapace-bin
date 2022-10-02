package cmd

import (
	"github.com/spf13/cobra"
)

var credential_storeCmd = &cobra.Command{
	Use:   "credential-store",
	Short: "Helper to store credentials on disk",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	credential_storeCmd.Flags().String("file", "", "fetch and store credentials in <path>")
	rootCmd.AddCommand(credential_storeCmd)
}
