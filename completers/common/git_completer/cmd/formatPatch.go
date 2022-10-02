package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var formatPatchCmd = &cobra.Command{
	Use:   "format-patch",
	Short: "Prepare patches for e-mail submission",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatPatchCmd).Standalone()

	formatPatchCmd.Flags().String("add-header", "", "add email header")
	formatPatchCmd.Flags().String("attach", "", "attach the patch")
	formatPatchCmd.Flags().String("base", "", "add prerequisite tree info to the patch series")
	formatPatchCmd.Flags().String("cc", "", "add Cc: header")
	formatPatchCmd.Flags().String("cover-from-description", "", "generate parts of a cover letter based on a branch's description")
	formatPatchCmd.Flags().Bool("cover-letter", false, "generate a cover letter")
	formatPatchCmd.Flags().String("creation-factor", "", "percentage by which creation is weighted ")
	formatPatchCmd.Flags().String("filename-max-length", "", "max length of output filename")
	formatPatchCmd.Flags().String("from", "", "set From address to <ident> (or committer ident if absent)")
	formatPatchCmd.Flags().Bool("ignore-if-in-upstream", false, "don't include a patch matching a commit upstream")
	formatPatchCmd.Flags().String("in-reply-to", "", "make first mail a reply to <message-id>")
	formatPatchCmd.Flags().String("inline", "", "inline the patch")
	formatPatchCmd.Flags().String("interdiff", "", "show changes against <rev> in cover letter or single patch")
	formatPatchCmd.Flags().BoolP("keep-subject", "k", false, "don't strip/add [PATCH]")
	formatPatchCmd.Flags().Bool("no-binary", false, "don't output binary diffs")
	formatPatchCmd.Flags().BoolP("no-numbered", "N", false, "use [PATCH] even with multiple patches")
	formatPatchCmd.Flags().BoolP("no-stat", "p", false, "show patch format instead of default (patch + stat)")
	formatPatchCmd.Flags().BoolP("numbered", "n", false, "use [PATCH n/m] even with a single patch")
	formatPatchCmd.Flags().Bool("numbered-files", false, "use simple number sequence for output file names")
	formatPatchCmd.Flags().StringP("output-directory", "o", "", "store resulting files in <dir>")
	formatPatchCmd.Flags().Bool("progress", false, "show progress while generating patches")
	formatPatchCmd.Flags().BoolP("quiet", "q", false, "don't print the patch filenames")
	formatPatchCmd.Flags().String("range-diff", "", "show changes against <refspec> in cover letter or single patch")
	formatPatchCmd.Flags().StringP("reroll-count", "v", "", "mark the series as Nth re-roll")
	formatPatchCmd.Flags().Bool("rfc", false, "use [RFC PATCH] instead of [PATCH]")
	formatPatchCmd.Flags().String("signature", "", "add a signature")
	formatPatchCmd.Flags().String("signature-file", "", "add a signature from a file")
	formatPatchCmd.Flags().BoolP("signoff", "s", false, "add a Signed-off-by trailer")
	formatPatchCmd.Flags().String("start-number", "", "start numbering patches at <n> instead of 1")
	formatPatchCmd.Flags().Bool("stdout", false, "print patches to standard out")
	formatPatchCmd.Flags().String("subject-prefix", "", "use [<prefix>] instead of [PATCH]")
	formatPatchCmd.Flags().String("suffix", "", "use <sfx> instead of '.patch'")
	formatPatchCmd.Flags().String("thread", "", "enable message threading, styles: shallow, deep")
	formatPatchCmd.Flags().String("to", "", "add To: header")
	formatPatchCmd.Flags().Bool("zero-commit", false, "output all-zero hash in From header")

	formatPatchCmd.Flag("from").NoOptDefVal = " "
	formatPatchCmd.Flag("attach").NoOptDefVal = " "
	formatPatchCmd.Flag("inline").NoOptDefVal = " "
	formatPatchCmd.Flag("thread").NoOptDefVal = " "

	rootCmd.AddCommand(formatPatchCmd)

	carapace.Gen(formatPatchCmd).FlagCompletion(carapace.ActionMap{
		"output-directory": carapace.ActionDirectories(),
		"thread":           carapace.ActionValues("shallow", "deep"),
	})

	carapace.Gen(formatPatchCmd).PositionalCompletion(
		carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionRefs(git.RefOption{}.Default())
			case 1:
				return git.ActionRefs(git.RefOption{}.Default())
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
