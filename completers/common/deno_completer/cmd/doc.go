package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Show documentation for a modul",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().String("import-map", "", "Load import map file")
	docCmd.Flags().Bool("json", false, "Output documentation in JSON format")
	docCmd.Flags().Bool("private", false, "Output private documentation")
	docCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	rootCmd.AddCommand(docCmd)

	docCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(docCmd).FlagCompletion(carapace.ActionMap{
		"import-map": carapace.ActionFiles(),
	})

	carapace.Gen(docCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
