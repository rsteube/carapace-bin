package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_image_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Create or update an ImagePolicy object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_image_policyCmd).Standalone()
	create_image_policyCmd.Flags().String("filter-extract", "", "replacement pattern (using capture groups from --filter-regex) to use for sorting")
	create_image_policyCmd.Flags().String("filter-regex", "", "regular expression pattern used to filter the image tags")
	create_image_policyCmd.Flags().String("image-ref", "", "the name of an image repository object")
	create_image_policyCmd.Flags().String("select-alpha", "", "use alphabetical sorting to select image; either \"asc\" meaning select the last, or \"desc\" meaning select the first")
	create_image_policyCmd.Flags().String("select-numeric", "", "use numeric sorting to select image; either \"asc\" meaning select the last, or \"desc\" meaning select the first")
	create_image_policyCmd.Flags().String("select-semver", "", "a semver range to apply to tags; e.g., '1.x'")
	create_imageCmd.AddCommand(create_image_policyCmd)
}
