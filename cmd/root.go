package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/doucol/c5o/internal"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "c5o",
	Short: "Project Calico utilities",
	Long:  fmt.Sprintf("c5o - %s\nA collection of Project Calico utilities which may or may not be helpful", Version),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}
		return nil
	},
}

var kubeConfig, kubeContext string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	var dflt string
	if home := homedir.HomeDir(); home != "" {
		dflt = filepath.Join(home, ".kube", "config")
	}
	kcev := os.Getenv("KUBECONFIG")
	if kcev != "" {
		dflt = kcev
	}
	rootCmd.PersistentFlags().StringVar(&kubeConfig, "kubeconfig", dflt, "absolute path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVar(&kubeContext, "kubecontext", "", "(optional) kubeconfig context to use")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() int {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		cmdContext, err := internal.NewCmdContext(kubeConfig, kubeContext)
		if err != nil {
			return err
		}
		cmd.SetContext(cmdContext.ToContext())
		return nil
	}
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}
	return 0
}
