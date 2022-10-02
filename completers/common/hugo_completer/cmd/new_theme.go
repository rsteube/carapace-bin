package cmd

import (
	"github.com/spf13/cobra"
)

var new_themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "Create a new theme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	newCmd.AddCommand(new_themeCmd)
}
