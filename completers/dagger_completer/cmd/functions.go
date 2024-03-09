package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var functionsCmd = &cobra.Command{
	Use:     "functions [flags] [FUNCTION]...",
	Short:   "List available functions",
	GroupID: "module",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(functionsCmd).Standalone()

	functionsCmd.PersistentFlags().Bool("focus", false, "Only show output for focused commands")
	functionsCmd.PersistentFlags().Bool("json", false, "Output as json")
	functionsCmd.PersistentFlags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	rootCmd.AddCommand(functionsCmd)

	carapace.Gen(functionsCmd).FlagCompletion(carapace.ActionMap{
		"mod": dagger.ActionMods(),
	})

	carapace.Gen(functionsCmd).PositionalAnyCompletion(
		dagger.ActionFunctions().FilterArgs(),
	)
}
