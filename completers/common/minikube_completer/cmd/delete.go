package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a local Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	deleteCmd.Flags().Bool("all", false, "Set flag to delete all profiles")
	deleteCmd.Flags().Bool("purge", false, "Set this flag to delete the '.minikube' folder from your user directory.")
	rootCmd.AddCommand(deleteCmd)
}
