package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nvim_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nvim",
	Short: "edit text",
	Long:  "https://neovim.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("M", "M", false, "Modifications in text not allowed")
	rootCmd.Flags().StringS("O", "O", "", "Open N vertical windows (default: one per file)")
	rootCmd.Flags().BoolS("R", "R", false, "Read-only mode")
	rootCmd.Flags().StringS("S", "S", "", "Source <session> after loading the first file")
	rootCmd.Flags().StringS("V", "V", "", "Verbose [level][file]")
	rootCmd.Flags().BoolS("Z", "Z", false, "Restricted mode")
	rootCmd.Flags().Bool("api-info", false, "Write msgpack-encoded API metadata to stdout")
	rootCmd.Flags().BoolS("b", "b", false, "Binary mode")
	rootCmd.Flags().String("cmd", "", "Execute <cmd> before any config")
	rootCmd.Flags().BoolS("d", "d", false, "Diff mode")
	rootCmd.Flags().Bool("embed", false, "Use stdin/stdout as a msgpack-rpc channel")
	rootCmd.Flags().Bool("headless", false, "Don't start a user interface")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help message")
	rootCmd.Flags().StringS("i", "i", "", "Use this shada file")
	rootCmd.Flags().String("listen", "", "Serve RPC API from this address")
	rootCmd.Flags().BoolS("m", "m", false, "Modifications (writing files) not allowed")
	rootCmd.Flags().BoolS("n", "n", false, "No swap file, use memory only")
	rootCmd.Flags().Bool("noplugin", false, "Don't load plugins")
	rootCmd.Flags().StringS("o", "o", "", "Open N windows (default: one per file)")
	rootCmd.Flags().StringS("p", "p", "", "Open N tab pages (default: one per file)")
	rootCmd.Flags().StringS("r", "r", "", "Recover edit state for this file")
	rootCmd.Flags().StringS("s", "s", "", "Read Normal mode commands from <scriptin>")
	rootCmd.Flags().String("startuptime", "", "Write startup timing messages to <file>")
	rootCmd.Flags().StringS("u", "u", "", "Use this config file")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"S":           carapace.ActionFiles(),
		"i":           carapace.ActionFiles(),
		"r":           action.ActionTemporaryFiles(),
		"startuptime": carapace.ActionFiles(),
		"u":           carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
