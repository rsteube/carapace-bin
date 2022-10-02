package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "archive a set of entrypoints into a toit archive file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archiveCmd).Standalone()
	archiveCmd.Flags().Bool("include-sdk", false, "if set, will include the used SDK files in the archive file")
	archiveCmd.Flags().String("out", "archive.tar", "the output file. Use: '-' for stdout")
	archiveCmd.Flags().String("sdk", "", "the default SDK path to use")
	archiveCmd.Flags().String("toitc", "", "the default toit compiler to use")
	archiveCmd.Flags().BoolP("verbose", "v", false, "")
	rootCmd.AddCommand(archiveCmd)

	carapace.Gen(archiveCmd).FlagCompletion(carapace.ActionMap{
		"out":   carapace.ActionFiles(),
		"sdk":   carapace.ActionDirectories(),
		"toitc": carapace.ActionFiles(),
	})

	carapace.Gen(archiveCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
