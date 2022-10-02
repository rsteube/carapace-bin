package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uname",
	Short: "print system information",
	Long:  "https://linux.die.net/man/1/uname",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print all information, in the following order,")
	rootCmd.Flags().BoolP("hardware-platform", "i", false, "print the hardware platform (non-portable)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("kernel-name", "s", false, "print the kernel name")
	rootCmd.Flags().BoolP("kernel-release", "r", false, "print the kernel release")
	rootCmd.Flags().BoolP("kernel-version", "v", false, "print the kernel version")
	rootCmd.Flags().BoolP("machine", "m", false, "print the machine hardware name")
	rootCmd.Flags().BoolP("nodename", "n", false, "print the network node hostname")
	rootCmd.Flags().BoolP("operating-system", "o", false, "print the operating system")
	rootCmd.Flags().BoolP("processor", "p", false, "print the processor type (non-portable)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
