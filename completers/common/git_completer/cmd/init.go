package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Git repository or reinitialize an existing one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("bare", false, "create a bare repository")
	initCmd.Flags().StringP("initial-branch", "b", "", "override the name of the initial branch")
	initCmd.Flags().String("object-format", "", "specify the hash algorithm to use")
	initCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	initCmd.Flags().String("separate-git-dir", "", "separate git dir from working tree")
	initCmd.Flags().String("shared", "", "specify that the git repository is to be shared amongst several users")
	initCmd.Flags().String("template", "", "directory from which templates will be used")
	rootCmd.AddCommand(initCmd)

	initCmd.Flag("shared").NoOptDefVal = "umask"

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"object-format":    carapace.ActionValues("sha1", "sha256"),
		"separate-git-dir": carapace.ActionFiles(),
		"shared": carapace.ActionValuesDescribed(
			"false", "use permissions reported by umask",
			"true", "make the repository group-writable",
			"umask", "use permissions reported by umask",
			"group", "make the repository group-writable",
			"all", "make repository readable by all users",
			"world", "make repository readable by all users",
			"everybody", "make repository readable by all users",
		),
		"template": carapace.ActionDirectories(),
	})

	carapace.Gen(initCmd).PositionalCompletion(carapace.ActionFiles())
}
