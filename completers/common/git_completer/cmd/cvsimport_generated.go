package cmd

import (
	"github.com/spf13/cobra"
)

var cvsimportCmd = &cobra.Command{
	Use:   "cvsimport",
	Short: "Salvage your data out of another SCM people love to hate",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(cvsimportCmd)
}
