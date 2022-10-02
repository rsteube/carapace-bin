package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connection_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import an external configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_importCmd).Standalone()

	connection_importCmd.Flags().Bool("temporary", false, "only temporary")
	connectionCmd.AddCommand(connection_importCmd)

	carapace.Gen(connection_importCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
