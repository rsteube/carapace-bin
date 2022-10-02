package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envsubst",
	Short: "Substitutes the values of environment variables",
	Long:  "https://linux.die.net/man/1/envsubst",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("variables", "v", false, "output the variables occurring in SHELL-FORMAT")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
}
