/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "Print version and exit",
	Long:  `Print just the version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
		fmt.Println(Revision)
		fmt.Println(Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
