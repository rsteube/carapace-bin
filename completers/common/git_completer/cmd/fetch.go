package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download objects and refs from another repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("all", false, "fetch from all remotes")
	fetchCmd.Flags().BoolP("append", "a", false, "append to .git/FETCH_HEAD instead of overwriting")
	fetchCmd.Flags().Bool("auto-gc", false, "run 'gc --auto' after fetching")
	fetchCmd.Flags().String("deepen", "", "deepen history of shallow clone")
	fetchCmd.Flags().String("depth", "", "deepen history of shallow clone")
	fetchCmd.Flags().Bool("dry-run", false, "dry run")
	fetchCmd.Flags().String("filter", "", "object filtering")
	fetchCmd.Flags().BoolP("force", "f", false, "force overwrite of local reference")
	fetchCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	fetchCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	fetchCmd.Flags().StringP("jobs", "j", "", "number of submodules fetched in parallel")
	fetchCmd.Flags().BoolP("keep", "k", false, "keep downloaded pack")
	fetchCmd.Flags().BoolP("multiple", "m", false, "fetch from multiple remotes")
	fetchCmd.Flags().BoolS("n", "n", false, "do not fetch all tags (--no-tags)")
	fetchCmd.Flags().String("negotiation-tip", "", "report that we have only objects reachable from this object")
	fetchCmd.Flags().Bool("progress", false, "force progress reporting")
	fetchCmd.Flags().BoolP("prune", "p", false, "prune remote-tracking branches no longer on remote")
	fetchCmd.Flags().BoolP("prune-tags", "P", false, "prune local tags no longer on remote and clobber changed tags")
	fetchCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	fetchCmd.Flags().String("recurse-submodules", "", "control recursive fetching of submodules")
	fetchCmd.Flags().String("refmap", "", "specify fetch refmap")
	fetchCmd.Flags().StringP("server-option", "o", "", "option to transmit")
	fetchCmd.Flags().Bool("set-upstream", false, "set upstream for git pull/fetch")
	fetchCmd.Flags().String("shallow-exclude", "", "deepen history of shallow clone, excluding rev")
	fetchCmd.Flags().String("shallow-since", "", "deepen history of shallow repository based on time")
	fetchCmd.Flags().Bool("show-forced-updates", false, "check for forced-updates on all updated branches")
	fetchCmd.Flags().BoolP("tags", "t", false, "fetch all tags and associated objects")
	fetchCmd.Flags().Bool("unshallow", false, "convert to a complete repository")
	fetchCmd.Flags().BoolP("update-head-ok", "u", false, "allow updating of HEAD ref")
	fetchCmd.Flags().Bool("update-shallow", false, "accept refs that update .git/shallow")
	fetchCmd.Flags().String("upload-pack", "", "path to upload pack on remote end")
	fetchCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	fetchCmd.Flags().Bool("write-commit-graph", false, "write the commit-graph after fetching")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !fetchCmd.Flag("all").Changed {
				return git.ActionRemotes().Invoke(c).Filter(c.Args).ToA()
			} else {
				return carapace.ActionValues()
			}
		}),
	)
}
