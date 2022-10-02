package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "redis-cli",
	Short: "Redis command line interface",
	Long:  "https://redis.io/docs/manual/cli/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("2", "2", false, "Start session in RESP2 protocol mode")
	rootCmd.Flags().BoolS("3", "3", false, "Start session in RESP3 protocol mode")
	rootCmd.Flags().StringS("D", "D", "", "Delimiter between responses for raw formatting")
	rootCmd.Flags().BoolS("X", "X", false, "Read <tag> argument from STDIN")
	rootCmd.Flags().StringS("a", "a", "", "Password to use when connecting to the server")
	rootCmd.Flags().Bool("askpass", false, "Force user to input password with mask from STDIN")
	rootCmd.Flags().Bool("bigkeys", false, "Sample Redis keys looking for keys with many elements")
	rootCmd.Flags().BoolS("c", "c", false, "Enable cluster mode")
	rootCmd.Flags().String("cacert", "", "CA Certificate file to verify with")
	rootCmd.Flags().String("cacertdir", "", "Directory where trusted CA certificates are stored")
	rootCmd.Flags().String("cert", "", "Client certificate to authenticate with")
	rootCmd.Flags().String("cluster", "", "Cluster Manager command and arguments") // TODO nargs
	rootCmd.Flags().Bool("csv", false, "Output in CSV format")
	rootCmd.Flags().StringS("d", "d", "", "Delimiter between response bulks for raw formatting")
	rootCmd.Flags().BoolS("e", "e", false, "Return exit error code when command execution fails")
	rootCmd.Flags().String("eval", "", "Send an EVAL command using the Lua script at <file>")
	rootCmd.Flags().String("functions-rdb", "", "Like --rdb but only get the functions when getting the RDB dump file")
	rootCmd.Flags().StringS("h", "h", "", "Server hostname")
	rootCmd.Flags().Bool("help", false, "Output this help and exit")
	rootCmd.Flags().Bool("hotkeys", false, "Sample Redis keys looking for hot keys")
	rootCmd.Flags().StringS("i", "i", "", "When -r is used, waits <interval> seconds per command")
	rootCmd.Flags().Bool("insecure", false, "Allow insecure TLS connection by skipping cert validation")
	rootCmd.Flags().String("intrinsic-latency", "", "Run a test to measure intrinsic system latency")
	rootCmd.Flags().Bool("json", false, "Output in JSON format")
	rootCmd.Flags().String("key", "", "Private key file to authenticate with")
	rootCmd.Flags().Bool("latency", false, "Enter a special mode continuously sampling latency")
	rootCmd.Flags().Bool("latency-dist", false, "Shows latency as a spectrum, requires xterm 256 colors")
	rootCmd.Flags().Bool("latency-history", false, "Like --latency but tracking latency changes over time")
	rootCmd.Flags().Bool("ldb", false, "Used with --eval enable the Redis Lua debugger")
	rootCmd.Flags().Bool("ldb-sync-mode", false, "Like --ldb but uses the synchronous Lua debugger")
	rootCmd.Flags().String("lru-test", "", "Simulate a cache workload with an 80-20 distribution")
	rootCmd.Flags().Bool("memkeys", false, "Sample Redis keys looking for keys consuming a lot of memory")
	rootCmd.Flags().String("memkeys-samples", "", "Sample Redis keys looking for keys consuming a lot of memory")
	rootCmd.Flags().StringS("n", "n", "", "Database number")
	rootCmd.Flags().Bool("no-auth-warning", false, "Don't show warning message when using password on command line interface")
	rootCmd.Flags().Bool("no-raw", false, "Force formatted output even when STDOUT is not a tty")
	rootCmd.Flags().StringS("p", "p", "", "Server port")
	rootCmd.Flags().String("pass", "", "Alias of -a for consistency with the new --user option")
	rootCmd.Flags().String("pattern", "", "Keys pattern when using the --scan, --bigkeys or --hotkeys options")
	rootCmd.Flags().Bool("pipe", false, "Transfer raw Redis protocol from stdin to server")
	rootCmd.Flags().String("pipe-timeout", "", "In --pipe mode, abort with error if after sending all data")
	rootCmd.Flags().Bool("quoted-input", false, "Force input to be handled as quoted strings")
	rootCmd.Flags().Bool("quoted-json", false, "Same as --json, but produce ASCII-safe quoted strings, not Unicode")
	rootCmd.Flags().String("quoted-pattern", "", "Same as --pattern, but the specified string can be quoted")
	rootCmd.Flags().StringS("r", "r", "", "Execute specified command N times")
	rootCmd.Flags().Bool("raw", false, "Use raw formatting for replies")
	rootCmd.Flags().String("rdb", "", "Transfer an RDB dump from remote server to local file")
	rootCmd.Flags().Bool("replica", false, "Simulate a replica showing commands received from the master")
	rootCmd.Flags().StringS("s", "s", "", "Server socket")
	rootCmd.Flags().Bool("scan", false, "List all keys using the SCAN command")
	rootCmd.Flags().String("show-pushes", "", "Whether to print RESP3 PUSH messages")
	rootCmd.Flags().String("sni", "", "Server name indication for TLS")
	rootCmd.Flags().Bool("stat", false, "Print rolling stats about server: mem, clients, ...")
	rootCmd.Flags().Bool("tls", false, "Establish a secure TLS connection")
	rootCmd.Flags().String("tls-ciphers", "", "Sets the list of preferred ciphers")
	rootCmd.Flags().String("tls-ciphersuites", "", "Sets the list of preferred ciphersuites")
	rootCmd.Flags().StringS("u", "u", "", "Server URI.")
	rootCmd.Flags().String("user", "", "Used to send ACL style 'AUTH username pass'")
	rootCmd.Flags().Bool("verbose", false, "Verbose mode")
	rootCmd.Flags().Bool("version", false, "Output version and exit")
	rootCmd.Flags().BoolS("x", "x", false, "Read last argument from STDIN")

	// TODO ciphersuites
	// TODO cluster `redis-cli --cluster help`
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cacert":        carapace.ActionFiles(),
		"cacertdir":     carapace.ActionDirectories(),
		"cert":          carapace.ActionFiles(),
		"eval":          carapace.ActionFiles(),
		"functions-rdb": carapace.ActionFiles(),
		"key":           carapace.ActionFiles(),
		"p":             net.ActionPorts(),
		"rdb":           carapace.ActionFiles(),
		"s":             carapace.ActionFiles(),
		"show-pushes":   carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"tls-ciphers": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return ssh.ActionCiphers().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
