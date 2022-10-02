package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run benchmarks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchCmd).Standalone()

	benchCmd.Flags().BoolP("allow-all", "A", false, "Allow all permissions")
	benchCmd.Flags().String("allow-env", "", "Allow environment access")
	benchCmd.Flags().String("allow-ffi", "", "Allow loading dynamic libraries")
	benchCmd.Flags().Bool("allow-hrtime", false, "Allow high resolution time measurement")
	benchCmd.Flags().String("allow-net", "", "Allow network access")
	benchCmd.Flags().String("allow-read", "", "Allow file system read access")
	benchCmd.Flags().String("allow-run", "", "Allow running subprocesses")
	benchCmd.Flags().String("allow-write", "", "Allow file system write access")
	benchCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	benchCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	benchCmd.Flags().String("check", "", "Type check modules")
	benchCmd.Flags().Bool("compat", false, "UNSTABLE: Node compatibility mode.")
	benchCmd.Flags().StringP("config", "c", "", "Specify the configuration file")
	benchCmd.Flags().String("filter", "", "Run benchmarks with this string or pattern in the bench name")
	benchCmd.Flags().BoolP("help", "h", false, "Print help information")
	benchCmd.Flags().String("ignore", "", "Ignore files")
	benchCmd.Flags().String("import-map", "", "Load import map file")
	benchCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	benchCmd.Flags().String("lock", "", "Check the specified lock file")
	benchCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	benchCmd.Flags().String("no-check", "", "Skip type checking modules")
	benchCmd.Flags().Bool("no-clear-screen", false, "Do not clear terminal screen when under watch mode")
	benchCmd.Flags().Bool("no-config", false, "Disable automatic loading of the configuration file.")
	benchCmd.Flags().Bool("no-prompt", false, "Always throw if required permission wasn't passed")
	benchCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	benchCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	benchCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	benchCmd.Flags().String("seed", "", "Set the random number generator seed")
	benchCmd.Flags().String("unsafely-ignore-certificate-errors", "", "DANGER: Disables verification of TLS certificates")
	benchCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	benchCmd.Flags().String("v8-flags", "", "Set V8 command line options")
	benchCmd.Flags().Bool("watch", false, "UNSTABLE: Watch for file changes and restart automatically")
	rootCmd.AddCommand(benchCmd)

	benchCmd.Flag("allow-env").NoOptDefVal = " "
	benchCmd.Flag("allow-ffi").NoOptDefVal = " "
	benchCmd.Flag("allow-net").NoOptDefVal = " "
	benchCmd.Flag("allow-read").NoOptDefVal = " "
	benchCmd.Flag("allow-run").NoOptDefVal = " "
	benchCmd.Flag("allow-write").NoOptDefVal = " "
	benchCmd.Flag("check").NoOptDefVal = " "
	benchCmd.Flag("ignore").NoOptDefVal = " "
	benchCmd.Flag("no-check").NoOptDefVal = " "
	benchCmd.Flag("reload").NoOptDefVal = " "
	benchCmd.Flag("unsafely-ignore-certificate-errors").NoOptDefVal = " "
	benchCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(benchCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
		"lock":       carapace.ActionFiles(),
	})

	carapace.Gen(benchCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
