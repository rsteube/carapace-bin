package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var set_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "The default components installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_profileCmd).Standalone()

	set_profileCmd.Flags().BoolP("help", "h", false, "Prints help information")
	set_profileCmd.Flags().BoolP("version", "V", false, "Prints version information")
	setCmd.AddCommand(set_profileCmd)

	carapace.Gen(set_profileCmd).PositionalCompletion(
		carapace.ActionValues("minimal", "default", "complete"),
	)
}
