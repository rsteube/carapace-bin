package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff live version against would-be applied version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	diffCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files contains the configuration to diff")
	diffCmd.Flags().Bool("force-conflicts", false, "If true, server-side apply will force the changes against conflicts.")
	diffCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	diffCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	diffCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	diffCmd.Flags().Bool("server-side", false, "If true, apply runs in the server instead of the client.")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
	})
}
