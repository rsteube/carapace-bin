package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Modify device properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_setCmd).Standalone()

	deviceCmd.AddCommand(device_setCmd)

	carapace.Gen(device_setCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
		carapace.ActionValues("autoconnect", "managed"),
		action.ActionYesNo(),
		carapace.ActionValues("autoconnect", "managed"),
		action.ActionYesNo(),
	)
}
