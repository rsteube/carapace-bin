package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Docker version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	versionCmd.Flags().String("kubeconfig", "", "Kubernetes config file")
	rootCmd.AddCommand(versionCmd)
}
