package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deno",
	Short: "A modern JavaScript and TypeScript runtime",
	Long:  "https://deno.land/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Prints help information")
	rootCmd.PersistentFlags().StringP("log-level", "L", "", "Set log level [possible values: debug, info]")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	rootCmd.PersistentFlags().Bool("unstable", false, "Enable unstable features and APIs")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-level": carapace.ActionValues("debug", "info").StyleF(style.ForLogLevel),
	})
}
