package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysctl",
	Short: "configure kernel parameters at runtime",
	Long:  "https://linux.die.net/man/8/sysctl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "alias of -a")
	rootCmd.Flags().BoolS("X", "X", false, "alias of -a")
	rootCmd.Flags().BoolP("all", "a", false, "display all variables")
	rootCmd.Flags().BoolP("binary", "b", false, "print value without new line")
	rootCmd.Flags().BoolS("d", "d", false, "alias of -h")
	rootCmd.Flags().Bool("deprecated", false, "include deprecated parameters to listing")
	rootCmd.Flags().BoolS("f", "f", false, "alias of -p")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore", "e", false, "ignore unknown variables errors")
	rootCmd.Flags().StringP("load", "p", "", "read values from file")
	rootCmd.Flags().BoolP("names", "N", false, "print variable names without values")
	rootCmd.Flags().BoolS("o", "o", false, "does nothing")
	rootCmd.Flags().StringP("pattern", "r", "", "select setting that match expression")
	rootCmd.Flags().BoolP("quiet", "q", false, "do not echo variable set")
	rootCmd.Flags().Bool("system", false, "read values from all system directories")
	rootCmd.Flags().BoolP("values", "n", false, "print only values of the given variable(s)")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
	rootCmd.Flags().BoolP("write", "w", false, "enable writing a value to variable")
	rootCmd.Flags().BoolS("x", "x", false, "does nothing")

	rootCmd.Flag("load").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"load": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return ActionVariableNames().Invoke(c).ToMultiPartsA(".")
			case 1:
				return ActionVariableValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}

func ActionVariableNames() carapace.Action {
	return carapace.ActionExecCommand("sysctl", "--all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			if splitted := strings.SplitN(line, " = ", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionVariableValues(name string) carapace.Action {
	return carapace.ActionExecCommand("sysctl", "--values", name)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
