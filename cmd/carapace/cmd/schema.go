package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "--codegen",
	Short: "json schema for spec files",
	Run: func(cmd *cobra.Command, args []string) {
		if schema, err := spec.Schema(); err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err.Error()) // TODO fail / exit 1 ?
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), schema)
		}
	},
}

func init() {
	carapace.Gen(schemaCmd).Standalone()

}
