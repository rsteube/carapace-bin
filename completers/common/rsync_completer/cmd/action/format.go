package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasSuffix(c.CallbackValue, "%") {
			return carapace.ActionValuesDescribed(
				"a", "the remote IP address",
				"b", "the number of bytes actually transferred",
				"B", "the permission bits of the file",
				"c", "the total size of the block checksums received for the basis file",
				"C", "the full-file checksum if it is known for the  file",
				"f", "the filename",
				"G", "the gid of the file",
				"h", "the remote host name",
				"i", "an itemized list of what is being updated",
				"l", "the length of the file in bytes",
				"L", "the string \" -> SYMLINK\", \" => HARDLINK\", or \"\"",
				"m", "the module name",
				"M", "the last-modified time of the file",
				"n", "the filename",
				"o", "the operation",
				"p", "the process ID of this rsync session",
				"P", "the module path",
				"t", "the current date time",
				"u", "the authenticated username or an empty string",
				"U", "the uid of the file",
			).Invoke(c).Prefix(c.CallbackValue).ToA().NoSpace()
		}
		return carapace.ActionValues()
	})
}
