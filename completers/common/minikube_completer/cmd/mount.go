package cmd

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mounts the specified directory into minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	mountCmd.Flags().String("9p-version", "9p2000.L", "Specify the 9p version that the mount should use")
	mountCmd.Flags().String("gid", "docker", "Default group id used for the mount")
	mountCmd.Flags().String("ip", "", "Specify the ip that the mount should be setup on")
	mountCmd.Flags().Bool("kill", false, "Kill the mount process spawned by minikube start")
	mountCmd.Flags().Uint("mode", 493, "File permissions used for the mount")
	mountCmd.Flags().Int("msize", 262144, "The number of bytes to use for 9p packet payload")
	mountCmd.Flags().StringSlice("options", []string{}, "Additional mount options, such as cache=fscache")
	mountCmd.Flags().String("type", "9p", "Specify the mount filesystem type (supported types: 9p)")
	mountCmd.Flags().String("uid", "docker", "Default user id used for the mount")
	rootCmd.AddCommand(mountCmd)

	carapace.Gen(mountCmd).FlagCompletion(carapace.ActionMap{
		"gid":  os.ActionGroups(),
		"mode": fs.ActionFileModesNumeric(),
		"type": carapace.ActionValues("9p"),
		"uid":  os.ActionUsers(),
	})

	carapace.Gen(mountCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionDirectories()
			case 1:
				path := filepath.Dir(c.CallbackValue)
				if path == "" {
					path = "/"
				}
				return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
					return carapace.ActionExecCommand("minikube", "ssh", "ls", `\-1`, `\-p`, path)(func(output []byte) carapace.Action {
						lines := strings.Split(string(output), "\n")
						return carapace.ActionValues(lines[:len(lines)-1]...)
					})
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
