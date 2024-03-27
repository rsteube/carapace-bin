package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new kustomization in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()
	createCmd.Flags().String("annotations", "", "Add one or more common annotations.")
	createCmd.Flags().Bool("autodetect", false, "Search for kubernetes resources in the current directory to be added to the kustomization file.")
	createCmd.Flags().String("labels", "", "Add one or more common labels.")
	createCmd.Flags().String("nameprefix", "", "Sets the value of the namePrefix field in the kustomization file.")
	createCmd.Flags().String("namespace", "", "Set the value of the namespace field in the customization file.")
	createCmd.Flags().String("namesuffix", "", "Sets the value of the nameSuffix field in the kustomization file.")
	createCmd.Flags().Bool("recursive", false, "Enable recursive directory searching for resource auto-detection.")
	createCmd.Flags().String("resources", "", "Name of a file containing a file to add to the kustomization file.")
	rootCmd.AddCommand(createCmd)
}
