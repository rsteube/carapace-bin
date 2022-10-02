package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install script as an executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("allow-all", "A", false, "Allow all permissions")
	installCmd.Flags().String("allow-env", "", "Allow environment access")
	installCmd.Flags().String("allow-ffi", "", "Allow loading dynamic libraries")
	installCmd.Flags().Bool("allow-hrtime", false, "Allow high resolution time measurement")
	installCmd.Flags().String("allow-net", "", "Allow network access")
	installCmd.Flags().String("allow-read", "", "Allow file system read access")
	installCmd.Flags().String("allow-run", "", "Allow running subprocesses")
	installCmd.Flags().String("allow-write", "", "Allow file system write access")
	installCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	installCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	installCmd.Flags().StringP("config", "c", "", "Load configuration file")
	installCmd.Flags().BoolP("force", "f", false, "Forcefully overwrite existing installation")
	installCmd.Flags().String("import-map", "", "Load import map file")
	installCmd.Flags().String("inspect", "", "Activate inspector on host:port (default: 127.0.0.1:9229)")
	installCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script")
	installCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	installCmd.Flags().String("lock", "", "Check the specified lock file")
	installCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	installCmd.Flags().StringP("name", "n", "", "Executable file name")
	installCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	installCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	installCmd.Flags().Bool("prompt", false, "Fallback to prompt if required permission wasn't passed")
	installCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	installCmd.Flags().String("root", "", "Installation root")
	installCmd.Flags().String("seed", "", "Seed Math.random()")
	installCmd.Flags().String("unsafely-ignore-certificate-errors", "", "DANGER: Disables verification of TLS certificates")
	installCmd.Flags().String("v8-flags", "", "Set V8 command line options (for help: --v8-flags=--help")
	rootCmd.AddCommand(installCmd)

	installCmd.Flag("allow-env").NoOptDefVal = " "
	installCmd.Flag("allow-ffi").NoOptDefVal = " "
	installCmd.Flag("allow-net").NoOptDefVal = " "
	installCmd.Flag("allow-read").NoOptDefVal = " "
	installCmd.Flag("allow-run").NoOptDefVal = " "
	installCmd.Flag("allow-write").NoOptDefVal = " "
	installCmd.Flag("inspect").NoOptDefVal = " "
	installCmd.Flag("inspect-brk").NoOptDefVal = " "
	installCmd.Flag("reload").NoOptDefVal = " "
	installCmd.Flag("unsafely-ignore-certificate-errors").NoOptDefVal = " "
	installCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(),
		"config":      carapace.ActionFiles(),
		"import-map":  carapace.ActionFiles(),
		"inspect":     action.ActionHostPort(),
		"inspect-brk": action.ActionHostPort(),
		"lock":        carapace.ActionFiles(),
		"root":        carapace.ActionDirectories(),
	})
}
