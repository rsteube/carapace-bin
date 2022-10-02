package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert your content to different formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	convertCmd.PersistentFlags().StringP("output", "o", "", "filesystem path to write files to")
	convertCmd.PersistentFlags().Bool("unsafe", false, "enable less safe operations, please backup first")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
