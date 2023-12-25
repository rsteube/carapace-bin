package cmd

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "--run",
	Short: "run spec",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, spec, err := loadSpec(args[0])
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
			os.Exit(1)
		}

		specCmd := spec.ToCobra()

		specCmd.SetArgs(args[1:])
		specCmd.Execute() // TODO handle error?
	},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(".yaml"),
	)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return spec.ActionSpec(c.Args[0]).Shift(1) // TODO
		}),
	)
}
