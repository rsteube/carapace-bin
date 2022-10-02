package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "aws",
	Short:              "Universal Command Line Interface for Amazon Web Services",
	Long:               "https://aws.amazon.com/cli/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			current := c.CallbackValue
			if c.CallbackValue == "-" {
				return carapace.ActionValues("--").NoSpace() // no shorthand flags so expand to longhand first (which is needed for the completer)
			}

			c.Setenv("COMP_LINE", "aws "+strings.Join(append(c.Args, current), " ")) // TODO escape/quote special characters
			return carapace.ActionExecCommand("aws_completer")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				if lines[0] == "" {
					return carapace.ActionValues()
				}
				for index, line := range lines {
					if strings.HasSuffix(line, " ") {
						lines[index] = strings.TrimSuffix(line, " ") // v1 has space suffix
					}
				}
				a := carapace.ActionValues(lines[:len(lines)-1]...)
				if strings.HasPrefix(current, "file://") ||
					strings.HasPrefix(current, "fileb://") {
					return a.NoSpace()
				}
				return a
			}).Invoke(c).ToA()
		}),
	)
}
