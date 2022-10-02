package action

import "github.com/rsteube/carapace"

func ActionProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"MODAL", "makes the window into a modal",
		"STICKY", "makes the window appear on all workspaces",
		"MAXIMIZED_VERT", "sizes the window maximized vertically",
		"MAXIMIZED_HORZ", "sizes the window maximized horizontally",
		"ABOVE", "Show window above all others (always on top)",
		"BELOW", "Show window below all others",
		"SKIP_TASKBAR", "hides the window from the taskbar",
		"SKIP_PAGER", "hides the window from the window pager",
		"FULLSCREEN", "makes window fullscreen",
		"HIDDEN", "unmaps the window",
		"SHADED", "rolls the window up",
		"DEMANDS_ATTENTION", "marks window urgent or needing attention",
	)
}
