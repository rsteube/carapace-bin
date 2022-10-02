package cmd

import (
	"github.com/spf13/cobra"
)

var show_branchCmd = &cobra.Command{
	Use:   "show-branch",
	Short: "Show branches and their commits",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	show_branchCmd.Flags().BoolP("all", "a", false, "show remote-tracking and local branches")
	show_branchCmd.Flags().String("color", "", "color '*!+-' corresponding to the branch")
	show_branchCmd.Flags().Bool("current", false, "include the current branch")
	show_branchCmd.Flags().Bool("date-order", false, "topologically sort, maintaining date order where possible")
	show_branchCmd.Flags().Bool("independent", false, "show refs unreachable from any other ref")
	show_branchCmd.Flags().Bool("list", false, "synonym to more=-1")
	show_branchCmd.Flags().Bool("merge-base", false, "show possible merge bases")
	show_branchCmd.Flags().String("more", "", "show <n> more commits after the common ancestor")
	show_branchCmd.Flags().Bool("no-name", false, "suppress naming strings")
	show_branchCmd.Flags().StringP("reflog", "g", "", "show <n> most recent ref-log entries starting at base")
	show_branchCmd.Flags().BoolP("remotes", "r", false, "show remote-tracking branches")
	show_branchCmd.Flags().Bool("sha1-name", false, "name commits with their object names")
	show_branchCmd.Flags().Bool("sparse", false, "show merges reachable from only one tip")
	show_branchCmd.Flags().Bool("topics", false, "show only commits not on the first branch")
	show_branchCmd.Flags().Bool("topo-order", false, "show commits in topological order")
	rootCmd.AddCommand(show_branchCmd)
}
