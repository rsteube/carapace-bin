package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create, list, delete or verify a tag object signed with GPG",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	tagCmd.Flags().BoolP("annotate", "a", false, "annotated tag, needs a message")
	tagCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	tagCmd.Flags().String("color", "", "respect format colors")
	tagCmd.Flags().String("column", "", "show tag list in columns")
	tagCmd.Flags().String("contains", "", "print only tags that contain the commit")
	tagCmd.Flags().Bool("create-reflog", false, "create a reflog")
	tagCmd.Flags().BoolP("delete", "d", false, "delete tags")
	tagCmd.Flags().BoolP("edit", "e", false, "force edit of tag message")
	tagCmd.Flags().StringP("file", "F", "", "read message from file")
	tagCmd.Flags().BoolP("force", "f", false, "replace the tag if exists")
	tagCmd.Flags().String("format", "", "format to use for the output")
	tagCmd.Flags().BoolP("ignore-case", "i", false, "sorting and filtering are case insensitive")
	tagCmd.Flags().BoolP("list", "l", false, "list tag names")
	tagCmd.Flags().StringP("local-user", "u", "", "use another key to sign the tag")
	tagCmd.Flags().String("merged", "", "print only tags that are merged")
	tagCmd.Flags().StringP("message", "m", "", "tag message")
	tagCmd.Flags().BoolS("n", "n", false, "print <n> lines of each tag message")
	tagCmd.Flags().String("no-contains", "", "print only tags that don't contain the commit")
	tagCmd.Flags().String("no-merged", "", "print only tags that are not merged")
	tagCmd.Flags().String("points-at", "", "print only tags of the object")
	tagCmd.Flags().BoolP("sign", "s", false, "annotated and GPG-signed tag")
	tagCmd.Flags().String("sort", "", "field name to sort on")
	tagCmd.Flags().BoolP("verify", "v", false, "verify tags")
	rootCmd.AddCommand(tagCmd)

	tagCmd.Flag("color").NoOptDefVal = " "

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":     git.ActionCleanupMode(),
		"color":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"contains":    git.ActionRefs(git.RefOption{Commits: 100}),
		"file":        carapace.ActionFiles(),
		"local-user":  os.ActionGpgKeyIds(),
		"merged":      git.ActionRefs(git.RefOption{Commits: 100}),
		"no-contains": git.ActionRefs(git.RefOption{Commits: 100}),
		"no-merged":   git.ActionRefs(git.RefOption{Commits: 100}),
		"points-at":   git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Commits: 100}),
	})

	carapace.Gen(tagCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if tagCmd.Flag("delete").Changed ||
				tagCmd.Flag("list").Changed ||
				tagCmd.Flag("verify").Changed {
				return git.ActionRefs(git.RefOption{Tags: true}).Invoke(c).Filter(c.Args).ToA()
			}
			switch len(c.Args) {
			case 0:
				return git.ActionRefs(git.RefOption{Tags: true})
			case 1:
				return git.ActionRefs(git.RefOption{LocalBranches: true, Commits: 100})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
