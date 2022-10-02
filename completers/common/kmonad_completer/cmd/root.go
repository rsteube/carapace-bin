package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kmonad",
	Short: "an onion of buttons",
	Long:  "https://github.com/kmonad/kmonad",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dry-run", "d", false, "only try parsing the config file")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help text")
	rootCmd.Flags().StringP("log-level", "l", "", "How much info to print out")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-level": carapace.ActionValues("debug", "info", "warn", "error").StyleF(style.ForLogLevel),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".kbd"),
	)
}
