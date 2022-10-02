package action

import "github.com/rsteube/carapace"

func ActionFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"BACKUP", "Mention files backed up",
		"COPY", "Mention files copied locally on the receiving side",
		"DEL", "Mention deletions on the receiving side",
		"FLIST", "Mention file-list receiving/sending (levels 1-2)",
		"MISC", "Mention miscellaneous information (levels 1-2)",
		"MOUNT", "Mention mounts that were found or skipped",
		"NAME", "Mention 1) updated file/dir names, 2) unchanged names",
		"PROGRESS", "Mention 1) per-file progress or 2) total transfer progress",
		"REMOVE", "Mention files removed on the sending side",
		"SKIP", "Mention files that are skipped due to options used",
		"STATS", "Mention statistics at end of run (levels 1-3)",
		"SYMSAFE", "Mention symlinks that are unsafe",
		"ALL", "Set all --info options (e.g. all4)",
		"NONE", "Silence all --info options (same as all0)",
		"HELP", "Output this help message",
	)
}
