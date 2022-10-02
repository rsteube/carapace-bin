package cmd

import (
	"github.com/spf13/cobra"
)

var unpack_fileCmd = &cobra.Command{
	Use:   "unpack-file",
	Short: "Creates a temporary file with a blob's contents",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(unpack_fileCmd)
}
