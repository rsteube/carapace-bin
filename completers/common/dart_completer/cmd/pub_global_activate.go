package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_global_activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Make a package's executables globally available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_global_activateCmd).Standalone()

	pub_global_activateCmd.Flags().StringP("executable", "x", "", "Executable(s) to place on PATH.")
	pub_global_activateCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_global_activateCmd.Flags().StringP("hosted-url", "u", "", "A custom pub server URL for the package.")
	pub_global_activateCmd.Flags().Bool("no-executables", false, "Do not put executables on PATH.")
	pub_global_activateCmd.Flags().Bool("overwrite", false, "Overwrite executables from other packages with the same")
	pub_global_activateCmd.Flags().StringP("source", "s", "", "The source used to find the package.")
	pub_globalCmd.AddCommand(pub_global_activateCmd)

	carapace.Gen(pub_global_activateCmd).FlagCompletion(carapace.ActionMap{
		// TODO executables
		"source": carapace.ActionValues("git", "hosted", "path"),
	})
}
