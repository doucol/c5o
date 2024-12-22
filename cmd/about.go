package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Print version, revision, and date of build",
	Long:  `Print version, revision, and date of build`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nRevision: %s\nDate: %s\n", Version, Revision, Date)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
