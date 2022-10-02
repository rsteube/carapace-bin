package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Type-check the dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	checkCmd.Flags().StringP("config", "c", "", "Specify the configuration file")
	checkCmd.Flags().BoolP("help", "h", false, "Print help information")
	checkCmd.Flags().String("import-map", "", "Load import map file")
	checkCmd.Flags().String("lock", "", "Check the specified lock file")
	checkCmd.Flags().Bool("no-config", false, "Disable automatic loading of the configuration file.")
	checkCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	checkCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	checkCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	checkCmd.Flags().Bool("remote", false, "Type-check all modules, including remote")
	checkCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
		"lock":       carapace.ActionFiles(),
	})

	carapace.Gen(checkCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
