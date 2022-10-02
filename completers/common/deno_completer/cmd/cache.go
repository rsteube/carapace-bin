package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache the dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	cacheCmd.Flags().StringP("config", "c", "", "Load configuration file")
	cacheCmd.Flags().String("import-map", "", "Load import map file")
	cacheCmd.Flags().String("lock", "", "Check the specified lock file")
	cacheCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	cacheCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	cacheCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	cacheCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	rootCmd.AddCommand(cacheCmd)

	cacheCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(cacheCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
	})

	carapace.Gen(cacheCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
