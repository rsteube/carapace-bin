package action

import "github.com/rsteube/carapace"

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"NS", "namespace identifier (inode number)",
		"TYPE", "kind of namespace",
		"PATH", "path to the namespace",
		"NPROCS", "number of processes in the namespace",
		"PID", "lowest PID in the namespace",
		"PPID", "PPID of the PID",
		"COMMAND", "command line of the PID",
		"UID", "UID of the PID",
		"USER", "username of the PID",
		"NETNSID", "namespace ID as used by network subsystem",
		"NSFS", "nsfs mountpoint (usually used network subsystem)",
		"PNS", "parent namespace identifier (inode number)",
		"ONS", "owner namespace identifier (inode number)",
	)
}
