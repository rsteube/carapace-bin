package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toitdocCmd = &cobra.Command{
	Use:   "toitdoc",
	Short: "Generates a toitdoc json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toitdocCmd).Standalone()
	toitdocCmd.Flags().Bool("exclude-sdk", false, "if set, will remove the sdk libraries from the toitdoc")
	toitdocCmd.Flags().Bool("include-private", false, "if set, will include private toitdoc for private elements")
	toitdocCmd.Flags().String("out", "toitdoc.json", "the output file")
	toitdocCmd.Flags().UintP("parallel", "p", 1, "parallelism")
	toitdocCmd.Flags().String("root-path", "", "root path to build paths from")
	toitdocCmd.Flags().String("sdk", "", "the SDK path to use")
	toitdocCmd.Flags().String("toitc", "", "the toit compiler to use")
	toitdocCmd.Flags().BoolP("verbose", "v", false, "")
	toitdocCmd.Flags().String("version", "", "version of the package to build toitdoc for")
	rootCmd.AddCommand(toitdocCmd)

	carapace.Gen(toitdocCmd).FlagCompletion(carapace.ActionMap{
		"out":       carapace.ActionFiles(),
		"root-path": carapace.ActionDirectories(),
		"sdk":       carapace.ActionDirectories(),
		"toitc":     carapace.ActionFiles(),
	})

	carapace.Gen(toitdocCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
