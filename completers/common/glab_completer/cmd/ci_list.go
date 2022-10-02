package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ci_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of CI pipelines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_listCmd).Standalone()
	ci_listCmd.Flags().StringP("orderBy", "o", "", "Order pipeline by <string>")
	ci_listCmd.Flags().IntP("page", "p", 1, "Page number")
	ci_listCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page. (default 30)")
	ci_listCmd.Flags().String("sort", "desc", "Sort pipeline by {asc|desc}. (Defaults to desc)")
	ci_listCmd.Flags().StringP("status", "s", "", "Get pipeline with status: {running|pending|success|failed|canceled|skipped|created|manual}")
	ciCmd.AddCommand(ci_listCmd)

	carapace.Gen(ci_listCmd).FlagCompletion(carapace.ActionMap{
		"orderBy": carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"sort":    carapace.ActionValues("asc", "desc"),
		"status":  carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual"),
	})
}
