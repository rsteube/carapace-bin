package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var project_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create or update a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_applyCmd).Standalone()

	project_applyCmd.Flags().Bool("app-status-poll", false, "Enable polling to continuously generate status reports for apps.")
	project_applyCmd.Flags().String("app-status-poll-interval", "", "Interval between polling to generate status reports if polling is enabled.")
	project_applyCmd.Flags().String("data-source", "", "The data source type to use.")
	project_applyCmd.Flags().String("from-waypoint-hcl", "", "waypoint.hcl formatted file to load settings from.")
	project_applyCmd.Flags().String("git-auth-type", "", "Authentication type for Git.")
	project_applyCmd.Flags().String("git-password", "", "Password for authentication when git-auth-type is 'basic'.")
	project_applyCmd.Flags().String("git-path", "", "Path is a subdirectory within the checked out repository to go into for the project's configuration.")
	project_applyCmd.Flags().String("git-private-key-password", "", "Password for the private key specified by git-private-key-path.")
	project_applyCmd.Flags().String("git-private-key-path", "", "Path to a PEM-encoded private key for 'ssh'-based auth.")
	project_applyCmd.Flags().String("git-ref", "", "Git ref to clone on new operations.")
	project_applyCmd.Flags().String("git-url", "", "URL of the Git repository to clone.")
	project_applyCmd.Flags().String("git-username", "", "Username for authentication when git-auth-type is 'basic'.")
	project_applyCmd.Flags().Bool("poll", false, "Enable polling.")
	project_applyCmd.Flags().String("poll-interval", "", "Interval between polling if polling is enabled.")
	project_applyCmd.Flags().String("runner-profile", "", "Name of a runner profile to be used for this project")
	project_applyCmd.Flags().String("waypoint-hcl", "", "Path to a waypoint.hcl file to associate with this project.")

	addGlobalOptions(project_applyCmd)

	projectCmd.AddCommand(project_applyCmd)

	carapace.Gen(project_applyCmd).FlagCompletion(carapace.ActionMap{
		"data-source":          carapace.ActionValues("git", "local"),
		"from-waypoint-hcl":    carapace.ActionFiles(".hcl"),
		"git-auth-type":        carapace.ActionValues("basic", "ssh"),
		"git-private-key-path": carapace.ActionFiles(),
		"git-ref": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := project_applyCmd.Flag("git-url"); flag.Changed {
				git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: flag.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"waypoint-hcl": carapace.ActionFiles(".hcl"),
	})

	carapace.Gen(project_applyCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
