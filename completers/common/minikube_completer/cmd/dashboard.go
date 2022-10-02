package cmd

import (
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Access the Kubernetes dashboard running within the minikube cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	dashboardCmd.Flags().Bool("url", false, "Display dashboard URL instead of opening a browser")
	rootCmd.AddCommand(dashboardCmd)
}
