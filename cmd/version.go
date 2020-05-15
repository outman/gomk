package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const gomarkVersion = "v0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "gomk version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(gomarkVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
