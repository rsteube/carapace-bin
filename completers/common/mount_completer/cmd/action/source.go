package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionSources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if util.HasPathPrefix(c.CallbackValue) {
			return carapace.Batch(
				fs.ActionBlockDevices(),
				carapace.ActionFiles(),
			).ToA()
		}
		return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"LABEL", "specifies device by filesystem label",
					"UUID", "specifies device by filesystem UUID",
					"PARTLABEL", "specifies device by partition label",
					"PARTUUID", "specifies device by partition UUID",
					"ID", "specifies device by udev hardware ID",
				).Invoke(c).Suffix("=").ToA()
			case 1:
				switch c.Parts[0] {
				case "LABEL":
					return fs.ActionLabels()
				case "UUID":
					return fs.ActionUuids()
				case "PARTLABEL":
					return fs.ActionPartLabels()
				case "PARTUUID":
					return fs.ActionPartUuids()
				case "ID":
					// TODO
					return carapace.ActionValues()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		})
	})
}
