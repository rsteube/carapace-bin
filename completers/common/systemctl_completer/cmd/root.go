package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/journalctl"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "systemctl",
	Short: "Query or send control commands to the system manager",
	Long:  "https://systemd.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("P", "P", "", "Equivalent to --value --property=NAME")
	rootCmd.Flags().BoolP("all", "a", false, "Show all properties/all units currently in memory, including dead/empty ones")
	rootCmd.Flags().String("boot-loader-entry", "", "Boot into a specific boot loader entry on next boot")
	rootCmd.Flags().String("boot-loader-menu", "", "Boot into boot loader menu on next boot")
	rootCmd.Flags().String("check-inhibitors", "", "Specify if checking inhibitors before shutting down, sleeping or hibernating")
	rootCmd.Flags().Bool("dry-run", false, "Only print what would be done")
	rootCmd.Flags().Bool("failed", false, "Shortcut for --state=failed")
	rootCmd.Flags().Bool("firmware-setup", false, "Tell the firmware to show the setup menu on next boot")
	rootCmd.Flags().BoolP("force", "f", false, "override existing symlinks or execute action immediately")
	rootCmd.Flags().BoolP("full", "l", false, "Don't ellipsize unit names on output")
	rootCmd.Flags().Bool("global", false, "Enable/disable/mask default user unit files globally")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().StringP("host", "H", "", "Operate on remote host")
	rootCmd.Flags().BoolS("i", "i", false, "Shortcut for --check-inhibitors=no")
	rootCmd.Flags().String("job-mode", "", "Specify how to deal with already queued jobs, when queueing a new job")
	rootCmd.Flags().String("kill-who", "", "Whom to send signal to")
	rootCmd.Flags().String("legend", "", "Enable/disable the legend (column headers and hints)")
	rootCmd.Flags().StringP("lines", "n", "", "Number of journal entries to show")
	rootCmd.Flags().Bool("marked", false, "Restart/reload previously marked units")
	rootCmd.Flags().Bool("mkdir", false, "Create directory before mounting, if missing")
	rootCmd.Flags().Bool("no-ask-password", false, "Do not ask for system passwords")
	rootCmd.Flags().Bool("no-block", false, "Do not wait until operation finished")
	rootCmd.Flags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.Flags().Bool("no-reload", false, "Don't reload daemon after en-/dis-abling unit files")
	rootCmd.Flags().Bool("no-wall", false, "Don't send wall message before halt/power-off/reboot")
	rootCmd.Flags().Bool("now", false, "Start or stop unit after enabling or disabling it")
	rootCmd.Flags().StringP("output", "o", "", "Change journal output mode")
	rootCmd.Flags().Bool("plain", false, "Print unit dependencies as a list instead of a tree")
	rootCmd.Flags().StringP("property", "p", "", "Show only properties by this name")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	rootCmd.Flags().Bool("read-only", false, "Create read-only bind mount")
	rootCmd.Flags().BoolP("recursive", "r", false, "Show unit list of host and local containers")
	rootCmd.Flags().Bool("reverse", false, "Show reverse dependencies with 'list-dependencies'")
	rootCmd.Flags().String("root", "", "Enable/disable/mask unit files in the specified root directory")
	rootCmd.Flags().Bool("runtime", false, "Enable/disable/mask unit files temporarily until next reboot")
	rootCmd.Flags().BoolP("show-transaction", "T", false, "When enqueuing a unit job, show full transaction")
	rootCmd.Flags().Bool("show-types", false, "When showing sockets, explicitly show their type")
	rootCmd.Flags().StringP("signal", "s", "", "Which signal to send")
	rootCmd.Flags().String("state", "", "List units with particular LOAD or SUB or ACTIVE state")
	rootCmd.Flags().Bool("system", false, "Connect to system manager")
	rootCmd.Flags().String("timestamp", "", "Change format of printed timestamps.")
	rootCmd.Flags().StringP("type", "t", "", "List units of a particular type")
	rootCmd.Flags().Bool("user", false, "Connect to user service manager")
	rootCmd.Flags().Bool("value", false, "When showing properties, only print the value")
	rootCmd.Flags().Bool("version", false, "Show package version")
	rootCmd.Flags().Bool("wait", false, "wait until service stopped or startup completed")
	rootCmd.Flags().String("what", "", "Which types of resources to remove")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"P":                action.ActionProperties(),
		"check-inhibitors": carapace.ActionValues("auto", "yes", "no").StyleF(style.ForKeyword),
		"host":             net.ActionHosts(),
		"job-mode": carapace.ActionValuesDescribed(
			"fail", "cause fail on conflict",
			"replace", "replace on conflict",
			"replace-irreversibly", "replace on conflict irreversibly",
			"isolate", "cause all other units to be stopped when starting unit",
			"ignore-dependencies", "ignore all unit dependencies",
			"ignore-requirements", "ignore all unit requirements",
			"flush", "cancel all queued jobs",
			"triggering", "stop unit and any active unit that may trigger it",
		),
		"kill-who": carapace.ActionValuesDescribed(
			"main", "main process",
			"control", "controll process",
			"all", "all processes",
		),
		"output":   journalctl.ActionOutputs(),
		"property": action.ActionProperties(),
		"root":     carapace.ActionDirectories(),
		"signal":   ps.ActionKillSignals(),
		"state":    action.ActionStates(),
		"timestamp": carapace.ActionValuesDescribed(
			"pretty", "Day YYYY-MM-DD HH:MM:SS TZ",
			"us", "Day YYYY-MM-DD HH:MM:SS.UUUUUU TZ",
			"utc", "Day YYYY-MM-DD HH:MM:SS UTC",
			"us+utc", "Day YYYY-MM-DD HH:MM:SS.UUUUUU UTC",
		),
		"type": action.ActionTypes(),
		"what": carapace.ActionValues("configuration", "state", "cache", "logs", "runtime"),
	})
}
