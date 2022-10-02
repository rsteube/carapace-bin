package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh",
	Short: "OpenSSH remote login client",
	Long:  "https://linux.die.net/man/1/ssh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Forces ssh to use IPv4 addresses only")
	rootCmd.Flags().BoolS("6", "6", false, "Forces ssh to use IPv6 addresses only")
	rootCmd.Flags().BoolS("A", "A", false, "Enables forwarding of connections from an authentication agent")
	rootCmd.Flags().StringS("B", "B", "", "Bind to given address before attempting to connect")
	rootCmd.Flags().BoolS("C", "C", false, "Requests compression of all data")
	rootCmd.Flags().StringS("D", "D", "", "Specifies a local “dynamic” application-level port forwarding")
	rootCmd.Flags().StringS("E", "E", "", "Append debug logs to log_file instead of standard error")
	rootCmd.Flags().StringS("F", "F", "", "Specifies an alternative per-user configuration file")
	rootCmd.Flags().BoolS("G", "G", false, "Causes ssh to print its configuration after evaluating Host and Match blocks and exit")
	rootCmd.Flags().StringS("I", "I", "", "Specify the PKCS#11 shared library ssh should use")
	rootCmd.Flags().StringS("J", "J", "", "Connect to the target host by using given jump host")
	rootCmd.Flags().BoolS("K", "K", false, "Enables GSSAPI-based authentication and forwarding of GSSAPI credentials")
	rootCmd.Flags().StringS("L", "L", "", "Specifies that connections are to be forwarded")
	rootCmd.Flags().BoolS("M", "M", false, "Places the ssh client into “master” mode for connection sharing")
	rootCmd.Flags().BoolS("N", "N", false, "Do not execute a remote command.  This is useful for just forwarding ports.")
	rootCmd.Flags().StringS("O", "O", "", "Control an active connection multiplexing master process")
	rootCmd.Flags().StringS("Q", "Q", "", "Queries ssh for the algorithms supported")
	rootCmd.Flags().StringS("R", "R", "", "Specifies that connections ar to be forwarded")
	rootCmd.Flags().StringS("S", "S", "", "Specifies the location of a control socket")
	rootCmd.Flags().BoolS("T", "T", false, "Disable pseudo-terminal allocation")
	rootCmd.Flags().BoolS("V", "V", false, "Display the version number and exit")
	rootCmd.Flags().StringS("W", "W", "", "Requests that standard input and output on the client be forwarded")
	rootCmd.Flags().BoolS("X", "X", false, "Enables X11 forwarding")
	rootCmd.Flags().BoolS("Y", "Y", false, "Enables trusted X11 forwarding")
	rootCmd.Flags().BoolS("a", "a", false, "Disables forwarding of the authentication agent connection.")
	rootCmd.Flags().StringS("b", "b", "", "Use bind_address on the local machine as the source address of the connection")
	rootCmd.Flags().StringS("c", "c", "", "Selects the cipher specification for encrypting the session")
	rootCmd.Flags().StringS("e", "e", "", "Sets the escape character for sessions with a pty")
	rootCmd.Flags().BoolS("f", "f", false, "Requests ssh to go to background just before command execution")
	rootCmd.Flags().BoolS("g", "g", false, "Allows remote hosts to connect to local forwarded ports")
	rootCmd.Flags().StringS("i", "i", "", "Selects a file from which the identity (private key) is read")
	rootCmd.Flags().BoolS("k", "k", false, "Disables forwarding of GSSAPI credentials to the server")
	rootCmd.Flags().StringS("l", "l", "", "Specifies the user to log in as on the remote machine")
	rootCmd.Flags().StringS("m", "m", "", "A comma-separated list of MAC algorithms")
	rootCmd.Flags().BoolS("n", "n", false, "Redirects stdin from /dev/null")
	rootCmd.Flags().StringS("o", "o", "", "Can be used to give options in the format used in the configuration file")
	rootCmd.Flags().StringS("p", "p", "", "Port to connect to on the remote host")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("s", "s", false, "May be used to request invocation of a subsystem on the remote system")
	rootCmd.Flags().BoolS("t", "t", false, "Force pseudo-terminal allocation")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
	rootCmd.Flags().StringS("w", "w", "", "Requests tunnel device forwarding")
	rootCmd.Flags().BoolS("x", "x", false, "Disables X11 forwarding")
	rootCmd.Flags().BoolS("y", "y", false, "Send log information using the syslog system module")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO complete bindings
		// "B"
		// "b"
		// "D"
		// "J"
		// "L"
		// ";"
		// "m"
		// "p"
		// "R"
		// "W"
		// "w"
		"E": carapace.ActionFiles(),
		"F": carapace.ActionFiles(),
		"I": carapace.ActionFiles(),
		"O": carapace.ActionValuesDescribed(
			"check", "check that the master process is running",
			"forward", "request forwardings without command execution",
			"cancel", "cancel forwardings",
			"exit", "request the master to exit",
			"stop", "request the master to stop accepting further multiplexing requests",
		),
		"Q": ActionQueryOptions(),
		"S": carapace.ActionFiles(),
		"c": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"i": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"o": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionOptions()
		}),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.Contains(c.CallbackValue, "@") {
				prefix := strings.SplitN(c.CallbackValue, "@", 2)[0]
				return net.ActionHosts().Invoke(c).Prefix(prefix + "@").ToA()
			} else {
				return net.ActionHosts()
			}
		}),
	)
}

func ActionQueryOptions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cipher", "supported symmetric ciphers",
		"cipher-auth", "supported symmetric ciphers that support authenticated encryption",
		"help", "supported query terms for use with the -Q flag",
		"mac", "supported message integrity codes",
		"kex", "key exchange algorithms",
		"key", "key types",
		"key-cert", "certificate key types",
		"key-plain", "non-certificate key types",
		"key-sig", "all key types and signature algorithms",
		"protocol-version", "supported SSH protocol versions",
		"sig", "supported signature algorithms",
	)
}
