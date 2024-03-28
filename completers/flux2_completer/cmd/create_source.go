package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Create or update sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_sourceCmd).Standalone()
	create_sourceCmd.PersistentFlags().String("fetch-timeout", "", "set a timeout for fetch operations performed by source-controller (e.g. 'git clone' or 'helm repo update')")
	createCmd.AddCommand(create_sourceCmd)
}
