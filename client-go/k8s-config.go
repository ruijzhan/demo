package main

import (
	"log"

	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sConfig struct {
	kubeconfigPath string
}

func newK8sConfig() *k8sConfig {
	path, err := homedir.Expand("~/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	return &k8sConfig{
		kubeconfigPath: path,
	}
}

func (c *k8sConfig) k8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", c.kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *k8sConfig) initClient() *kubernetes.Clientset {
	return kubernetes.NewForConfigOrDie(c.k8sRestConfig())
}

func (c *k8sConfig) initDynamicClient() dynamic.Interface {
	return dynamic.NewForConfigOrDie(c.k8sRestConfig())
}

func (c *k8sConfig) initDiscoveryClient() *discovery.DiscoveryClient {
	return discovery.NewDiscoveryClientForConfigOrDie(c.k8sRestConfig())
}
