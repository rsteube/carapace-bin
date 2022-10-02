package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile the script into a self contained executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().BoolP("allow-all", "A", false, "Allow all permissions")
	compileCmd.Flags().String("allow-env", "", "Allow environment access")
	compileCmd.Flags().String("allow-ffi", "", "Allow loading dynamic libraries")
	compileCmd.Flags().Bool("allow-hrtime", false, "Allow high resolution time measurement")
	compileCmd.Flags().String("allow-net", "", "Allow network access")
	compileCmd.Flags().String("allow-read", "", "Allow file system read access")
	compileCmd.Flags().String("allow-run", "", "Allow running subprocesses")
	compileCmd.Flags().String("allow-write", "", "Allow file system write access")
	compileCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	compileCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	compileCmd.Flags().StringP("config", "c", "", "Load configuration file")
	compileCmd.Flags().String("import-map", "", "Load import map file")
	compileCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	compileCmd.Flags().String("lock", "", "Check the specified lock file")
	compileCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	compileCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	compileCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	compileCmd.Flags().StringP("output", "o", "", "Output file (defaults to $PWD/<inferred-name>)")
	compileCmd.Flags().Bool("prompt", false, "Fallback to prompt if required permission wasn't passed")
	compileCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	compileCmd.Flags().String("seed", "", "Seed Math.random()")
	compileCmd.Flags().String("target", "", "Target OS architecture [possible values: x86_64-unknown-linux-gnu, x86_64-pc-windows-msvc, x86_64-apple-darwin, aarch64-apple-darwin]")
	compileCmd.Flags().String("unsafely-ignore-certificate-errors", "", "DANGER: Disables verification of TLS certificates")
	compileCmd.Flags().String("v8-flags", "", "Set V8 command line options (for help: --v8-flags=--help")
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flag("allow-env").NoOptDefVal = " "
	compileCmd.Flag("allow-ffi").NoOptDefVal = " "
	compileCmd.Flag("allow-net").NoOptDefVal = " "
	compileCmd.Flag("allow-read").NoOptDefVal = " "
	compileCmd.Flag("allow-run").NoOptDefVal = " "
	compileCmd.Flag("allow-write").NoOptDefVal = " "
	compileCmd.Flag("reload").NoOptDefVal = " "
	compileCmd.Flag("unsafely-ignore-certificate-errors").NoOptDefVal = " "
	compileCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(compileCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
		"lock":       carapace.ActionFiles(),
		"output":     carapace.ActionFiles(),
		"target":     carapace.ActionValues("x86_64-unknown-linux-gnu", "x86_64-pc-windows-msvc", "x86_64-apple-darwin", "aarch64-apple-darwin"),
	})

	carapace.Gen(compileCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
