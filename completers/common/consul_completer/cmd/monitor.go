package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Stream logs from a Consul agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorCmd).Standalone()
	addClientFlags(monitorCmd)
	monitorCmd.Flags().Bool("log-json", false, "Output logs in JSON format.")
	monitorCmd.Flags().String("log-level", "", "Log level of the agent.")

	rootCmd.AddCommand(monitorCmd)

	carapace.Gen(monitorCmd).FlagCompletion(carapace.ActionMap{
		"log-level": carapace.ActionValues("trace", "debug", "info", "warn", "err").StyleF(style.ForLogLevel),
	})
}
