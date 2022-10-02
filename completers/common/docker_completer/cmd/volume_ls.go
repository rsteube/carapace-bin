package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volume_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_lsCmd).Standalone()

	volume_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. 'dangling=true')")
	volume_lsCmd.Flags().String("format", "", "Pretty-print volumes using a Go template")
	volume_lsCmd.Flags().BoolP("quiet", "q", false, "Only display volume names")
	volumeCmd.AddCommand(volume_lsCmd)
}
