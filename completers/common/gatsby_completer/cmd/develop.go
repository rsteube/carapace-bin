package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "Start development server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developCmd).Standalone()

	developCmd.Flags().String("ca-file", "", "Custom HTTPS CA certificate file.")
	developCmd.Flags().StringP("cert-file", "c", "", "Custom HTTPS cert file.")
	developCmd.Flags().Bool("graphql-tracing", false, "Trace every graphql resolver.")
	developCmd.Flags().StringP("host", "H", "", "Set host.")
	developCmd.Flags().BoolP("https", "S", false, "Use HTTPS.")
	developCmd.Flags().String("inspect", "", "Opens a port for debugging.")
	developCmd.Flags().String("inspect-brk", "", "Opens a port for debugging.")
	developCmd.Flags().StringP("key-file", "k", "", "Custom HTTPS key file.")
	developCmd.Flags().BoolP("open", "o", false, "Open the site in your (default) browser for you.")
	developCmd.Flags().String("open-tracing-config-file", "", "Tracer configuration file.")
	developCmd.Flags().StringP("port", "p", "", "Set port.")
	rootCmd.AddCommand(developCmd)

	carapace.Gen(developCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":                  carapace.ActionFiles(),
		"cert-file":                carapace.ActionFiles(),
		"key-file":                 carapace.ActionFiles(),
		"open-tracing-config-file": carapace.ActionFiles(),
		"port":                     net.ActionPorts(),
	})
}
