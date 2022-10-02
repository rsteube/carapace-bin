package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an image from a Dockerfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_buildCmd).Standalone()

	image_buildCmd.Flags().String("add-host", "", "Add a custom host-to-IP mapping (host:ip)")
	image_buildCmd.Flags().String("build-arg", "", "Set build-time variables")
	image_buildCmd.Flags().String("cache-from", "", "Images to consider as cache sources")
	image_buildCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	image_buildCmd.Flags().Bool("compress", false, "Compress the build context using gzip")
	image_buildCmd.Flags().String("cpu-period", "", "Limit the CPU CFS (Completely Fair Scheduler) period")
	image_buildCmd.Flags().String("cpu-quota", "", "Limit the CPU CFS (Completely Fair Scheduler) quota")
	image_buildCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	image_buildCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	image_buildCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	image_buildCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (default true)")
	image_buildCmd.Flags().StringP("file", "f", "", "Name of the Dockerfile (Default is 'PATH/Dockerfile')")
	image_buildCmd.Flags().Bool("force-rm", false, "Always remove intermediate containers")
	image_buildCmd.Flags().String("iidfile", "", "Write the image ID to the file")
	image_buildCmd.Flags().String("isolation", "", "Container isolation technology")
	image_buildCmd.Flags().String("label", "", "Set metadata for an image")
	image_buildCmd.Flags().StringP("memory", "m", "", "Memory limit")
	image_buildCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	image_buildCmd.Flags().String("network", "", "Set the networking mode for the RUN instructions during build")
	image_buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	image_buildCmd.Flags().StringP("output", "o", "", "Output destination (format: type=local,dest=path)")
	image_buildCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	image_buildCmd.Flags().String("progress", "", "Set type of progress output (auto, plain, tty). Use plain to")
	image_buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image")
	image_buildCmd.Flags().BoolP("quiet", "q", false, "Suppress the build output and print image ID on success")
	image_buildCmd.Flags().Bool("rm", false, "Remove intermediate containers after a successful build")
	image_buildCmd.Flags().String("secret", "", "Secret file to expose to the build (only if BuildKit enabled):")
	image_buildCmd.Flags().String("security-opt", "", "Security options")
	image_buildCmd.Flags().String("shm-size", "", "Size of /dev/shm")
	image_buildCmd.Flags().Bool("squash", false, "Squash newly built layers into a single new layer")
	image_buildCmd.Flags().String("ssh", "", "SSH agent socket or keys to expose to the build (only if")
	image_buildCmd.Flags().Bool("stream", false, "Stream attaches to server to negotiate build context")
	image_buildCmd.Flags().StringP("tag", "t", "", "Name and optionally a tag in the 'name:tag' format")
	image_buildCmd.Flags().String("target", "", "Set the target build stage to build.")
	image_buildCmd.Flags().String("ulimit", "", "Ulimit options (default [])")
	imageCmd.AddCommand(image_buildCmd)

	rootAlias(image_buildCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"file":      carapace.ActionFiles(),
			"iidfile":   carapace.ActionFiles(),
			"isolation": carapace.ActionValues("default", "hyperv", "process"),
			"network":   carapace.ActionValues("bridge", "container", "host", "none"),
			"progress":  carapace.ActionValues("auto", "plain", "tty"),
		})

		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionFiles(),
		)
	})
}
