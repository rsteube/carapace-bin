package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsmem",
	Short: "list the ranges of available memory with their online status",
	Long:  "https://www.man7.org/linux/man-pages/man1/lsmem.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "list each individual memory block")
	rootCmd.Flags().BoolP("bytes", "b", false, "print SIZE in bytes rather than in human readable format")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().StringP("output", "o", "", "output columns")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().BoolP("pairs", "P", false, "use key=\"value\" output format")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format")
	rootCmd.Flags().StringP("split", "S", "", "split ranges by specified columns")
	rootCmd.Flags().String("summary", "", "print summary information")
	rootCmd.Flags().StringP("sysroot", "s", "", "use the specified directory as system root")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("summary").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"split": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"summary": carapace.ActionValues("never", "always", "only").StyleF(style.ForKeyword),
		"sysroot": carapace.ActionDirectories(),
	})
}

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"RANGE", "start and end address of the memory range",
		"SIZE", "size of the memory range",
		"STATE", "online status of the memory range",
		"REMOVABLE", "memory is removable",
		"BLOCK", "memory block number or blocks range",
		"NODE", "numa node of memory",
		"ZONES", "valid zones for the memory range",
	)
}
