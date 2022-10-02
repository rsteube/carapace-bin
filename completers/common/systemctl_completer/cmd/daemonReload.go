package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var daemonReloadCmd = &cobra.Command{
	Use:   "daemon-reload",
	Short: "Reload systemd manager configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonReloadCmd).Standalone()

	rootCmd.AddCommand(daemonReloadCmd)
}
