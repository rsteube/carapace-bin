package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var xfrmCmd = &cobra.Command{
	Use:   "xfrm",
	Short: "manage IPSec policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xfrmCmd).Standalone()

	rootCmd.AddCommand(xfrmCmd)
}
