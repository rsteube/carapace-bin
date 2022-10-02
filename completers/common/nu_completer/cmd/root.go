package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nu",
	Short: "Nushell",
	Long:  "https://www.nushell.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("commands", "c", "", "")
	rootCmd.Flags().String("debug", "", "")
	rootCmd.Flags().String("develop", "", "")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().StringP("loglevel", "l", "", "[possible values: error, warn, info, debug, trace]")
	rootCmd.Flags().Bool("stdin", false, "")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"loglevel": carapace.ActionValues("error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
