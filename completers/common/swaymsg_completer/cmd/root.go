package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaymsg",
	Short: "Send messages to a running instance of sway over the IPC socket",
	Long:  "https://github.com/swaywm/sway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().BoolP("monitor", "m", false, "Monitor until killed (-t SUBSCRIBE only)")
	rootCmd.Flags().BoolP("pretty", "p", false, "Use pretty output even when not using a tty")
	rootCmd.Flags().BoolP("quiet", "q", false, "Be quiet.")
	rootCmd.Flags().BoolP("raw", "r", false, "Use raw output even if using a tty")
	rootCmd.Flags().StringP("socket", "s", "", "Use the specified socket.")
	rootCmd.Flags().StringP("type", "t", "", "Specify the message type.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"socket": carapace.ActionFiles(),
		"type": carapace.ActionValuesDescribed(
			"<command>", "The message is a sway command",
			"get_workspaces", "Gets a list of workspaces and their status.",
			"get_inputs", "Gets a list of current inputs.",
			"get_outputs", "Gets a list of current outputs.",
			"get_tree", "Gets a JSON-encoded layout tree of all open windows, containers, outputs, workspaces",
			"get_seats", "Gets a list of all seats, its properties and all assigned devices.",
			"get_marks", "Get a JSON-encoded list of marks.",
			"get_bar_config", "Get a JSON-encoded configuration for swaybar.",
			"get_version", "Get version information for the running instance of sway.",
			"get_binding_modes", "Gets a JSON-encoded list of currently configured binding modes.",
			"get_binding_state", "Gets JSON-encoded info about the current binding state.",
			"get_config", "Gets a copy of the current configuration. Doesn't expand includes.",
			"send_tick", "Sends a tick event to all subscribed clients.",
			"subscribe", "Subscribe to a list of event types.",
		),
	})
}
