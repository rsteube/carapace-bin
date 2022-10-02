package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove containers, networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()
	downCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	downCmd.Flags().String("rmi", "", "Remove images used by services. \"local\" remove only images that don't have a custom tag (\"local\"|\"all\")")
	downCmd.Flags().IntP("timeout", "t", 10, "Specify a shutdown timeout in seconds")
	downCmd.Flags().BoolP("volumes", "v", false, " Remove named volumes declared in the `volumes` section of the Compose file and anonymous volumes attached to containers.")
	rootCmd.AddCommand(downCmd)

	carapace.Gen(downCmd).FlagCompletion(carapace.ActionMap{
		"rmi": carapace.ActionValuesDescribed(
			"all", "Remove all images used by any service.",
			"local", "Remove only images that don't have a custom tag set by the image field.",
		),
	})
}
