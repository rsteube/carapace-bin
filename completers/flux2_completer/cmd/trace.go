package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace in-cluster objects throughout the GitOps delivery pipeline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(traceCmd).Standalone()
	traceCmd.Flags().String("api-version", "", "the Kubernetes object API version, e.g. 'apps/v1'")
	traceCmd.Flags().String("kind", "", "the Kubernetes object kind, e.g. Deployment'")
	rootCmd.AddCommand(traceCmd)
}
