package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/dict_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dict",
	Short: "Query a dictd server for the definition of a word",
	Long:  "https://linux.die.net/man/1/dict",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "forces dict to use IPv4 addresses only.")
	rootCmd.Flags().BoolS("6", "6", false, "forces dict to use IPv6 addresses only.")
	rootCmd.Flags().String("client", "", "additional text for client command")
	rootCmd.Flags().StringP("config", "c", "", "specify configuration file")
	rootCmd.Flags().StringP("database", "d", "", "select a database to search")
	rootCmd.Flags().BoolP("dbs", "D", false, "show available databases")
	rootCmd.Flags().String("debug", "", "set debugging flag")
	rootCmd.Flags().BoolP("formatted", "f", false, "use strict tabbed format of output")
	rootCmd.Flags().Bool("help", false, "display this help")
	rootCmd.Flags().StringP("host", "h", "", "specify server")
	rootCmd.Flags().StringP("info", "i", "", "show information about a database")
	rootCmd.Flags().StringP("key", "k", "", "shared secret for authentication")
	rootCmd.Flags().BoolP("license", "L", false, "display copyright and license information")
	rootCmd.Flags().BoolP("match", "m", false, "match instead of define")
	rootCmd.Flags().BoolP("mime", "M", false, "send OPTION MIME command if server supports it")
	rootCmd.Flags().BoolP("noauth", "a", false, "disable authentication")
	rootCmd.Flags().BoolP("nocorrect", "C", false, "disable attempted spelling correction")
	rootCmd.Flags().String("pipesize", "", "specify buffer size for pipelining (256)")
	rootCmd.Flags().StringP("port", "p", "", "specify port")
	rootCmd.Flags().BoolP("raw", "r", false, "trace raw transaction")
	rootCmd.Flags().BoolP("serverhelp", "H", false, "show server help")
	rootCmd.Flags().BoolP("serverinfo", "I", false, "show information about the server")
	rootCmd.Flags().StringP("strategy", "s", "", "strategy for matching or defining")
	rootCmd.Flags().BoolP("strats", "S", false, "show available search strategies")
	rootCmd.Flags().StringP("user", "u", "", "username for authentication")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "display version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"database": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDatabases(rootCmd.Flag("host").Value.String())
		}),
		"host": net.ActionHosts(),
		"port": net.ActionPorts(),
		"strategy": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionStrategies(rootCmd.Flag("host").Value.String())
		}),
		"user": os.ActionUsers(),
	})
}
