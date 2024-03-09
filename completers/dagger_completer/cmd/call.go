package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var callCmd = &cobra.Command{
	Use:     "call [flags] [FUNCTION]...",
	Short:   "Call a module function",
	GroupID: "module",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(callCmd).Standalone()
	callCmd.Flags().SetInterspersed(false)

	callCmd.PersistentFlags().Bool("focus", false, "Only show output for focused commands")
	callCmd.PersistentFlags().Bool("json", false, "Present result as JSON")
	callCmd.PersistentFlags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	callCmd.PersistentFlags().StringP("output", "o", "", "Path in the host to save the result to")
	rootCmd.AddCommand(callCmd)

	carapace.Gen(callCmd).FlagCompletion(carapace.ActionMap{
		"mod":    dagger.ActionMods(),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(callCmd).PositionalAnyCompletion(
		dagger.ActionFunctions(),
	)
}
