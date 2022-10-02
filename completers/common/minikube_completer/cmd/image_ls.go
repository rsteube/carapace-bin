package cmd

import (
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	imageCmd.AddCommand(image_lsCmd)
}
