package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsb_release",
	Short: "prints certain LSB (Linux Standard Base) and Distribution information",
	Long:  "https://linux.die.net/man/1/lsb_release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Display all of the above information.")
	rootCmd.Flags().BoolP("codename", "c", false, "Display the codename according to the distribution release.")
	rootCmd.Flags().BoolP("description", "d", false, "Display the single line text description of the distribution.")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message.")
	rootCmd.Flags().BoolP("id", "i", false, "Display the string id of the distributor.")
	rootCmd.Flags().BoolP("release", "r", false, "Display the release number of the distribution.")
	rootCmd.Flags().BoolP("short", "s", false, "Display all of the above information in short output format.")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the LSB specification against which the distribution is compliant.")
}
