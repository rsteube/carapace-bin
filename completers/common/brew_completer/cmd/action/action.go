package action

import (
	"github.com/spf13/cobra"
)

func defaultArgs(cmd *cobra.Command) []string {
	args := make([]string, 0)

	if flag := cmd.Flag("formula"); flag != nil && flag.Changed {
		args = append(args, "--formula")
	} else if flag := cmd.Flag("formulae"); flag != nil && flag.Changed {
		args = append(args, "--formulae")
	}

	if flag := cmd.Flag("cask"); flag != nil && flag.Changed {
		args = append(args, "--cask")
	} else if flag := cmd.Flag("casks"); flag != nil && flag.Changed {
		args = append(args, "--casks")
	}
	return args
}
