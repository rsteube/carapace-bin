package cmd

import (
	"github.com/spf13/cobra"
)

var mod_npm_packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Experimental: Prepares and writes a composite package.json file for your project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	mod_npmCmd.AddCommand(mod_npm_packCmd)
}
