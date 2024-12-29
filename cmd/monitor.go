package cmd

import (
	"context"
	"fmt"

	"github.com/doucol/c5o/internal"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor a variety of Calico resources",
	Long:  `Monitor a variety of Calico resources`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdCtx := internal.CmdContextFromContext(cmd.Context())
		clientset := cmdCtx.Clientset
		pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		for _, pod := range pods.Items {
			fmt.Printf("Pod name: %s\n", pod.Name)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
