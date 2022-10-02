package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionRepo() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if util.HasPathPrefix(c.CallbackValue) {
			return carapace.ActionDirectories()
		}
		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"sftp", "SFTP",
					"rest", "HTTP",
					"s3", "Amazon S3",
					"swift", "OpenStack Swift",
					"b2", "Backblaze B2",
					"azure", "Microsoft Azure Blob Storage",
					"gs", "Google Cloud Storage",
					"rclone", "rclone",
				).Invoke(c).Suffix(":").ToA()
			default:
				// TODO completion
				return carapace.ActionValues()
			}
		})
	})
}
