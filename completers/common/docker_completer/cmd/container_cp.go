package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var container_cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cpCmd).Standalone()

	container_cpCmd.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	container_cpCmd.Flags().BoolP("follow-link", "L", false, "Always follow symbol link in SRC_PATH")
	containerCmd.AddCommand(container_cpCmd)

	rootAlias(container_cpCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if util.HasPathPrefix(c.CallbackValue) {
					return carapace.ActionFiles()
				}
				return docker.ActionContainerPath()
			}),
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if util.HasPathPrefix(c.Args[0]) {
					return docker.ActionContainerPath()
				}
				return carapace.ActionFiles()
			}),
		)
	})
}
