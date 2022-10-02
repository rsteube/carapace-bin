package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamconf",
	Short: "Clam AntiVirus configuration utility",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config-dir", "c", "", "Read configuration files from DIR")
	rootCmd.Flags().StringP("generate-config", "g", "", "Generate example config file")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("non-default", "n", false, "Only display non-default settings")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir":      carapace.ActionDirectories(),
		"generate-config": carapace.ActionFiles(),
	})
}
