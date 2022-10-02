package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "resume a suspended vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resumeCmd).Standalone()

	resumeCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	resumeCmd.Flags().Bool("provision", false, "Enable provisioning")
	resumeCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	rootCmd.AddCommand(resumeCmd)

	carapace.Gen(resumeCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProvisioners().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(resumeCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
