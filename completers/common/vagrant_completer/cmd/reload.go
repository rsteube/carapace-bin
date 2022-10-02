package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "restarts vagrant machine, loads new Vagrantfile configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()

	reloadCmd.Flags().BoolP("force", "f", false, "Force shut down (equivalent of pulling power)")
	reloadCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	reloadCmd.Flags().Bool("provision", false, "Enable provisioning")
	reloadCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	rootCmd.AddCommand(reloadCmd)

	carapace.Gen(reloadCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProvisioners().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(reloadCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
