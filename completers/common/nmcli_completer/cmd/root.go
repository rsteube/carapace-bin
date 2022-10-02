package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nmcli",
	Short: "command-line tool for controlling NetworkManager",
	Long:  "https://developer-old.gnome.org/NetworkManager/stable/nmcli.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("ask", "a", false, "ask for missing parameters")
	rootCmd.Flags().StringP("colors", "c", "", "whether to use colors in output")
	rootCmd.Flags().StringP("escape", "e", "", "escape columns separators in values")
	rootCmd.Flags().StringP("fields", "f", "", "specify fields to output")
	rootCmd.Flags().StringP("get-values", "g", "", "shortcut for -m tabular -t -f")
	rootCmd.Flags().BoolP("help", "h", false, "print this help")
	rootCmd.Flags().StringP("mode", "m", "", "output mode")
	rootCmd.Flags().BoolP("overview", "o", false, "overview mode")
	rootCmd.Flags().BoolP("pretty", "p", false, "pretty output")
	rootCmd.Flags().BoolP("show-secrets", "s", false, "allow displaying passwords")
	rootCmd.Flags().BoolP("terse", "t", false, "terse output")
	rootCmd.Flags().BoolP("version", "v", false, "show program version")
	rootCmd.Flags().StringP("wait", "w", "", "set timeout waiting for finishing operations")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"colors":     carapace.ActionValues("auto", "yes", "no").StyleF(style.ForKeyword),
		"escape":     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"fields":     ActionMultiFields(),
		"get-values": ActionMultiFields(),
		"mode":       carapace.ActionValues("tabular", "multiline"),
	})
}

func ActionMultiFields() carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return ActionFields().Invoke(c).Filter(c.Parts).Suffix(",").ToA()
	})
}

func ActionFields() carapace.Action {
	return carapace.ActionValues(
		"ACTIVE", "CONNECTION", "RATE", "TIMESTAMP",
		"ACTIVE-PATH", "CONNECTIVITY", "READONLY", "TIMESTAMP-REAL",
		"all", "CON-PATH", "RUNNING", "TYPE",
		"AUTOCONNECT", "CON-UUID", "SECURITY", "UUID",
		"AUTOCONNECT-PRIORITY", "DBUS-PATH", "SIGNAL", "VERSION",
		"BARS", "DEVICE", "SLAVE", "WIFI",
		"BSSID", "MODE", "SSID", "WIFI-HW",
		"CHAN", "NAME", "STARTUP", "WWAN",
		"common", "NETWORKING", "STATE", "WWAN-HW",
	)
}
