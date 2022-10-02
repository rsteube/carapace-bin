package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generates a description of the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()
	describeCmd.Flags().Bool("allow-local-deps", false, "Allow local dependencies and don't report them")
	describeCmd.Flags().Bool("disallow-local-deps", false, "Always disallow local dependencies and report them as error")
	describeCmd.Flags().String("out-dir", "", "Output directory of description files")
	describeCmd.Flags().BoolP("verbose", "v", false, "Show more information")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).FlagCompletion(carapace.ActionMap{
		"out-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(describeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
