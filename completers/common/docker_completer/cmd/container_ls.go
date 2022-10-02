package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_lsCmd).Standalone()

	container_lsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	container_lsCmd.Flags().StringArrayP("filter", "f", []string{}, "Filter output based on conditions provided")
	container_lsCmd.Flags().String("format", "", "Pretty-print containers using a Go template")
	container_lsCmd.Flags().StringP("last", "n", "", "Show n last created containers (includes all states) (default -1)")
	container_lsCmd.Flags().BoolP("latest", "l", false, "Show the latest created container (includes all states)")
	container_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	container_lsCmd.Flags().BoolP("quiet", "q", false, "Only display numeric IDs")
	container_lsCmd.Flags().BoolP("size", "s", false, "Display total file sizes")
	containerCmd.AddCommand(container_lsCmd)

	rootAlias(container_lsCmd, func(cmd *cobra.Command, isAlias bool) {
		if isAlias {
			cmd.Use = "ps"
		}

		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"filter": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"ancestor", "Filters containers which share a given image as an ancestor",
						"before", "Filters containers created before or after a given container ID or name",
						"exited", "An integer representing the container’s exit code",
						"expose", "Filters containers which publish or expose a given port",
						"health", "Filters containers based on their healthcheck status",
						"id", "Container’s ID",
						"is-task", "Filters containers that are a “task” for a service",
						"isolation", "Windows daemon only",
						"label", "An arbitrary string representing either a key or a key-value pair",
						"name", "Container’s name",
						"network", "Filters running containers connected to a given network",
						"publish", "Filters containers which publish or expose a given port",
						"since", "Filters containers created before or after a given container ID or name",
						"status", "Container status",
						"volume", "	Filters running containers which have mounted a given volume or bind mount",
					).Invoke(c).Suffix("=").ToA()
				case 1:
					switch c.Parts[0] {
					case "ancestor":
						return docker.ActionRepositoryTags()
					case "before":
						return docker.ActionContainers()
					case "expose":
						return docker.ActionPorts()
					case "health":
						return carapace.ActionValues("starting", "healthy", "unhealthy", "none")
					case "is-task":
						return carapace.ActionValues("true", "false")
					case "id":
						return docker.ActionContainerIds()
					case "isolation":
						return carapace.ActionValues("default", "process", "hyperv")
					case "name":
						return docker.ActionContainers()
					case "network":
						return docker.ActionNetworks()
					case "publish":
						return docker.ActionPorts()
					case "since":
						return docker.ActionContainers()
					case "status":
						return carapace.ActionValues("created", "restarting", "removing", "running", "paused", "exited", "dead")
					case "volume":
						return docker.ActionVolumes()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			}),
		})
	})
}
