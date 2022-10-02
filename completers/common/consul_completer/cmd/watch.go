package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for changes in Consul",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()
	addClientFlags(watchCmd)
	addServerFlags(watchCmd)

	watchCmd.Flags().String("key", "", "Specifies the key to watch.")
	watchCmd.Flags().String("name", "", "")
	watchCmd.Flags().String("passingonly", "", "Specifies if only hosts passing all checks are displayed.")
	watchCmd.Flags().String("prefix", "", "Specifies the key prefix to watch.")
	watchCmd.Flags().String("service", "", "Specifies the service to watch.")
	watchCmd.Flags().Bool("shell", false, "Use a shell to run the command.")
	watchCmd.Flags().String("state", "", "Specifies the states to watch.")
	watchCmd.Flags().StringArray("tag", []string{}, "Specifies the service tag(s) to filter on.")
	watchCmd.Flags().String("type", "", "Specifies the watch type.")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).FlagCompletion(carapace.ActionMap{
		// TODO complete flags
		"passingonly": carapace.ActionValues("true", "false"),
		"type":        carapace.ActionValues("key", "keyprefix", "services", "nodes", "service", "checks", "event"),
	})

	// TODO positional completion
}
