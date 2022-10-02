package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uptime",
	Short: "Tell how long the system has been running",
	Long:  "https://linux.die.net/man/1/uptime",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("pretty", "p", false, "show uptime in pretty format")
	rootCmd.Flags().BoolP("since", "s", false, "system up since")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
}
