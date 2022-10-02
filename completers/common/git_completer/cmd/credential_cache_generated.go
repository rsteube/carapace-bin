package cmd

import (
	"github.com/spf13/cobra"
)

var credential_cacheCmd = &cobra.Command{
	Use:   "credential-cache",
	Short: "Helper to temporarily store passwords in memory",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	credential_cacheCmd.Flags().String("socket", "", "path of cache-daemon socket")
	credential_cacheCmd.Flags().String("timeout", "", "number of seconds to cache credentials")
	rootCmd.AddCommand(credential_cacheCmd)
}
