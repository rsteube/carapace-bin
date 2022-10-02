package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-fmt",
	Short: "format all bin and lib files of the current crate",
	Long:  "https://github.com/rust-lang/rustfmt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Format all packages (only usable in workspaces)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().String("manifest-path", "", "Specify path to Cargo.toml")
	rootCmd.Flags().String("message-format", "", "Specify message-format: short|json|human")
	rootCmd.Flags().StringP("package", "p", "", "Specify package to format (only usable in workspaces)")
	rootCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	rootCmd.Flags().Bool("version", false, "Print rustfmt version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path":  carapace.ActionFiles(".toml"),
		"message-format": carapace.ActionValues("short", "json", "human"),
	})
}
