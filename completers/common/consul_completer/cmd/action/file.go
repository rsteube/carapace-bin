package action

import "github.com/rsteube/carapace"

func ActionOptionalFiles(suffix ...string) carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 1:
			return carapace.ActionFiles(suffix...)
		default:
			return carapace.ActionValues()
		}
	})
}
