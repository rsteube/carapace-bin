package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pamac",
	Short: "package manager utility",
	Long:  "https://wiki.manjaro.org/index.php/Pamac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("help", "h", "", "show help")
	rootCmd.Flags().Bool("version", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"help": carapace.ActionValues("search", "list", "info", "install", "reinstall", "remove", "checkupdates", "update", "upgrade", "clone", "build", "clean"),
	})

}
