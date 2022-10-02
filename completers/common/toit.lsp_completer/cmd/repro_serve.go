package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var repro_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repro_serveCmd).Standalone()
	repro_serveCmd.Flags().Bool("json", false, "output format as json")
	repro_serveCmd.Flags().Int("port", 0, "port to use for file-server")
	reproCmd.AddCommand(repro_serveCmd)

	carapace.Gen(repro_serveCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})

	carapace.Gen(repro_serveCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
