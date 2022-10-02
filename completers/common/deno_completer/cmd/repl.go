package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Read Eval Print Loop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replCmd).Standalone()

	replCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	replCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	replCmd.Flags().StringP("config", "c", "", "Load configuration file")
	replCmd.Flags().String("eval", "", "Evaluates the provided code when the REPL starts.")
	replCmd.Flags().String("import-map", "", "Load import map file")
	replCmd.Flags().String("inspect", "", "Activate inspector on host:port (default: 127.0.0.1:9229)")
	replCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script")
	replCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	replCmd.Flags().String("lock", "", "Check the specified lock file")
	replCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	replCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	replCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	replCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	replCmd.Flags().String("seed", "", "Seed Math.random()")
	replCmd.Flags().String("v8-flags", "", "Set V8 command line options (for help: --v8-flags=--help")
	rootCmd.AddCommand(replCmd)

	replCmd.Flag("inspect").NoOptDefVal = " "
	replCmd.Flag("inspect-brk").NoOptDefVal = " "
	replCmd.Flag("reload").NoOptDefVal = " "
	replCmd.Flag("v8-flags").NoOptDefVal = " "

	carapace.Gen(replCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(),
		"config":      carapace.ActionFiles(),
		"import-map":  carapace.ActionFiles(),
		"inspect":     action.ActionHostPort(),
		"inspect-brk": action.ActionHostPort(),
		"lock":        carapace.ActionFiles(),
	})
}
