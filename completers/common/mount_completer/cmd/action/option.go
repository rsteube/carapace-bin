package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
)

func ActionOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"async", "All I/O to the filesystem should be done asynchronously.",
				"atime", "Do not use the noatime feature.",
				"noatime", "Do not update inode access times on this filesystem.",
				"auto", "Can be mounted with the -a option.",
				"noauto", "Can only be mounted explicitly.",
				"context=", "Set context",
				"fscontext=", "Set context",
				"defcontext=", "Set context",
				"rootcontext=", "Set context",
				"defaults", "Use the default options",
				"dev", "Interpret character or block special devices on the filesystem.",
				"nodev", "Do not interpret character or block special devices on the filesystem.",
				"diratime", "Update directory inode access times on this filesystem.",
				"nodiratime", "Do not update directory inode access times on this filesystem.",
				"dirsync", "All directory updates within the filesystem should be done synchronously.",
				"exec", "Permit execution of binaries.",
				"noexec", "Do not permit direct execution of any binaries on the mounted filesystem.",
				"group", "Allow an ordinary user to mount the filesystem if one of that user’s groups matches the group of the device.",
				"iversion", "Every time the inode is modified, the i_version field will be incremented.",
				"noiversion", "Do not increment the i_version inode field.",
				"mand", "Allow mandatory locks on this filesystem.",
				"nomand", "Do not allow mandatory locks on this filesystem.",
				"_netdev", "The filesystem resides on a device that requires network access",
				"nofail", "Do not report errors for this device if it does not exist.",
				"relatime", "Update inode access times relative to modify or change time.",
				"norelatime", "Do not use the relatime feature.",
				"strictatime", "Allows to explicitly request full atime updates.",
				"nostrictatime", "Use the kernel’s default behavior for inode access time updates.",
				"lazytime", "Only update times on the in-memory version of the file inode.",
				"nolazytime", "Do not use the lazytime feature.",
				"suid", "Honor set-user-ID and set-group-ID bits or file capabilities",
				"nosuid", "Do not honor set-user-ID and set-group-ID bits or file capabilities",
				"silent", "Turn on the silent flag.",
				"loud", "Turn off the silent flag.",
				"owner", "Allow an ordinary user to mount the filesystem if that user is the owner of the device.",
				"remount", "Attempt to remount an already-mounted filesystem.",
				"ro", "Mount the filesystem read-only.",
				"rw", "Mount the filesystem read-write.",
				"sync", "All I/O to the filesystem should be done synchronously.",
				"user", "Allow an ordinary user to mount the filesystem.",
				"nouser", "Forbid an ordinary user to mount the filesystem.",
				"users", "Allow any user to mount and to unmount the filesystem",
				"X-mount.mkdir", "Allow to make a target directory if it does not exist yet.",
				"X-mount.subdir=", "Allow mounting sub-directory from a filesystem instead of the root directory.",
				"nosymfollow", "Do not follow symlinks when resolving paths.",
			)
		case 1:
			switch c.Parts[0] {
			case "X-mount.mkdir":
				return fs.ActionFileModesNumeric()

			default:
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	})
}
