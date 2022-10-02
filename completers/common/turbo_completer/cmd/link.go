package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your local directory to a Vercel organization and enable remote caching",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().String("api", "", "api url")
	linkCmd.Flags().Bool("no-gitignore", false, "Do not create or modify .gitignore (default false)")
	rootCmd.AddCommand(linkCmd)

	linkCmd.Flag("api").NoOptDefVal = " "
}
