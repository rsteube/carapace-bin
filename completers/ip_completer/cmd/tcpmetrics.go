package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tcpmetricsCmd = &cobra.Command{
	Use:   "tcpmetrics",
	Short: "manage TCP Metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcpmetricsCmd).Standalone()

	rootCmd.AddCommand(tcpmetricsCmd)
}
