package cmd

import (
	"github.com/spf13/cobra"
)

var fetch_packCmd = &cobra.Command{
	Use:   "fetch-pack",
	Short: "Receive missing objects from another repository",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(fetch_packCmd)
}
