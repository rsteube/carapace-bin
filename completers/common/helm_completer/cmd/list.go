package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list releases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "show all releases without any filter applied")
	listCmd.Flags().BoolP("all-namespaces", "A", false, "list releases across all namespaces")
	listCmd.Flags().BoolP("date", "d", false, "sort by release date")
	listCmd.Flags().Bool("deployed", false, "show deployed releases. If no other is specified, this will be automatically enabled")
	listCmd.Flags().Bool("failed", false, "show failed releases")
	listCmd.Flags().StringP("filter", "f", "", "a regular expression (Perl compatible). Any releases that match the expression will be included in the results")
	listCmd.Flags().IntP("max", "m", 256, "maximum number of releases to fetch")
	listCmd.Flags().Int("offset", 0, "next release index in the list, used to offset from start value")
	listCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	listCmd.Flags().Bool("pending", false, "show pending releases")
	listCmd.Flags().BoolP("reverse", "r", false, "reverse the sort order")
	listCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Works only for secret(default) and configmap storage backends.")
	listCmd.Flags().BoolP("short", "q", false, "output short (quiet) listing format")
	listCmd.Flags().Bool("superseded", false, "show superseded releases")
	listCmd.Flags().String("time-format", "", "format time using golang time formatter. Example: --time-format \"2006-01-02 15:04:05Z0700\"")
	listCmd.Flags().Bool("uninstalled", false, "show uninstalled releases (if 'helm uninstall --keep-history' was used)")
	listCmd.Flags().Bool("uninstalling", false, "show releases that are currently being uninstalled")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})
}
