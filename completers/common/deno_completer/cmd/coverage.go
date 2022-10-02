package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var coverageCmd = &cobra.Command{
	Use:   "coverage",
	Short: "Print coverage reports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coverageCmd).Standalone()

	coverageCmd.Flags().String("exclude", "", "Exclude source files from the report")
	coverageCmd.Flags().String("ignore", "", "Ignore coverage files")
	coverageCmd.Flags().String("include", "", "Include source files in the report [default: ^file:]")
	coverageCmd.Flags().Bool("lcov", false, "Output coverage report in lcov format")
	rootCmd.AddCommand(coverageCmd)

	coverageCmd.Flag("exclude").NoOptDefVal = " "
	coverageCmd.Flag("ignore").NoOptDefVal = " "
	coverageCmd.Flag("include").NoOptDefVal = " "

	carapace.Gen(coverageCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
