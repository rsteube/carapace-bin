package cmd

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:   "--style [config]",
	Short: "set style",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := setStyle(args[0]); err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
		}
	},
}

func init() {
	carapace.Gen(styleCmd).Standalone()

	carapace.Gen(styleCmd).PositionalCompletion(
		carapace.ActionStyleConfig(),
	)
}

func setStyle(s string) error {
	if splitted := strings.SplitN(s, "=", 2); len(splitted) == 2 {
		return style.Set(splitted[0], splitted[1])
	}
	return fmt.Errorf("invalid format: '%v'", s)
}
