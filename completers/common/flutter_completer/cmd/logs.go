package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show log output for running Flutter apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().BoolP("clear", "c", false, "Clear log history before reading from logs.")
	logsCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	logsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(logsCmd)
}
