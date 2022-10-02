package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dlv",
	Short: "Delve is a debugger for the Go programming language.",
	Long:  "https://github.com/go-delve/delve",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("accept-multiclient", false, "Allows a headless server to accept multiple client connections via JSON-RPC or DAP.")
	rootCmd.PersistentFlags().Bool("allow-non-terminal-interactive", false, "Allows interactive sessions of Delve that don't have a terminal as stdin, stdout and stderr")
	rootCmd.PersistentFlags().Int("api-version", 1, "Selects JSON-RPC API version when headless. New clients should use v2. Can be reset via RPCServer.SetApiVersion. See Documentation/api/json-rpc/README.md.")
	rootCmd.PersistentFlags().String("backend", "default", "Backend selection (see 'dlv help backend').")
	rootCmd.PersistentFlags().String("build-flags", "", "Build flags, to be passed to the compiler. For example: --build-flags=\"-tags=integration -mod=vendor -cover -v\"")
	rootCmd.PersistentFlags().Bool("check-go-version", true, "Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve.")
	rootCmd.PersistentFlags().Bool("disable-aslr", false, "Disables address space randomization")
	rootCmd.PersistentFlags().Bool("headless", false, "Run debug server only, in headless mode. Server will accept both JSON-RPC or DAP client connections.")
	rootCmd.Flags().BoolP("help", "h", false, "help for dlv")
	rootCmd.PersistentFlags().String("init", "", "Init file, executed by the terminal client.")
	rootCmd.PersistentFlags().StringP("listen", "l", "127.0.0.1:0", "Debugging server listen address.")
	rootCmd.PersistentFlags().Bool("log", false, "Enable debugging server logging.")
	rootCmd.PersistentFlags().String("log-dest", "", "Writes logs to the specified file or file descriptor (see 'dlv help log').")
	rootCmd.PersistentFlags().String("log-output", "", "Comma separated list of components that should produce debug output (see 'dlv help log')")
	rootCmd.PersistentFlags().Bool("only-same-user", true, "Only connections from the same user that started this instance of Delve are allowed to connect.")
	rootCmd.PersistentFlags().StringArrayP("redirect", "r", []string{}, "Specifies redirect rules for target process (see 'dlv help redirect')")
	rootCmd.PersistentFlags().String("wd", "", "Working directory for running the program.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backend": carapace.ActionValuesDescribed(
			"default", "Uses lldb on macOS, native everywhere else.",
			"native", "Native backend.",
			"lldb", "Uses lldb-server or debugserver.",
			"rr", "Uses mozilla rr (https://github.com/mozilla/rr)",
		),
		"init":     carapace.ActionFiles(),
		"log-dest": carapace.ActionFiles(),
		"log-output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValuesDescribed(
				"debugger", "Log debugger commands",
				"gdbwire", "Log connection to gdbserial backend",
				"lldbout", "Copy output from debugserver/lldb to standard output",
				"debuglineerr", "Log recoverable errors reading .debug_line",
				"rpc", "Log all RPC messages",
				"dap", "Log all DAP messages",
				"fncall", "Log function call protocol",
				"minidump", "Log minidump loading",
			).Invoke(c).Filter(c.Parts).ToA()
		}),
		"redirect": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("stdin", "stdout", "stderr").Invoke(c).Suffix(":").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
