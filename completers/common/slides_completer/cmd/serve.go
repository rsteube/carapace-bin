package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start an SSH server to run slide",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().BoolP("help", "h", false, "help for serve")
	serveCmd.Flags().String("host", "", "Server host to bind to (default \"localhost\")")
	serveCmd.Flags().String("keyPath", "", "Server private key path (default \"slides\")")
	serveCmd.Flags().String("port", "", "Server port to bind to (default 53531)")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"keyPath": carapace.ActionFiles(),
		"port":    net.ActionPorts(),
	})

	carapace.Gen(serveCmd).PositionalCompletion(
		carapace.ActionFiles(".md"),
	)
}
