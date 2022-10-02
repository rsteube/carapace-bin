package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolP("allow-all", "A", false, "Allow all permissions")
	testCmd.Flags().String("allow-env", "", "Allow environment access")
	testCmd.Flags().String("allow-ffi", "", "Allow loading dynamic libraries")
	testCmd.Flags().Bool("allow-hrtime", false, "Allow high resolution time measurement")
	testCmd.Flags().String("allow-net", "", "Allow network access")
	testCmd.Flags().Bool("allow-none", false, "Don't return error code if no test files are found")
	testCmd.Flags().String("allow-read", "", "Allow file system read access")
	testCmd.Flags().String("allow-run", "", "Allow running subprocesses")
	testCmd.Flags().String("allow-write", "", "Allow file system write access")
	testCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	testCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	testCmd.Flags().StringP("config", "c", "", "Load configuration file")
	testCmd.Flags().String("coverage", "", "UNSTABLE: Collect coverage profile data into DIR")
	testCmd.Flags().Bool("doc", false, "UNSTABLE: type check code blocks")
	testCmd.Flags().String("fail-fast", "", "Stop after N errors. Defaults to stopping after first failure.")
	testCmd.Flags().String("filter", "", "Run tests with this string or pattern in the test name")
	testCmd.Flags().String("ignore", "", "Ignore files")
	testCmd.Flags().String("import-map", "", "Load import map file")
	testCmd.Flags().String("inspect", "", "Activate inspector on host:port (default: 127.0.0.1:9229)")
	testCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script")
	testCmd.Flags().StringP("jobs", "j", "", "Number of parallel workers, defaults to # of CPUs when no value is provided. Defaults to 1 when the option is not present.")
	testCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	testCmd.Flags().String("lock", "", "Check the specified lock file")
	testCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	testCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	testCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	testCmd.Flags().Bool("no-run", false, "Cache test modules, but don't run tests")
	testCmd.Flags().Bool("prompt", false, "Fallback to prompt if required permission wasn't passed")
	testCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	testCmd.Flags().String("seed", "", "Seed Math.random()")
	testCmd.Flags().String("shuffle", "", "(UNSTABLE): Shuffle the order in which the tests are run")
	testCmd.Flags().String("unsafely-ignore-certificate-errors", "", "DANGER: Disables verification of TLS certificates")
	testCmd.Flags().String("v8-flags", "", "Set V8 command line options (for help: --v8-flags=--help")
	testCmd.Flags().Bool("watch", false, "UNSTABLE: Watch for file changes and restart process automatically")
	rootCmd.AddCommand(testCmd)

	testCmd.Flag("allow-env").NoOptDefVal = " "
	testCmd.Flag("allow-ffi").NoOptDefVal = " "
	testCmd.Flag("allow-net").NoOptDefVal = " "
	testCmd.Flag("allow-read").NoOptDefVal = " "
	testCmd.Flag("allow-run").NoOptDefVal = " "
	testCmd.Flag("allow-write").NoOptDefVal = " "
	testCmd.Flag("coverage").NoOptDefVal = " "
	testCmd.Flag("fail-fast").NoOptDefVal = " "
	testCmd.Flag("ignore").NoOptDefVal = " "
	testCmd.Flag("inspect").NoOptDefVal = " "
	testCmd.Flag("inspect-brk").NoOptDefVal = " "
	testCmd.Flag("reload").NoOptDefVal = " "
	testCmd.Flag("shuffle").NoOptDefVal = " "
	testCmd.Flag("unsafely-ignore-certificate-errors").NoOptDefVal = " "
	testCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(),
		"coverage":    carapace.ActionDirectories(),
		"import-map":  carapace.ActionFiles(),
		"inspect":     action.ActionHostPort(),
		"inspect-brk": action.ActionHostPort(),
		"lock":        carapace.ActionFiles(),
	})

	carapace.Gen(testCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
