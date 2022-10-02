package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mosh",
	Short: "mobile shell with roaming and intelligent local echo",
	Long:  "https://mosh.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "use IPv4 only")
	rootCmd.Flags().BoolS("6", "6", false, "use IPv6 only")
	rootCmd.Flags().BoolS("a", "a", false, "use local echo even on fast links")
	rootCmd.Flags().String("bind-server", "", "ask the server to reply from an IP address")
	rootCmd.Flags().String("client", "", "mosh client on local machine")
	rootCmd.Flags().String("experimental-remote-ip", "", "select the method for discovering the remote IP address")
	rootCmd.Flags().String("family", "", "net family ")
	rootCmd.Flags().Bool("help", false, "show help message")
	rootCmd.Flags().Bool("local", false, "run mosh-server locally without using ssh")
	rootCmd.Flags().BoolS("n", "n", false, "never use local echo")
	rootCmd.Flags().Bool("no-init", false, "do not send terminal initialization string")
	rootCmd.Flags().Bool("no-ssh-pty", false, "do not allocate a pseudo tty on ssh connection")
	rootCmd.Flags().StringP("port", "p", "", "server-side UDP port or range")
	rootCmd.Flags().String("predict", "", "predict setting")
	rootCmd.Flags().String("server", "", "mosh server on remote machine")
	rootCmd.Flags().String("ssh", "", "ssh command to run when setting up session")
	rootCmd.Flags().Bool("version", false, "version and copyright information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bind-server":            carapace.ActionValues("ssh", "any", "IP"),
		"client":                 carapace.ActionFiles(),
		"experimental-remote-ip": carapace.ActionValues("local", "remote", "proxy"),
		"family": carapace.ActionValuesDescribed(
			"inet", "use IPv4 only",
			"inet6", "use IPv6 only",
			"auto", "autodetect network type for single-family hosts only",
			"all", "try all network types",
			"prefer-inet", "use all network types, but try IPv4 first",
			"prefer-inet6", "use all network types, but try IPv6 first",
		).StyleF(style.ForKeyword),
		"port": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionPorts()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"predict": carapace.ActionValuesDescribed(
			"adaptive", "local echo for slower links",
			"always", "use local echo even on fast links",
			"never", "never use local echo",
			"experimental", "aggressively echo even when incorrect",
		).StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					users := os.ActionUsers().Invoke(c).Suffix("@")
					hosts := net.ActionHosts().Invoke(c)
					return users.Merge(hosts).ToA()
				case 1:
					return net.ActionHosts()
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
