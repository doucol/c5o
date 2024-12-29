package internal

import (
	"context"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

type CmdContextKeyType string

const CmdContextKey CmdContextKeyType = "CmdContextKey"

type CmdContext struct {
	KubeConfig  string
	KubeContext string
	*kubernetes.Clientset
}

func (c *CmdContext) ToContext() context.Context {
	return context.WithValue(context.Background(), CmdContextKey, c)
}

func CmdContextFromContext(ctx context.Context) *CmdContext {
	return ctx.Value(CmdContextKey).(*CmdContext)
}

func NewCmdContext(kubeConfig, kubeContext string) (*CmdContext, error) {
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfig}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: kubeContext}
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &CmdContext{
		KubeConfig:  kubeConfig,
		KubeContext: kubeContext,
		Clientset:   clientset,
	}, nil
}
