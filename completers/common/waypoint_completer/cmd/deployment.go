package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Deployment creation and management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deploymentCmd).Standalone()

	rootCmd.AddCommand(deploymentCmd)
}
