package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_kustomizationsCmd = &cobra.Command{
	Use:   "kustomizations",
	Short: "Get Kustomization statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_kustomizationsCmd).Standalone()
	getCmd.AddCommand(get_kustomizationsCmd)
}
