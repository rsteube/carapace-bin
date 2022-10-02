package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "node",
	Short: "server-side JavaScript runtime",
	Long:  "https://nodejs.org/en/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("abort-on-uncaught-exception", false, "Aborting instead of exiting causes a core file to be generated for analysis.")
	rootCmd.Flags().BoolP("check", "c", false, "Check the script's syntax without executing it.")
	rootCmd.Flags().Bool("completion-bash", false, "Print source-able bash completion script for Node.js.")
	rootCmd.Flags().String("conditions", "", "Use custom conditional exports conditions.")
	rootCmd.Flags().Bool("cpu-prof", false, "Start the V8 CPU profiler on start up")
	rootCmd.Flags().String("cpu-prof-dir", "", "The directory where the CPU profiles will be placed.")
	rootCmd.Flags().String("cpu-prof-interval", "", "The sampling interval in microseconds for the CPU profiles.")
	rootCmd.Flags().String("cpu-prof-name", "", "File name of the V8 CPU profile.")
	rootCmd.Flags().String("diagnostic-dir", "", "Set the directory for all diagnostic output files.")
	rootCmd.Flags().String("disable-proto", "", "Disable the `Object.prototype.__proto__` property.")
	rootCmd.Flags().Bool("disallow-code-generation-from-strings", false, "Make built-in language features that generate code from strings throw an exception instead.")
	rootCmd.Flags().Bool("enable-fips", false, "Enable FIPS-compliant crypto at startup.")
	rootCmd.Flags().Bool("enable-source-maps", false, "Enable Source Map V3 support for stack traces.")
	rootCmd.Flags().StringP("eval", "e", "", "Evaluate string as JavaScript.")
	rootCmd.Flags().Bool("experimental-import-meta-resolve", false, "Enable experimental ES modules support for import.meta.resolve().")
	rootCmd.Flags().Bool("experimental-json-modules", false, "Enable experimental JSON interop support for the ES Module loader.")
	rootCmd.Flags().String("experimental-loader", "", "Specify the module to use as a custom module loader.")
	rootCmd.Flags().String("experimental-policy", "", "Use the specified file as a security policy.")
	rootCmd.Flags().Bool("experimental-repl-await", false, "Enable experimental top-level await keyword support in REPL.")
	rootCmd.Flags().String("experimental-specifier-resolution", "", "Select extension resolution algorithm for ES Modules.")
	rootCmd.Flags().Bool("experimental-vm-modules", false, "Enable experimental ES module support in VM module.")
	rootCmd.Flags().Bool("experimental-wasi-unstable-preview1", false, "Enable experimental WebAssembly System Interface support.")
	rootCmd.Flags().Bool("experimental-wasm-modules", false, "Enable experimental WebAssembly module support.")
	rootCmd.Flags().Bool("force-context-aware", false, "Disable loading native addons that are not context-aware.")
	rootCmd.Flags().Bool("force-fips", false, "Force FIPS-compliant crypto on startup.")
	rootCmd.Flags().Bool("frozen-intrinsics", false, "Enable experimental frozen intrinsics support.")
	rootCmd.Flags().Bool("heap-prof", false, "Start the V8 heap profiler on start up.")
	rootCmd.Flags().String("heap-prof-dir", "", "The directory where the heap profiles will be placed.")
	rootCmd.Flags().Bool("heap-prof-interval", false, "The average sampling interval in bytes for the heap profiles.")
	rootCmd.Flags().String("heap-prof-name", "", "File name of the V8 heap profile.")
	rootCmd.Flags().String("heapsnapshot-near-heap-limit", "", "Generate heap snapshot when the V8 heap usage is approaching the heap limit.")
	rootCmd.Flags().String("heapsnapshot-signal", "", "Generate heap snapshot on specified signal.")
	rootCmd.Flags().BoolP("help", "h", false, "Print command-line options.")
	rootCmd.Flags().String("icu-data-dir", "", "Specify ICU data load path.")
	rootCmd.Flags().String("input-type", "", "Set the module resolution type for input via --eval, --print or STDIN.")
	rootCmd.Flags().Bool("insecure-http-parser", false, "Use an insecure HTTP parser that accepts invalid HTTP headers.")
	rootCmd.Flags().String("inspect", "", "Activate inspector on host:port.")
	rootCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script.")
	rootCmd.Flags().String("inspect-port", "", "Set the host:port to be used when the inspector is activated.")
	rootCmd.Flags().String("inspect-publish-uid", "", "Specify how the inspector WebSocket URL is exposed.")
	rootCmd.Flags().BoolP("interactive", "i", false, "Open the REPL even if stdin does not appear to be a terminal.")
	rootCmd.Flags().Bool("jitless", false, "Disable runtime allocation of executable memory.")
	rootCmd.Flags().String("max-http-header-size", "", "Specify the maximum size of HTTP headers in bytes.")
	rootCmd.Flags().Bool("napi-modules", false, "This option is a no-op.")
	rootCmd.Flags().Bool("no-deprecation", false, "Silence deprecation warnings.")
	rootCmd.Flags().Bool("no-force-async-hooks-checks", false, "Disable runtime checks for `async_hooks`.")
	rootCmd.Flags().Bool("no-warnings", false, "Silence all process warnings (including deprecations).")
	rootCmd.Flags().Bool("node-memory-debug", false, "Enable extra debug checks for memory leaks in Node.js internals.")
	rootCmd.Flags().String("openssl-config", "", "Load an OpenSSL configuration file on startup.")
	rootCmd.Flags().Bool("pending-deprecation", false, "Emit pending deprecation warnings.")
	rootCmd.Flags().String("policy-integrity", "", "Instructs Node.js to error prior to running any code if the policy does not have the specified integrity.")
	rootCmd.Flags().Bool("preserve-symlinks", false, "Instructs the module loader to preserve symbolic links when resolving and caching modules other than the main module.")
	rootCmd.Flags().Bool("preserve-symlinks-main", false, "Instructs the module loader to preserve symbolic links when resolving and caching the main module.")
	rootCmd.Flags().StringP("print", "p", "", "Identical to -e, but prints the result.")
	rootCmd.Flags().Bool("prof", false, "Generate V8 profiler output.")
	rootCmd.Flags().Bool("prof-process", false, "Process V8 profiler output generated using the V8 option --prof.")
	rootCmd.Flags().String("redirect-warnings", "", "Write process warnings to the given file instead of printing to stderr.")
	rootCmd.Flags().Bool("report-compact", false, "Write diagnostic reports in a compact format, single-line JSON.")
	rootCmd.Flags().Bool("report-dir", false, "Location at which the diagnostic report will be generated.")
	rootCmd.Flags().Bool("report-directory", false, "Location at which the diagnostic report will be generated.")
	rootCmd.Flags().Bool("report-filename", false, "Name of the file to which the diagnostic report will be written.")
	rootCmd.Flags().Bool("report-on-fatalerror", false, "Enables the diagnostic report to be triggered on fatal errors.")
	rootCmd.Flags().Bool("report-on-signal", false, "Enables diagnostic report to be generated upon receiving the specified signal.")
	rootCmd.Flags().Bool("report-signal", false, "Sets or resets the signal for diagnostic report generation.")
	rootCmd.Flags().Bool("report-uncaught-exception", false, "Enables diagnostic report to be generated on un-caught exceptions.")
	rootCmd.Flags().StringP("require", "r", "", "Preload the specified module at startup.")
	rootCmd.Flags().String("secure-heap", "", "Specify the size of the OpenSSL secure heap.")
	rootCmd.Flags().String("secure-heap-min", "", "Specify the minimum allocation from the OpenSSL secure heap.")
	rootCmd.Flags().Bool("throw-deprecation", false, "Throw errors for deprecations.")
	rootCmd.Flags().String("title", "", "Specify process.title on startup.")
	rootCmd.Flags().String("tls-cipher-list", "", "Specify an alternative default TLS cipher list.")
	rootCmd.Flags().String("tls-keylog", "", "Log TLS key material to a file.")
	rootCmd.Flags().Bool("tls-max-v1.2", false, "Set default  maxVersion to 'TLSv1.2'.")
	rootCmd.Flags().Bool("tls-max-v1.3", false, "Set default  maxVersion to 'TLSv1.3'.")
	rootCmd.Flags().Bool("tls-min-v1.0", false, "Set default minVersion to 'TLSv1'.")
	rootCmd.Flags().Bool("tls-min-v1.1", false, "Set default minVersion to 'TLSv1.1'.")
	rootCmd.Flags().Bool("tls-min-v1.2", false, "Set default minVersion to 'TLSv1.2'.")
	rootCmd.Flags().Bool("tls-min-v1.3", false, "Set default minVersion to 'TLSv1.3'.")
	rootCmd.Flags().Bool("trace-atomics-wait", false, "Print short summaries of calls to Atomics.wait().")
	rootCmd.Flags().Bool("trace-deprecation", false, "Print stack traces for deprecations.")
	rootCmd.Flags().String("trace-event-categories", "", "A comma-separated list of categories that should be traced when trace event tracing is enabled.")
	rootCmd.Flags().String("trace-event-file-pattern", "", "Template string specifying the filepath for the trace event data.")
	rootCmd.Flags().Bool("trace-events-enabled", false, "Enable the collection of trace event tracing information.")
	rootCmd.Flags().Bool("trace-exit", false, "Prints a stack trace whenever an environment is exited proactively.")
	rootCmd.Flags().Bool("trace-sigint", false, "Prints a stack trace on SIGINT.")
	rootCmd.Flags().Bool("trace-sync-io", false, "Print a stack trace whenever synchronous I/O is detected after the first turn of the event loop.")
	rootCmd.Flags().Bool("trace-tls", false, "Prints TLS packet trace information to stderr.")
	rootCmd.Flags().Bool("trace-uncaught", false, "Print stack traces for uncaught exceptions.")
	rootCmd.Flags().Bool("trace-warnings", false, "Print stack traces for process warnings (including deprecations).")
	rootCmd.Flags().Bool("track-heap-objects", false, "Track heap object allocations for heap snapshots.")
	rootCmd.Flags().String("unhandled-rejections", "", "Define the behavior for unhandled rejections.")
	rootCmd.Flags().Bool("use-bundled-ca", false, "Use bundled Mozilla CA store as supplied by current Node.js version or use OpenSSL's default CA store.")
	rootCmd.Flags().String("use-largepages", "", "Re-map the Node.js static code to large memory pages at startup.")
	rootCmd.Flags().Bool("use-openssl-ca", false, "Use bundled Mozilla CA store as supplied by current Node.js version or use OpenSSL's default CA store.")
	rootCmd.Flags().Bool("v8-options", false, "Print V8 command-line options.")
	rootCmd.Flags().String("v8-pool-size", "", "Set V8's thread pool size which will be used to allocate background jobs.")
	rootCmd.Flags().BoolP("version", "v", false, "Print node's version.")
	rootCmd.Flags().Bool("zero-fill-buffers", false, "Automatically zero-fills all newly allocated Buffer and SlowBuffer instances.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpu-prof-dir":        carapace.ActionDirectories(),
		"diagnostic-dir":      carapace.ActionDirectories(),
		"disable-proto":       carapace.ActionValues("delete", "throw"),
		"heapsnapshot-signal": ps.ActionKillSignals(),
		"icu-data-dir":        carapace.ActionDirectories(),
		"inspect-publish-uid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("http", "stderr").Invoke(c).Filter(c.Parts).ToA()
		}),
		"openssl-config":       carapace.ActionFiles(),
		"redirect-warnings":    carapace.ActionFiles(),
		"report-dir":           carapace.ActionDirectories(),
		"report-directory":     carapace.ActionDirectories(),
		"tls-keylog":           carapace.ActionFiles(),
		"unhandled-rejections": carapace.ActionValues("strict", "warn", "none"),
		"use-largepages":       carapace.ActionValues("off", "on", "silent"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles().Invoke(c).Merge(carapace.ActionValues("inspect").Invoke(c)).ToA()
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if c.Args[0] == "inspect" {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)
}
