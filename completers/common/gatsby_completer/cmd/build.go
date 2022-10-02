package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Gatsby project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("graphql-tracing", false, "Trace every graphql resolver, may have performance implications.")
	buildCmd.Flags().Bool("no-uglify", false, "Build site without uglifying JS bundles.")
	buildCmd.Flags().String("open-tracing-config-file", "", "Tracer configuration file.")
	buildCmd.Flags().Bool("prefix-paths", false, "Build site with link paths prefixed with the pathPrefix.")
	buildCmd.Flags().Bool("profile", false, "Build site with react profiling.")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"open-tracing-config-file": carapace.ActionFiles(),
	})
}
