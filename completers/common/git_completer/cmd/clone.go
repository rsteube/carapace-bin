package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository into a new directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().Bool("bare", false, "create a bare repository")
	cloneCmd.Flags().StringP("branch", "b", "", "checkout <branch> instead of the remote's HEAD")
	cloneCmd.Flags().StringP("config", "c", "", "set config inside the new repository")
	cloneCmd.Flags().String("depth", "", "create a shallow clone of that depth")
	cloneCmd.Flags().Bool("dissociate", false, "use --reference only while cloning")
	cloneCmd.Flags().String("filter", "", "object filtering")
	cloneCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	cloneCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	cloneCmd.Flags().StringP("jobs", "j", "", "number of submodules cloned in parallel")
	cloneCmd.Flags().BoolP("local", "l", false, "to clone from a local repository")
	cloneCmd.Flags().Bool("mirror", false, "create a mirror repository (implies bare)")
	cloneCmd.Flags().BoolP("no-checkout", "n", false, "don't create a checkout")
	cloneCmd.Flags().Bool("no-hardlinks", false, "don't use local hardlinks, always copy")
	cloneCmd.Flags().Bool("no-tags", false, "don't clone any tags, and make later fetches not to follow them")
	cloneCmd.Flags().StringP("origin", "o", "", "use <name> instead of 'origin' to track upstream")
	cloneCmd.Flags().Bool("progress", false, "force progress reporting")
	cloneCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	cloneCmd.Flags().String("recurse-submodules", "", "initialize submodules in the clone")
	cloneCmd.Flags().String("recursive", "", "alias of --recurse-submodules")
	cloneCmd.Flags().String("reference", "", "reference repository")
	cloneCmd.Flags().String("reference-if-able", "", "reference repository")
	cloneCmd.Flags().Bool("remote-submodules", false, "any cloned submodules will use their remote-tracking branch")
	cloneCmd.Flags().String("separate-git-dir", "", "separate git dir from working tree")
	cloneCmd.Flags().String("server-option", "", "option to transmit")
	cloneCmd.Flags().String("shallow-exclude", "", "deepen history of shallow clone, excluding rev")
	cloneCmd.Flags().String("shallow-since", "", "create a shallow clone since a specific time")
	cloneCmd.Flags().Bool("shallow-submodules", false, "any cloned submodules will be shallow")
	cloneCmd.Flags().BoolP("shared", "s", false, "setup as shared repository")
	cloneCmd.Flags().Bool("single-branch", false, "clone only one branch, HEAD or --branch")
	cloneCmd.Flags().Bool("sparse", false, "initialize sparse-checkout file to include only files at root")
	cloneCmd.Flags().String("template", "", "directory from which templates will be used")
	cloneCmd.Flags().StringP("upload-pack", "u", "", "path to git-upload-pack on the remote")
	cloneCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"separate-git-dir": carapace.ActionFiles(),
		"template":         carapace.ActionDirectories(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		git.ActionRepositorySearch(),
		carapace.ActionDirectories(),
	)
}
