package cmd

import (
	"github.com/spf13/cobra"
)

var fast_exportCmd = &cobra.Command{
	Use:   "fast-export",
	Short: "Git data exporter",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	fast_exportCmd.Flags().Bool("anonymize", false, "anonymize output")
	fast_exportCmd.Flags().String("export-marks", "", "Dump marks to this file")
	fast_exportCmd.Flags().Bool("fake-missing-tagger", false, "Fake a tagger when tags lack one")
	fast_exportCmd.Flags().Bool("full-tree", false, "Output full tree for each commit")
	fast_exportCmd.Flags().String("import-marks", "", "Import marks from this file")
	fast_exportCmd.Flags().String("import-marks-if-exists", "", "Import marks from this file if it exists")
	fast_exportCmd.Flags().Bool("mark-tags", false, "Label tags with mark ids")
	fast_exportCmd.Flags().Bool("no-data", false, "Skip output of blob data")
	fast_exportCmd.Flags().String("progress", "", "show progress after <n> objects")
	fast_exportCmd.Flags().String("reencode", "", "select handling of commit messages in an alternate encoding")
	fast_exportCmd.Flags().Bool("reference-excluded-parents", false, "Reference parents which are not in fast-export stream by object id")
	fast_exportCmd.Flags().String("refspec", "", "Apply refspec to exported refs")
	fast_exportCmd.Flags().Bool("show-original-ids", false, "Show original object ids of blobs/commits")
	fast_exportCmd.Flags().String("signed-tags", "", "select handling of signed tags")
	fast_exportCmd.Flags().String("tag-of-filtered-object", "", "select handling of tags that tag filtered objects")
	fast_exportCmd.Flags().Bool("use-done-feature", false, "Use the done feature to terminate the stream")
	rootCmd.AddCommand(fast_exportCmd)
}
