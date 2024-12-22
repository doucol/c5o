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
	Short: "Print version, revision, and date of build",
	Long:  `Print version, revision, and date of build`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nRevision: %s\nDate: %s\n", Version, Revision, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
