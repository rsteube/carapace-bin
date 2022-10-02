package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload locally saved terminal session to asciinema.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
