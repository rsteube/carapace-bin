package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var general_loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "Get or change NetworkManager logging level and domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_loggingCmd).Standalone()

	generalCmd.AddCommand(general_loggingCmd)

	carapace.Gen(general_loggingCmd).PositionalCompletion(
		carapace.ActionValues("level"),
		carapace.ActionValuesDescribed(
			"ERR", "log only critical errors",
			"WARN", "log warning messages",
			"INFO", "log various informational messages",
			"DEBUG", "verbose logging for debugging purposes",
		),
		carapace.ActionValues("domains"),
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionDomains().Invoke(c).Filter(c.Parts).Suffix(",").ToA()
		}),
	)
}

func ActionDomains() carapace.Action {
	return carapace.ActionValues(
		"PLATFORM", "RFKILL", "ETHER", "WIFI", "BT", "MB", "DHCP4", "DHCP6", "PPP", "IP4", "IP6", "AUTOIP4", "DNS", "VPN", "SHARING", "SUPPLICANT", "AGENTS", "SETTINGS", "SUSPEND", "CORE", "DEVICE", "OLPC", "INFINIBAND", "FIREWALL", "ADSL", "BOND", "VLAN", "BRIDGE", "TEAM", "CONCHECK", "DCB", "DISPATCH", "AUDIT", "SYSTEMD",
	)
}
