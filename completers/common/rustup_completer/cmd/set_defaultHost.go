package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var set_defaultHostCmd = &cobra.Command{
	Use:   "default-host",
	Short: "The triple used to identify toolchains when not specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_defaultHostCmd).Standalone()

	set_defaultHostCmd.Flags().BoolP("help", "h", false, "Prints help information")
	set_defaultHostCmd.Flags().BoolP("version", "V", false, "Prints version information")
	setCmd.AddCommand(set_defaultHostCmd)
}
