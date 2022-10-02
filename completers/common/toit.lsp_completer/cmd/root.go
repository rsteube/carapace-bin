package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toit.lsp",
	Short: "start the lsp server",
	Long:  "https://github.com/toitlang/toit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().String("cpuprofile", "", "write cpu profile to `file`")
	rootCmd.Flags().BoolP("help", "h", false, "help for toitlsp")
	rootCmd.Flags().String("memprofile", "", "write mem profile to `file`")
	rootCmd.Flags().String("sdk", "", "the default SDK path to use")
	rootCmd.Flags().String("toitc", "", "the default toit compiler to use")
	rootCmd.Flags().BoolP("verbose", "v", false, "")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"memprofile": carapace.ActionFiles(),
		"sdk":        carapace.ActionDirectories(),
		"toitc":      carapace.ActionFiles(),
	})
}
