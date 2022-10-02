package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generate_libraryCmd = &cobra.Command{
	Use:   "library",
	Short: "Creates a new, generic library project in the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_libraryCmd).Standalone()

	generate_libraryCmd.Flags().String("entry-file", "", "The path at which to create the library's public API file")
	generate_libraryCmd.Flags().StringP("prefix", "p", "", "A prefix to apply to generated selectors.")
	generate_libraryCmd.Flags().Bool("skip-install", false, "Do not install dependency packages.")
	generate_libraryCmd.Flags().Bool("skip-package-json", false, "Do not add dependencies to the \"package.json\" file.")
	generate_libraryCmd.Flags().Bool("skip-ts-config", false, "Do not update \"tsconfig.json\" to add a path mapping for the new library.")
	generateCmd.AddCommand(generate_libraryCmd)
}
