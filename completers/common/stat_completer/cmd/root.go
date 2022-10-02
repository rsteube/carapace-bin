package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stat",
	Short: "display file or file system status",
	Long:  "https://linux.die.net/man/2/stat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cached", "", "specify how to use cached attributes;")
	rootCmd.Flags().BoolP("dereference", "L", false, "follow links")
	rootCmd.Flags().BoolP("file-system", "f", false, "display file system status instead of file status")
	rootCmd.Flags().StringP("format", "c", "", "use the specified FORMAT instead of the default;")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("printf", "", "like --format, but interpret backslash escapes,")
	rootCmd.Flags().BoolP("terse", "t", false, "print the information in terse form")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cached": carapace.ActionValuesDescribed(
			"always", "use cached attributes if available",
			"never", "try to synchronize with the latest attributes",
			"default", "leave it up to the underlying file system",
		).StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
