package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "visudo",
	Short: "safely edit the sudoers file",
	Long:  "https://linux.die.net/man/8/visudo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("check", "c", false, "check-only mode")
	rootCmd.Flags().StringP("file", "f", "", "specify sudoers file location")
	rootCmd.Flags().BoolP("help", "h", false, "display help message and exit")
	rootCmd.Flags().BoolP("quiet", "q", false, "less verbose (quiet) syntax error messages")
	rootCmd.Flags().BoolP("strict", "s", false, "strict syntax checking")
	rootCmd.Flags().BoolP("version", "V", false, "display version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
