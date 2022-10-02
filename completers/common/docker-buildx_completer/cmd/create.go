package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()
	createCmd.Flags().Bool("append", false, "Append a node to builder instead of changing it")
	createCmd.Flags().Bool("bootstrap", false, "Boot builder after creation")
	createCmd.Flags().String("buildkitd-flags", "", "Flags for buildkitd daemon")
	createCmd.Flags().String("config", "", "BuildKit config file")
	createCmd.Flags().String("driver", "", "Driver to use (available: \"docker\", \"docker-container\", \"kubernetes\")")
	createCmd.Flags().StringArray("driver-opt", []string{}, "Options for the driver")
	createCmd.Flags().Bool("leave", false, "Remove a node from builder instead of changing it")
	createCmd.Flags().String("name", "", "Builder instance name")
	createCmd.Flags().String("node", "", "Create/modify node with given name")
	createCmd.Flags().StringArray("platform", []string{}, "Fixed platforms for current node")
	createCmd.Flags().Bool("use", false, "Set the current builder instance")
	rootCmd.AddCommand(createCmd)

	// TODO flag completion
	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"driver": carapace.ActionValues("docker", "docker-container", "kubernetes"),
	})

	// TODO positional completion
}
