package cmd

import (
	"github.com/spf13/cobra"
)

var pack_objectsCmd = &cobra.Command{
	Use:   "pack-objects",
	Short: "Create a packed archive of objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	pack_objectsCmd.Flags().Bool("all", false, "include objects reachable from any reference")
	pack_objectsCmd.Flags().Bool("all-progress", false, "show progress meter during object writing phase")
	pack_objectsCmd.Flags().Bool("all-progress-implied", false, "similar to --all-progress when progress meter is shown")
	pack_objectsCmd.Flags().String("compression", "", "pack compression level")
	pack_objectsCmd.Flags().Bool("delta-base-offset", false, "use OFS_DELTA objects")
	pack_objectsCmd.Flags().Bool("delta-islands", false, "respect islands during delta compression")
	pack_objectsCmd.Flags().String("depth", "", "maximum length of delta chain allowed in the resulting pack")
	pack_objectsCmd.Flags().Bool("exclude-promisor-objects", false, "do not pack objects in promisor packfiles")
	pack_objectsCmd.Flags().String("filter", "", "object filtering")
	pack_objectsCmd.Flags().Bool("honor-pack-keep", false, "ignore packs that have companion .keep file")
	pack_objectsCmd.Flags().Bool("include-tag", false, "include tag objects that refer to objects to be packed")
	pack_objectsCmd.Flags().Bool("incremental", false, "ignore packed objects")
	pack_objectsCmd.Flags().String("index-version", "", "write the pack index file in the specified idx format version")
	pack_objectsCmd.Flags().Bool("indexed-objects", false, "include objects referred to by the index")
	pack_objectsCmd.Flags().String("keep-pack", "", "ignore this pack")
	pack_objectsCmd.Flags().Bool("keep-true-parents", false, "do not hide commits by grafts")
	pack_objectsCmd.Flags().Bool("keep-unreachable", false, "keep unreachable objects")
	pack_objectsCmd.Flags().Bool("local", false, "ignore borrowed objects from alternate object store")
	pack_objectsCmd.Flags().String("max-pack-size", "", "maximum size of each output pack file")
	pack_objectsCmd.Flags().String("missing", "", "handling for missing objects")
	pack_objectsCmd.Flags().Bool("non-empty", false, "do not create an empty pack output")
	pack_objectsCmd.Flags().Bool("pack-loose-unreachable", false, "pack loose unreachable objects")
	pack_objectsCmd.Flags().Bool("progress", false, "show progress meter")
	pack_objectsCmd.Flags().BoolP("quiet", "q", false, "do not show progress meter")
	pack_objectsCmd.Flags().Bool("reflog", false, "include objects referred by reflog entries")
	pack_objectsCmd.Flags().Bool("reuse-delta", false, "reuse existing deltas")
	pack_objectsCmd.Flags().Bool("reuse-object", false, "reuse existing objects")
	pack_objectsCmd.Flags().Bool("revs", false, "read revision arguments from standard input")
	pack_objectsCmd.Flags().Bool("shallow", false, "create packs suitable for shallow fetches")
	pack_objectsCmd.Flags().Bool("sparse", false, "use the sparse reachability algorithm")
	pack_objectsCmd.Flags().Bool("stdout", false, "output pack to stdout")
	pack_objectsCmd.Flags().Bool("thin", false, "create thin packs")
	pack_objectsCmd.Flags().String("threads", "", "use threads when searching for best delta matches")
	pack_objectsCmd.Flags().String("unpack-unreachable", "", "unpack unreachable objects newer than <time>")
	pack_objectsCmd.Flags().Bool("unpacked", false, "limit the objects to those that are not yet packed")
	pack_objectsCmd.Flags().Bool("use-bitmap-index", false, "use a bitmap index if available to speed up counting objects")
	pack_objectsCmd.Flags().String("window", "", "limit pack window by objects")
	pack_objectsCmd.Flags().String("window-memory", "", "limit pack window by memory in addition to object limit")
	pack_objectsCmd.Flags().Bool("write-bitmap-index", false, "write a bitmap index together with the pack index")
	rootCmd.AddCommand(pack_objectsCmd)
}
