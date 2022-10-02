package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate a formula or a cask for the downloadable file at URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("HEAD", false, "Indicate that URL points to the package's repository rather than a file.")
	createCmd.Flags().Bool("autotools", false, "Create a basic template for an Autotools-style build.")
	createCmd.Flags().Bool("cask", false, "Create a basic template for a cask.")
	createCmd.Flags().Bool("cmake", false, "Create a basic template for a CMake-style build.")
	createCmd.Flags().Bool("crystal", false, "Create a basic template for a Crystal build.")
	createCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	createCmd.Flags().BoolP("force", "f", false, "Ignore errors for disallowed formula names")
	createCmd.Flags().Bool("go", false, "Create a basic template for a Go build.")
	createCmd.Flags().BoolP("help", "h", false, "Show this message.")
	createCmd.Flags().Bool("meson", false, "Create a basic template for a Meson-style build.")
	createCmd.Flags().Bool("no-fetch", false, "Homebrew will not download URL to the cache")
	createCmd.Flags().Bool("node", false, "Create a basic template for a Node build.")
	createCmd.Flags().Bool("perl", false, "Create a basic template for a Perl build.")
	createCmd.Flags().Bool("python", false, "Create a basic template for a Python build.")
	createCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	createCmd.Flags().Bool("ruby", false, "Create a basic template for a Ruby build.")
	createCmd.Flags().Bool("rust", false, "Create a basic template for a Rust build.")
	createCmd.Flags().Bool("set-license", false, "Explicitly set the license of the new formula.")
	createCmd.Flags().Bool("set-name", false, "Explicitly set the name of the new formula or cask.")
	createCmd.Flags().Bool("set-version", false, "Explicitly set the version of the new formula or cask.")
	createCmd.Flags().Bool("tap", false, "Generate the new formula within the given tap")
	createCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(createCmd)
}
