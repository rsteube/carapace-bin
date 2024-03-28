package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Create or update a Alert resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_alertCmd).Standalone()
	create_alertCmd.Flags().String("event-severity", "", "severity of events to send alerts for")
	create_alertCmd.Flags().StringSlice("event-source", []string{}, "sources that should generate alerts (<kind>/<name>), also accepts comma-separated values")
	create_alertCmd.Flags().String("provider-ref", "", "reference to provider")
	createCmd.AddCommand(create_alertCmd)
}
