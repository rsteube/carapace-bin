package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func ActionFiles(cmd *cobra.Command, service string, index int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return compose.ActionFiles(compose.ServicePathOpts{Files: files(cmd), Service: service, Index: index})
	})
}
