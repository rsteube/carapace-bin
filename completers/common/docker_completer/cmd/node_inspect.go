package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_inspectCmd).Standalone()

	node_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	node_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	nodeCmd.AddCommand(node_inspectCmd)

	carapace.Gen(node_inspectCmd).PositionalAnyCompletion(docker.ActionNodes())
}
