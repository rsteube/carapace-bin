package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var connection_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_addCmd).Standalone()

	connectionCmd.AddCommand(connection_addCmd)

	carapace.Gen(connection_addCmd).PositionalCompletion(
		carapace.ActionValues("type"),
		ActionTypes(),
	)

	carapace.Gen(connection_addCmd).PositionalAnyCompletion(ActionOption())
}

func ActionTypes() carapace.Action {
	vals := make([]string, 0)
	for key := range options {
		if key != "common" {
			vals = append(vals, key)
		}
	}
	return carapace.ActionValues(vals...)
}

func ActionOption() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		typeOptions := map[string]carapace.Action{}
		for key, value := range options["common"] { // add common options
			typeOptions[key] = value
		}
		for key, value := range options[c.Args[1]] { // add type options (assuming type is declared at arg[1])
			typeOptions[key] = value
		}
		for index, arg := range c.Args { // remove already declared options (but leave the last/current one)
			if index%2 == 0 && index < len(c.Args)-1 {
				delete(typeOptions, arg)
			}
		}

		if len(c.Args)%2 == 0 { // complete option names
			vals := make([]string, 0)
			for key := range typeOptions {
				vals = append(vals, key)
			}
			return carapace.ActionValues(vals...)
		} else { // complete option values
			optionName := c.Args[len(c.Args)-1]
			if action, ok := typeOptions[optionName]; !ok {
				return carapace.ActionMessage("unknown option: " + optionName)
			} else {
				return action
			}
		}
	})
}

var options = map[string]map[string]carapace.Action{
	"common": {
		"ifname":      carapace.ActionValues(),
		"con-name":    carapace.ActionValues(),
		"autoconnect": action.ActionYesNo(),
		"save":        action.ActionYesNo(),
		"master":      carapace.ActionValues(),
		"slave-type":  carapace.ActionValues(),
	},
	"ethernet": {
		"mac":        carapace.ActionValues(),
		"cloned-mac": carapace.ActionValues(),
		"mtu":        carapace.ActionValues(),
	},
	"wifi": {
		"ssid":       carapace.ActionValues(),
		"mac":        carapace.ActionValues(),
		"cloned-mac": carapace.ActionValues(),
		"mtu":        carapace.ActionValues(),
		"mode":       carapace.ActionValues("infrastructure", "ap", "adhoc"),
	},
	"wimax": {
		"mac": carapace.ActionValues(),
		"nsp": carapace.ActionValues(),
	},
	"pppoe": {
		"username": carapace.ActionValues(),
		"password": carapace.ActionValues(),
		"service":  carapace.ActionValues(),
		"mtu":      carapace.ActionValues(),
		"mac":      carapace.ActionValues(),
	},
	"gsm": {
		"apn":      carapace.ActionValues(),
		"user":     carapace.ActionValues(),
		"password": carapace.ActionValues(),
	},
	"cdma": {
		"user":     carapace.ActionValues(),
		"password": carapace.ActionValues(),
	},
	"infiniband": {
		"mac":            carapace.ActionValues(),
		"mtu":            carapace.ActionValues(),
		"transport-mode": carapace.ActionValues("datagram", "connected"),
		"parent":         carapace.ActionValues(), // TODO ifname
		"p-key":          carapace.ActionValues(),
	},
	"bluetooth": {
		"addr":    carapace.ActionValues(),
		"bt-type": carapace.ActionValues("panu", "nap", "dun-gsm", "dun-cdma"),
	},
	"vlan": {
		"dev":     carapace.ActionValues(),
		"id":      carapace.ActionValues(),
		"flags":   carapace.ActionValues(),
		"ingress": carapace.ActionValues(),
		"egress":  carapace.ActionValues(),
		"mtu":     carapace.ActionValues(),
	},
	"bond": {
		"mode":          carapace.ActionValues("balance-rr", "active-backup", "balance-xor", "broadcast", "802.3ad", "balance-tlb", "balance-alb"),
		"primary":       carapace.ActionValues(), // TODO ifname
		"miimon":        carapace.ActionValues(),
		"downdelay":     carapace.ActionValues(),
		"updelay":       carapace.ActionValues(),
		"arp-interval":  carapace.ActionValues(),
		"arp-ip-target": carapace.ActionValues(),
		"lacp-rate":     carapace.ActionValues("slow", "fast"),
	},
	"bond-slave": {
		"master": carapace.ActionValues(),
	},
	"team": {
		"config": carapace.ActionFiles(),
	},
	"team-slave": {
		"master": carapace.ActionValues(),
		"config": carapace.ActionFiles(),
	},
	"bridge": {
		"stp":                action.ActionYesNo(),
		"priority":           carapace.ActionValues(),
		"forward-delay":      carapace.ActionValues(),
		"hello-time":         carapace.ActionValues(),
		"max-age":            carapace.ActionValues(),
		"ageing-time":        carapace.ActionValues(),
		"multicast-snooping": action.ActionYesNo(),
		"mac":                carapace.ActionValues(),
	},
	"bridge-slave": {
		"master":    carapace.ActionValues(),
		"priority":  carapace.ActionValues(),
		"path-cost": carapace.ActionValues(),
		"hairpin":   action.ActionYesNo(),
	},
	"vpn": {
		"vpn-type": carapace.ActionValues("vpnc", "openvpn", "pptp", "openconnect", "openswan", "libreswan", "ssh", "l2tp", "iodine"),
		"user":     carapace.ActionValues(),
	},
	"olpc-mesh": {
		"ssid":         carapace.ActionValues(),
		"channel":      carapace.ActionValues(),
		"dhcp-anycast": carapace.ActionValues(),
	},
	"adsl": {
		"username":      carapace.ActionValues(),
		"protocol":      carapace.ActionValues("pppoa", "pppoe", "ipoatm"),
		"password":      carapace.ActionValues(),
		"encapsulation": carapace.ActionValues("vcmux", "llc"),
	},
	"tun": {
		"mode":        carapace.ActionValues("tun", "tap"),
		"owner":       os.ActionUsers(),
		"group":       os.ActionGroups(),
		"pi":          action.ActionYesNo(),
		"vnet-hdr":    action.ActionYesNo(),
		"multi-queue": action.ActionYesNo(),
	},
	"ip-tunnel": {
		"mode":   carapace.ActionValues("ipip", "gre", "sit", "isatap", "vti", "ip6ip6", "ipip6", "ip6gre", "vti6"),
		"remote": carapace.ActionValues(),
		"local":  carapace.ActionValues(),
		"dev":    carapace.ActionValues(),
	},
	"macsec": {
		"dev":     carapace.ActionValues(),
		"mode":    carapace.ActionValues("psk", "eap"),
		"cak":     carapace.ActionValues(),
		"encrypt": action.ActionYesNo(),
		"port":    carapace.ActionValues(),
	},
	"macvlan": {
		"dev":  carapace.ActionValues(),
		"mode": carapace.ActionValues("vepa", "bridge", "private", "passthru", "source"),
		"tap":  action.ActionYesNo(),
	},
	"vxlan": {
		"id":               carapace.ActionValues(),
		"remote":           carapace.ActionValues(),
		"local":            carapace.ActionValues(),
		"dev":              carapace.ActionValues(),
		"source-port-min":  carapace.ActionValues(),
		"source-port-max":  carapace.ActionValues(),
		"destination-port": carapace.ActionValues(),
	},
	"wpan": {
		"short-addr": carapace.ActionValues(),
		"pan-id":     carapace.ActionValues(),
		"page":       carapace.ActionValues(),
		"channel":    carapace.ActionValues(),
		"mac":        carapace.ActionValues(),
	},
	"6lowpan": {
		"dev": carapace.ActionValues(),
	},
	"dummy": {},
}

/**
  // TODO are these dependant on a specific option?

  SLAVE_OPTIONS:
    bridge:       [priority <0-63>]
                  [path-cost <1-65535>]
                  [hairpin yes|no]

    team:         [config <file>|<raw JSON data>]

  IP_OPTIONS:
                  [ip4 <IPv4 address>] [gw4 <IPv4 gateway>]
                  [ip6 <IPv6 address>] [gw6 <IPv6 gateway>]
**/
