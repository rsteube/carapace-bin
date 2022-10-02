package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var engine_activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate Enterprise Edition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(engine_activateCmd).Standalone()

	engine_activateCmd.Flags().String("containerd", "", "override default location of containerd endpoint")
	engine_activateCmd.Flags().Bool("display-only", false, "only display license information and exit")
	engine_activateCmd.Flags().String("engine-image", "", "Specify engine image")
	engine_activateCmd.Flags().String("format", "", "Pretty-print licenses using a Go template")
	engine_activateCmd.Flags().String("license", "", "License File")
	engine_activateCmd.Flags().Bool("quiet", false, "Only display available licenses by ID")
	engine_activateCmd.Flags().String("registry-prefix", "", "Override the default location where engine images are pulled")
	engine_activateCmd.Flags().String("version", "", "Specify engine version (default is to use currently running")
	engineCmd.AddCommand(engine_activateCmd)

	carapace.Gen(engine_activateCmd).FlagCompletion(carapace.ActionMap{
		"license": carapace.ActionFiles(),
	})
}
