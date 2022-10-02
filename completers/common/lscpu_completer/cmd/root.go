package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lscpu",
	Short: "display information about the CPU architecture",
	Long:  "https://linux.die.net/man/1/lscpu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print both online and offline CPUs (default for -e)")
	rootCmd.Flags().BoolP("bytes", "B", false, "print sizes in bytes rather than in human readable format")
	rootCmd.Flags().StringP("caches", "C", "", "info about caches in extended readable format")
	rootCmd.Flags().StringP("extended", "e", "", "print out an extended readable format")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("hex", "x", false, "print hexadecimal masks rather than lists of CPUs")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON for default or extended format")
	rootCmd.Flags().BoolP("offline", "c", false, "print offline CPUs only")
	rootCmd.Flags().BoolP("online", "b", false, "print online CPUs only (default for -p)")
	rootCmd.Flags().Bool("output-all", false, "print all available columns for -e, -p or -C")
	rootCmd.Flags().StringP("parse", "p", "", "print out a parsable format")
	rootCmd.Flags().BoolP("physical", "y", false, "print physical instead of logical IDs")
	rootCmd.Flags().StringP("sysroot", "s", "", "use specified directory as system root")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("caches").NoOptDefVal = " "
	rootCmd.Flag("extended").NoOptDefVal = " "
	rootCmd.Flag("parse").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"caches": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionCacheColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"extended": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFormatColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"parse": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFormatColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sysroot": carapace.ActionDirectories(),
	})
}

func ActionFormatColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"CPU", "logical CPU number",
		"CORE", "logical core number",
		"SOCKET", "logical socket number",
		"NODE", "logical NUMA node number",
		"BOOK", "logical book number",
		"DRAWER", "logical drawer number",
		"CACHE", "shows how caches are shared between CPUs",
		"POLARIZATION", "CPU dispatching mode on virtual hardware",
		"ADDRESS", "physical address of a CPU",
		"CONFIGURED", "shows if the hypervisor has allocated the CPU",
		"ONLINE", "shows if Linux currently makes use of the CPU",
		"MAXMHZ", "shows the maximum MHz of the CPU",
		"MINMHZ", "shows the minimum MHz of the CPU",
	)
}

func ActionCacheColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ALL-SIZE", "size of all system caches",
		"LEVEL", "cache level",
		"NAME", "cache name",
		"ONE-SIZE", "size of one cache",
		"TYPE", "cache type",
		"WAYS", "ways of associativity",
		"ALLOC-POLICY", "allocation policy",
		"WRITE-POLICY", "write policy",
		"PHY-LINE", "number of physical cache line per cache t",
		"SETS", "number of sets in the cache; set lines has the same cache index",
		"COHERENCY-SIZE", "minimum amount of data in bytes transferred from memory to cache",
	)
}
