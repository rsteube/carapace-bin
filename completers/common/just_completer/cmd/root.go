package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/just"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "just",
	Short: "Just a command runner",
	Long:  "https://github.com/casey/just",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("changelog", false, "Print changelog")
	rootCmd.Flags().Bool("check", false, "Run `--fmt` in 'check' mode. Exits with 0 if justfile is formatted correctly. Exits with 1 and prints a diff if formatting is required.")
	rootCmd.Flags().Bool("choose", false, "Select one or more recipes to run using a binary. If `--chooser` is not passed the chooser defaults to the value of $JUST_CHOOSER, falling back to `fzf`")
	rootCmd.Flags().String("chooser", "", "Override binary invoked by `--choose`")
	rootCmd.Flags().Bool("clear-shell-args", false, "Clear shell arguments")
	rootCmd.Flags().String("color", "", "Print colorful output [default: auto]  [possible values: auto, always, never]")
	rootCmd.Flags().StringP("command", "c", "", "Run an arbitrary command with the working directory, `.env`, overrides, and exports set")
	rootCmd.Flags().String("completions", "", "Print shell completion script for <SHELL> [possible values: zsh, bash, fish, powershell, elvish]")
	rootCmd.Flags().String("dotenv-filename", "", "Search for environment file named <DOTENV-FILENAME> instead of `.env`")
	rootCmd.Flags().String("dotenv-path", "", "Load environment file at <DOTENV-PATH> instead of searching for one")
	rootCmd.Flags().Bool("dry-run", false, "Print what just would do without doing it")
	rootCmd.Flags().Bool("dump", false, "Print justfile")
	rootCmd.Flags().String("dump-format", "", "Dump justfile as <FORMAT> [default: just]  [possible values: just, json]")
	rootCmd.Flags().BoolP("edit", "e", false, "Edit justfile with editor given by $VISUAL or $EDITOR, falling back to `vim`")
	rootCmd.Flags().Bool("evaluate", false, "Evaluate and print all variables. If a variable name is given as an argument, only print that variable's value.")
	rootCmd.Flags().Bool("fmt", false, "Format and overwrite justfile")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().Bool("highlight", false, "Highlight echoed recipe lines in bold")
	rootCmd.Flags().Bool("init", false, "Initialize new justfile in project root")
	rootCmd.Flags().StringP("justfile", "f", "", "Use <JUSTFILE> as justfile")
	rootCmd.Flags().BoolP("list", "l", false, "List available recipes and their arguments")
	rootCmd.Flags().String("list-heading", "", "Print <TEXT> before list")
	rootCmd.Flags().String("list-prefix", "", "Print <TEXT> before each list item")
	rootCmd.Flags().Bool("no-dotenv", false, "Don't load `.env` file")
	rootCmd.Flags().Bool("no-highlight", false, "Don't highlight echoed recipe lines in bold")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress all output")
	rootCmd.Flags().String("set", "", "Override <VARIABLE> with <VALUE>")
	rootCmd.Flags().String("shell", "", "Invoke <SHELL> to run recipes")
	rootCmd.Flags().String("shell-arg", "", "Invoke shell with <SHELL-ARG> as an argument")
	rootCmd.Flags().Bool("shell-command", false, "Invoke <COMMAND> with the shell used to run recipe lines and backticks")
	rootCmd.Flags().StringP("show", "s", "", "Show information about <RECIPE>")
	rootCmd.Flags().Bool("summary", false, "List names of available recipes")
	rootCmd.Flags().BoolP("unsorted", "u", false, "Return list and summary entries in source order")
	rootCmd.Flags().Bool("unstable", false, "Enable unstable features")
	rootCmd.Flags().Bool("variables", false, "List names of variables")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")
	rootCmd.Flags().StringP("working-directory", "d", "", "Use <WORKING-DIRECTORY> as working directory. --justfile must also be set")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chooser": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"color":       carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"completions": carapace.ActionValues("zsh", "bash", "fish", "powershell", "elvish"),
		"dotenv-path": carapace.ActionDirectories(),
		"dump-format": carapace.ActionValues("just", "json"),
		"justfile":    carapace.ActionFiles(),
		"set": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionVariables(rootCmd.Flag("justfile").Value.String())
		}),
		"shell": os.ActionShells(),
		"show": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionRecipes(rootCmd.Flag("justfile").Value.String())
		}),
		"working-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionRecipes(rootCmd.Flag("justfile").Value.String()).Invoke(c).Filter(c.Args).ToA()
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		if f := rootCmd.Flag("working-directory"); f.Changed {
			return action.Chdir(f.Value.String())
		}
		return action
	})
}
