package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-doc",
	Short: "show documentation for package or symbol",
	Long:  "https://pkg.go.dev/cmd/doc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	rootCmd.Flags().BoolS("all", "all", false, "Show all the documentation for the package")
	rootCmd.Flags().BoolS("c", "c", false, "Respect case when matching symbols")
	rootCmd.Flags().BoolS("cmd", "cmd", false, "Treat a command like a regular package")
	rootCmd.Flags().BoolS("short", "short", false, "One-line representation for each symbol")
	rootCmd.Flags().BoolS("src", "src", false, "Show the full source code for the symbol")
	rootCmd.Flags().BoolS("u", "u", false, "Show documentation for unexported as well")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return golang.ActionPackages()
			default:
				// TODO complete versions for other packages
				var url string
				switch {
				case strings.HasPrefix(c.Parts[0], "github.com/"):
					if splitted := strings.Split(c.Parts[0], "/"); len(splitted) > 2 {
						url = "https://" + strings.Join(splitted[:3], "/")
					}
				case strings.HasPrefix(c.Parts[0], "golang.org/x/tools/"):
					url = "https://github.com/golang/tools"
				}

				if url != "" {
					return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: url, Tags: true})
				}
				return carapace.ActionValues()
			}
		}),
		carapace.ActionMultiPartsN(".", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return golang.ActionSymbols(golang.SymbolOpts{
					Package:    c.Args[0],
					Unexported: rootCmd.Flag("u").Changed,
				}).NoSpace()

			default:
				return golang.ActionMethodOrFields(golang.MethodOrFieldOpts{
					Package:    c.Args[0],
					Symbol:     c.Parts[0],
					Unexported: rootCmd.Flag("u").Changed,
				})
			}
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})
}
