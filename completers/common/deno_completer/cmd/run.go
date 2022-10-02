package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a JavaScript or TypeScript program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("allow-all", "A", false, "Allow all permissions")
	runCmd.Flags().String("allow-env", "", "Allow environment access")
	runCmd.Flags().String("allow-ffi", "", "Allow loading dynamic libraries")
	runCmd.Flags().Bool("allow-hrtime", false, "Allow high resolution time measurement")
	runCmd.Flags().String("allow-net", "", "Allow network access")
	runCmd.Flags().String("allow-read", "", "Allow file system read access")
	runCmd.Flags().String("allow-run", "", "Allow running subprocesses")
	runCmd.Flags().String("allow-write", "", "Allow file system write access")
	runCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	runCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	runCmd.Flags().StringP("config", "c", "", "Load configuration file")
	runCmd.Flags().String("import-map", "", "Load import map file")
	runCmd.Flags().String("inspect", "", "Activate inspector on host:port (default: 127.0.0.1:9229)")
	runCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script")
	runCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	runCmd.Flags().String("lock", "", "Check the specified lock file")
	runCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	runCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	runCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	runCmd.Flags().Bool("prompt", false, "Fallback to prompt if required permission wasn't passed")
	runCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	runCmd.Flags().String("seed", "", "Seed Math.random()")
	runCmd.Flags().String("unsafely-ignore-certificate-errors", "", "DANGER: Disables verification of TLS certificates")
	runCmd.Flags().String("v8-flags", "", "Set V8 command line options (for help: --v8-flags=--help")
	runCmd.Flags().Bool("watch", false, "UNSTABLE: Watch for file changes and restart process automatically")
	rootCmd.AddCommand(runCmd)

	runCmd.Flag("allow-env").NoOptDefVal = " "
	runCmd.Flag("allow-ffi").NoOptDefVal = " "
	runCmd.Flag("allow-net").NoOptDefVal = " "
	runCmd.Flag("allow-read").NoOptDefVal = " "
	runCmd.Flag("allow-run").NoOptDefVal = " "
	runCmd.Flag("allow-write").NoOptDefVal = " "
	runCmd.Flag("inspect").NoOptDefVal = " "
	runCmd.Flag("inspect-brk").NoOptDefVal = " "
	runCmd.Flag("reload").NoOptDefVal = " "
	runCmd.Flag("unsafely-ignore-certificate-errors").NoOptDefVal = " "
	runCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(),
		"config":      carapace.ActionFiles(),
		"import-map":  carapace.ActionFiles(),
		"inspect":     action.ActionHostPort(),
		"inspect-brk": action.ActionHostPort(),
		"lock":        carapace.ActionFiles(),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
