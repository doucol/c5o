package monitor

import (
	"context"
	"fmt"

	"github.com/doucol/c5o/internal"
	"github.com/spf13/cobra"
	k8sv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor a variety of Calico resources",
	Long:  `Monitor a variety of Calico resources`,
	RunE: func(cmd *cobra.Command, args []string) error {
		clientset := internal.ClientsetFromContext(cmd.Context())
		pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), k8sv1.ListOptions{})
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
	// Add all subcommands here
}
