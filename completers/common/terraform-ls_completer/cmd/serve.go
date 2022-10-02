package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Language Server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "file into which to write CPU profile")
	serveCmd.Flags().StringS("log-file", "log-file", "", "path to a file to log into with support for variables")
	serveCmd.Flags().StringS("memprofile", "memprofile", "", "file into which to write memory profile")
	serveCmd.Flags().StringS("port", "port", "", "port number to listen on")
	serveCmd.Flags().StringS("req-concurrency", "req-concurrency", "", "number of RPC requests to process concurrently")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"log-file":   carapace.ActionFiles(),
		"memprofile": carapace.ActionFiles(),
		"port":       net.ActionPorts(),
	})
}
