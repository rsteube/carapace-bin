package cmd

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:                "git",
	Short:              "execute a git command",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(gitCmd).Standalone()

	rootCmd.AddCommand(gitCmd)

	carapace.Gen(gitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			home, err := os.UserHomeDir()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			location := fmt.Sprintf("%v/.password-store", home)
			if dir, ok := os.LookupEnv("PASSWORD_STORE_DIR"); ok {
				location = dir
			}
			return bridge.ActionCarapaceBin("git").Chdir(location)
		}),
	)
}
