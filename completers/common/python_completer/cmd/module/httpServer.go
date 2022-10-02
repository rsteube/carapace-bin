package module

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

func ActionInvokeHttpServer() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cmd := &cobra.Command{
			Use:   "http.server",
			Short: "simple http server",
			Run:   func(cmd *cobra.Command, args []string) {},
		}

		carapace.Gen(cmd).Standalone()

		cmd.Flags().StringP("bind", "b", "", "Specify alternate bind address")
		cmd.Flags().Bool("cgi", false, "Run as CGI Server")
		cmd.Flags().StringP("directory", "d", "", "Specify alternative directory")
		cmd.Flags().BoolP("help", "h", false, "show this help message and exit")

		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"directory": carapace.ActionDirectories(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			net.ActionPorts(),
		)

		return carapace.ActionExecute(cmd)
	})
}
