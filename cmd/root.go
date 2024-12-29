package cmd

import (
	"fmt"
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

var KubeConfig, KubeContext string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// var cfgFile string
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.c5o.yaml)")

	var def string
	if home := homedir.HomeDir(); home != "" {
		def = filepath.Join(home, ".kube", "config")
	}
	rootCmd.PersistentFlags().StringVar(&KubeConfig, "kubeconfig", def, "absolute path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVar(&KubeContext, "kubecontext", "", "(optional) kubeconfig context to use")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() int {
	fmt.Printf("config: %s, ctx: %s", KubeConfig, KubeContext)
	cmdContext, err := internal.NewCmdContext(KubeConfig, KubeContext)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return 1
	}
	err = rootCmd.ExecuteContext(cmdContext.ToContext())
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return 1
	}
	return 0
}
