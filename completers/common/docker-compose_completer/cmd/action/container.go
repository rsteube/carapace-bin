package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func ActionContainers(cmd *cobra.Command, opts compose.ContainerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts.Files = files(cmd)
		return compose.ActionContainers(opts)
	})
}
