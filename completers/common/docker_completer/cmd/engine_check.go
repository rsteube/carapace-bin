package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var engine_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for available engine updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(engine_checkCmd).Standalone()

	engine_checkCmd.Flags().String("containerd", "", "override default location of containerd endpoint")
	engine_checkCmd.Flags().Bool("downgrades", false, "Report downgrades (default omits older versions)")
	engine_checkCmd.Flags().String("engine-image", "", "Specify engine image (default uses the same image as currently")
	engine_checkCmd.Flags().String("format", "", "Pretty-print updates using a Go template")
	engine_checkCmd.Flags().Bool("pre-releases", false, "Include pre-release versions")
	engine_checkCmd.Flags().BoolP("quiet", "q", false, "Only display available versions")
	engine_checkCmd.Flags().String("registry-prefix", "", "Override the existing location where engine images are pulled")
	engine_checkCmd.Flags().Bool("upgrades", false, "Report available upgrades (default true)")
	engineCmd.AddCommand(engine_checkCmd)
}
