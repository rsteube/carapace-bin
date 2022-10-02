package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build or rebuild services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().StringArray("build-arg", []string{}, "Set build-time variables for services.")
	buildCmd.Flags().Bool("compress", true, "Compress the build context using gzip. DEPRECATED")
	buildCmd.Flags().Bool("force-rm", true, "Always remove intermediate containers. DEPRECATED")
	buildCmd.Flags().StringP("memory", "m", "", "Set memory limit for the build container. Not supported on buildkit yet.")
	buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	buildCmd.Flags().Bool("no-rm", false, "Do not remove intermediate containers after a successful build. DEPRECATED")
	buildCmd.Flags().Bool("parallel", true, "Build images in parallel. DEPRECATED")
	buildCmd.Flags().String("progress", "auto", "Set type of progress output (auto, tty, plain, quiet)")
	buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image.")
	buildCmd.Flags().BoolP("quiet", "q", false, "Don't print anything to STDOUT")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "tty", "plain", "quiet"),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(buildCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
