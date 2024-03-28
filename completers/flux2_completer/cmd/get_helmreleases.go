package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_helmreleasesCmd = &cobra.Command{
	Use:   "helmreleases",
	Short: "Get HelmRelease statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_helmreleasesCmd).Standalone()
	getCmd.AddCommand(get_helmreleasesCmd)
}
