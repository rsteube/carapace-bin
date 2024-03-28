package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check requirements and installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()
	checkCmd.Flags().StringSlice("components", []string{"source-controller", "kustomize-controller", "helm-controller", "notification-controller"}, "list of components, accepts comma-separated values")
	checkCmd.Flags().StringSlice("components-extra", []string{}, "list of components in addition to those supplied or defaulted, accepts comma-separated values")
	checkCmd.Flags().String("poll-interval", "", "how often the health checker should poll the cluster for the latest state of the resources.")
	checkCmd.Flags().Bool("pre", false, "only run pre-installation checks")
	rootCmd.AddCommand(checkCmd)
}
