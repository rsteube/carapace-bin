package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamav-config",
	Short: "clamav config",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("cflags", false, "print pre-processor and compiler flags")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("libs", false, "print library linking information")
	rootCmd.Flags().String("prefix", "", "change libclamav prefix [default /usr]")
	rootCmd.Flags().Bool("version", false, "output version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
	})
}
