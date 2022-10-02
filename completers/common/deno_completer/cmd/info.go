package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info about cache or info related to source file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	infoCmd.Flags().String("import-map", "", "Load import map file")
	infoCmd.Flags().Bool("json", false, "UNSTABLE: Outputs the information in JSON format")
	infoCmd.Flags().String("location", "", "Show files used for origin bound APIs like the Web Storage API when running a script with '--location=<HREF>'")
	infoCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
	})

	carapace.Gen(infoCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
