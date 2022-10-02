package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve previously built Gatsby site",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().StringP("host", "H", "", "Set host.")
	serveCmd.Flags().BoolP("open", "o", false, "Open the site in your (default) browser for you.")
	serveCmd.Flags().String("open-tracing-config-file", "", "Tracer configuration file.")
	serveCmd.Flags().StringP("port", "p", "", "Set port.")
	serveCmd.Flags().Bool("prefix-paths", false, "Serve site with link paths prefixed with the pathPrefix.")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"open-tracing-config-file": carapace.ActionFiles(),
	})
}
