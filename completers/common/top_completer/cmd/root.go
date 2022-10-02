package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/top_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "top",
	Short: "display Linux processes",
	Long:  "https://linux.die.net/man/1/top",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "Starts top with the last remembered Cpu States portion of the  summary  area  reversed.")
	rootCmd.Flags().StringS("E", "E", "", "Instructs top to force summary area memory to be scaled as")
	rootCmd.Flags().BoolS("H", "H", false, "Instructs  top  to  display  individual  threads.")
	rootCmd.Flags().StringS("O", "O", "", "print available field names")
	rootCmd.Flags().BoolS("S", "S", false, "Starts  top  with the last remembered `S' state reversed.")
	rootCmd.Flags().StringS("U", "U", "", "Display  only  processes  with  a  user  id or user name matching that given.")
	rootCmd.Flags().BoolS("b", "b", false, "Starts top in Batch mode")
	rootCmd.Flags().BoolS("c", "c", false, "Starts top with the last remembered `c' state reversed")
	rootCmd.Flags().StringS("d", "d", "", "Specifies the delay between screen updates")
	rootCmd.Flags().StringS("e", "e", "", "Instructs top to force task area memory to be scaled as")
	rootCmd.Flags().BoolS("h", "h", false, "show help")
	rootCmd.Flags().BoolS("i", "i", false, "Starts top with the last remembered `i' state reversed")
	rootCmd.Flags().StringS("n", "n", "", "Specifies  the  maximum  number  of  iterations,  or  frames, top should produce before ending.")
	rootCmd.Flags().StringS("o", "o", "", "Specifies the name of the field on which tasks will be sorted")
	rootCmd.Flags().StringS("p", "p", "", "Monitor  only  processes with specified process IDs")
	rootCmd.Flags().BoolS("s", "s", false, "Starts top with secure mode forced, even for root.")
	rootCmd.Flags().StringS("u", "u", "", "Display  only  processes  with  a  user  id or user name matching that given.")
	rootCmd.Flags().BoolS("v", "v", false, "show version")
	rootCmd.Flags().BoolS("w", "w", false, "val    In Batch mode, when used without an argument top will format output using the  COLUMNS= and  LINES=  environment  variables, if  set.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"E": carapace.ActionValuesDescribed(
			"k", "kibibytes",
			"m", "mebibytes",
			"g", "gibibytes",
			"t", "tebibytes",
			"p", "pebibytes",
			"e", "exbibytes",
		),
		"U": os.ActionUsers(),
		"e": carapace.ActionValuesDescribed(
			"k", "kibibytes",
			"m", "mebibytes",
			"g", "gibibytes",
			"t", "tebibytes",
			"p", "pebibytes",
		),
		"o": action.ActionFields(),
		"p": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"u": os.ActionUsers(),
	})
}
