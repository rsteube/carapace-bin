package cmd

import (
	"github.com/spf13/cobra"
)

var mod_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete the Hugo Module cache for the current project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	mod_cleanCmd.Flags().Bool("all", false, "clean entire module cache")
	mod_cleanCmd.Flags().String("pattern", "", "pattern matching module paths to clean (all if not set), e.g. \"**hugo*\"")
	modCmd.AddCommand(mod_cleanCmd)
}
