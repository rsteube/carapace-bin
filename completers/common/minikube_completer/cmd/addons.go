package cmd

import (
	"github.com/spf13/cobra"
)

var addonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Enable or disable a minikube addon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(addonsCmd)
}
