package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run command in the background.")
	execCmd.Flags().StringArrayP("env", "e", []string{}, "Set environment variables")
	execCmd.Flags().Int("index", 1, "index of the container if there are multiple instances of a service [default: 1].")
	execCmd.Flags().BoolP("no-TTY", "T", false, "Disable pseudo-TTY allocation. By default `docker compose exec` allocates a TTY.")
	execCmd.Flags().Bool("privileged", false, "Give extended privileges to the process.")
	execCmd.Flags().StringP("user", "u", "", "Run the command as this user.")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				if index, err := execCmd.Flags().GetInt("index"); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					return action.ActionContainerUsers(execCmd, c.Args[0], index)
				}
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionServices(execCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if index, err := execCmd.Flags().GetInt("index"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return action.ActionFiles(execCmd, c.Args[0], index)
			}
		}),
	)
}
