package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Eval script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evalCmd).Standalone()

	evalCmd.Flags().Bool("cached-only", false, "Require that remote dependencies are already cached")
	evalCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	evalCmd.Flags().StringP("config", "c", "", "Load configuration file")
	evalCmd.Flags().String("ext", "", "Set standard input (stdin) content type [default: js]  [possible values: ts, tsx, js, jsx]")
	evalCmd.Flags().String("import-map", "", "Load import map file")
	evalCmd.Flags().String("inspect", "", "Activate inspector on host:port (default: 127.0.0.1:9229)")
	evalCmd.Flags().String("inspect-brk", "", "Activate inspector on host:port and break at start of user script")
	evalCmd.Flags().String("location", "", "Value of 'globalThis.location' used by some web APIs")
	evalCmd.Flags().String("lock", "", "Check the specified lock file")
	evalCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	evalCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	evalCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	evalCmd.Flags().BoolP("print", "p", false, "print result to stdout")
	evalCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	evalCmd.Flags().String("seed", "", "Seed Math.random()")
	rootCmd.AddCommand(evalCmd)

	carapace.Gen(evalCmd).FlagCompletion(carapace.ActionMap{
		"cert":        carapace.ActionFiles(),
		"config":      carapace.ActionFiles(),
		"ext":         carapace.ActionValues("ts", "tsx", "js", "jsx"),
		"import-map":  carapace.ActionFiles(),
		"inspect":     action.ActionHostPort(),
		"inspect-brk": action.ActionHostPort(),
		"lock":        carapace.ActionFiles(),
	})
}
