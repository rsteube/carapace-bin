package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify configuration values in .condarc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().String("add", "", "Add one configuration value to the beginning of a list key.")
	configCmd.Flags().String("append", "", "Add one configuration value to the end of a list key.")
	configCmd.Flags().String("describe", "", "Describe given configuration parameters.")
	configCmd.Flags().Bool("env", false, "Write to the active conda environment .condarc file")
	configCmd.Flags().String("file", "", "Write to the given file.")
	configCmd.Flags().String("get", "", "Get a configuration value.")
	configCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	configCmd.Flags().Bool("json", false, "Report all output as json.")
	configCmd.Flags().String("prepend", "", "Add one configuration value to the beginning of a list key.")
	configCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	configCmd.Flags().String("remove", "", "Remove a configuration value from a list key.")
	configCmd.Flags().String("remove-key", "", "Remove a configuration key (and all its values).")
	configCmd.Flags().String("set", "", "Set a boolean or string key")
	configCmd.Flags().String("show", "", "Display configuration values as calculated and compiled.")
	configCmd.Flags().Bool("show-sources", false, "Display all identified configuration sources.")
	configCmd.Flags().Bool("stdin", false, "Apply configuration information given in yaml format piped through stdin.")
	configCmd.Flags().Bool("system", false, "Write to the system .condarc file at '/opt/miniconda3/.condarc'.")
	configCmd.Flags().Bool("validate", false, "Validate all configuration sources.")
	configCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	configCmd.Flags().Bool("write-default", false, "Write the default configuration to a file.")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"add":        action.ActionConfigKeys(),
		"append":     action.ActionConfigKeys(),
		"describe":   action.ActionConfigKeys(),
		"file":       carapace.ActionFiles(),
		"get":        action.ActionConfigKeys(),
		"prepend":    action.ActionConfigKeys(),
		"remove":     action.ActionConfigKeys(),
		"remove-key": action.ActionConfigKeys(),
		"set":        action.ActionConfigKeys(),
		"show":       action.ActionConfigKeys(),
	})

}
