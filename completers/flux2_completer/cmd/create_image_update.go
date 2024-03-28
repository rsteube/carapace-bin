package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Create or update an ImageUpdateAutomation object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_image_updateCmd).Standalone()
	create_image_updateCmd.Flags().String("author-email", "", "the email to use for commit author")
	create_image_updateCmd.Flags().String("author-name", "", "the name to use for commit author")
	create_image_updateCmd.Flags().String("checkout-branch", "", "the branch to checkout")
	create_image_updateCmd.Flags().String("commit-template", "", "a template for commit messages")
	create_image_updateCmd.Flags().String("git-repo-path", "", "path to the directory containing the manifests to be updated, defaults to the repository root")
	create_image_updateCmd.Flags().String("git-repo-ref", "", "the name of a GitRepository resource with details of the upstream Git repository")
	create_image_updateCmd.Flags().String("push-branch", "", "the branch to push commits to, defaults to the checkout branch if not specified")
	create_imageCmd.AddCommand(create_image_updateCmd)
}
