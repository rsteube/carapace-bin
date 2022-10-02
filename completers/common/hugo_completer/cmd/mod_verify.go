package cmd

import (
	"github.com/spf13/cobra"
)

var mod_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify dependencies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	mod_verifyCmd.Flags().Bool("clean", false, "delete module cache for dependencies that fail verification")
	modCmd.AddCommand(mod_verifyCmd)
}
