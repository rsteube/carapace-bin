package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shred",
	Short: "overwrite a file to hide its contents, and optionally delete it",
	Long:  "https://linux.die.net/man/1/shred",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("exact", "x", false, "do not round file sizes up to the next full block;")
	rootCmd.Flags().BoolP("force", "f", false, "change permissions to allow writing if necessary")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("iterations", "n", "", "overwrite N times instead of the default (3)")
	rootCmd.Flags().String("random-source", "", "get random bytes from FILE")
	rootCmd.Flags().String("remove", "", "like -u but give control on HOW to delete;  See below")
	rootCmd.Flags().StringP("size", "s", "", "shred this many bytes (suffixes like K, M, G accepted)")
	rootCmd.Flags().BoolS("u", "u", false, "deallocate and remove file after overwriting")
	rootCmd.Flags().BoolP("verbose", "v", false, "show progress")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero", "z", false, "add a final overwrite with zeros to hide shredding")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"random-source": carapace.ActionFiles(),
		"remove": carapace.ActionValuesDescribed(
			"unlink", "use standard unlink call",
			"wipe", "like unlink, but obfuscate bytes in name first",
			"wipesync", "like wipe, but sync each obfuscated byte to disk",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
