package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hostname",
	Short: "show or set system host name",
	Long:  "https://linux.die.net/man/1/hostname",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("aliases", "a", false, "alias names")
	rootCmd.Flags().BoolP("domain", "d", false, "DNS domain name")
	rootCmd.Flags().StringP("file", "F", "", "set host name or NIS domain name from FILE")
	rootCmd.Flags().StringP("fqdn,", "f", "", "DNS host name or FQDN")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().BoolP("ip-addresses", "i", false, "addresses for the host name")
	rootCmd.Flags().BoolP("short", "s", false, "short host name")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")
	rootCmd.Flags().StringP("yp,", "y", "", "NIS/YP domain name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
