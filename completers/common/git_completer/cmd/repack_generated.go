package cmd

import (
	"github.com/spf13/cobra"
)

var repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Pack unpacked objects in a repository",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	repackCmd.Flags().BoolS("A", "A", false, "same as -a, and turn unreachable objects loose")
	repackCmd.Flags().BoolS("F", "F", false, "pass --no-reuse-object to git-pack-objects")
	repackCmd.Flags().BoolS("a", "a", false, "pack everything in a single pack")
	repackCmd.Flags().BoolS("d", "d", false, "remove redundant packs, and run git-prune-packed")
	repackCmd.Flags().BoolP("delta-islands", "i", false, "pass --delta-islands to git-pack-objects")
	repackCmd.Flags().String("depth", "", "limits the maximum delta depth")
	repackCmd.Flags().BoolS("f", "f", false, "pass --no-reuse-delta to git-pack-objects")
	repackCmd.Flags().String("keep-pack", "", "do not repack this pack")
	repackCmd.Flags().BoolP("keep-unreachable", "k", false, "with -a, repack unreachable objects")
	repackCmd.Flags().BoolP("local", "l", false, "pass --local to git-pack-objects")
	repackCmd.Flags().String("max-pack-size", "", "maximum size of each packfile")
	repackCmd.Flags().BoolS("n", "n", false, "do not run git-update-server-info")
	repackCmd.Flags().Bool("pack-kept-objects", false, "repack objects in packs marked with .keep")
	repackCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	repackCmd.Flags().String("threads", "", "limits the maximum number of threads")
	repackCmd.Flags().String("unpack-unreachable", "", "with -A, do not loosen objects older than this")
	repackCmd.Flags().String("window", "", "size of the window used for delta compression")
	repackCmd.Flags().String("window-memory", "", "same as the above, but limit memory size instead of entries count")
	repackCmd.Flags().BoolP("write-bitmap-index", "b", false, "write bitmap index")
	rootCmd.AddCommand(repackCmd)
}
