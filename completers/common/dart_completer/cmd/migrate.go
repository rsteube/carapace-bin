package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Perform null safety migration on a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().Bool("apply-changes", false, "Apply the proposed null safety changes to the files on disk.")
	migrateCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	migrateCmd.Flags().Bool("ignore-errors", false, "Attempt to perform null safety analysis even if the project has analysis errors.")
	migrateCmd.Flags().Bool("no-web-preview", false, "Print proposed changes to the console.")
	migrateCmd.Flags().String("preview-hostname", "", "Run the preview server on the specified hostname.")
	migrateCmd.Flags().String("preview-port", "", "Run the preview server on the specified port.")
	migrateCmd.Flags().Bool("skip-import-check", false, "Go ahead with migration even if some imported files have not yet been migrated.")
	migrateCmd.Flags().String("summary", "", "Output a machine-readable summary of migration changes.")
	migrateCmd.Flags().BoolP("verbose", "v", false, "Show additional command output.")
	migrateCmd.Flags().Bool("web-preview", false, "Show an interactive preview of the proposed null safety changes in a browser window.")
	rootCmd.AddCommand(migrateCmd)

	carapace.Gen(migrateCmd).FlagCompletion(carapace.ActionMap{
		"preview-hostname": carapace.ActionValues("localhost", "any"),
		"summary":          carapace.ActionFiles(),
	})

	carapace.Gen(migrateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
