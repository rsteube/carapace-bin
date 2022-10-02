package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func files(cmd *cobra.Command) (result []string) {
	result, _ = cmd.Root().Flags().GetStringArray("file")
	return
}

func ActionContainerUsers(cmd *cobra.Command, service string, index int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return compose.ActionUsers(compose.ContainerUserOpts{Files: files(cmd), Service: service, Index: index})
	})
}
