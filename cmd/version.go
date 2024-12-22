package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Date     string
	Revision string
	Version  string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
