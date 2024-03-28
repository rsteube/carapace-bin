package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the client and server-side components version information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().Bool("client", false, "print only client version")
	versionCmd.Flags().StringP("output", "o", "yaml", "the format in which the information should be printed. can be 'json' or 'yaml'")
	rootCmd.AddCommand(versionCmd)
}
