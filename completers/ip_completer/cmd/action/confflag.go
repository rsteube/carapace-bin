package action

import "github.com/rsteube/carapace"

func ActionConfFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"home", "(IPv6 only) designates this address the \"home address\"",
		"nodad", "(IPv6 only) do not perform Duplicate Address Detection",
		"mngtmpaddr", "(IPv6 only) make the kernel manage temporary addresses created from this one as template",
		"noprefixroute", "Do not automatically create a route for the network prefix of the added address",
		"autojoin", "enable autojoin",
	)
}
