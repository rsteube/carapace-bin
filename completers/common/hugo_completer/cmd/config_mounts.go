package cmd

import (
	"github.com/spf13/cobra"
)

var config_mountsCmd = &cobra.Command{
	Use:   "mounts",
	Short: "Print the configured file mounts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(config_mountsCmd)
}
