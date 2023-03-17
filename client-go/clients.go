package main

import (
	"context"
	"fmt"
	"log"

	_ "embed"

	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func clientSet() {
	cliset := newK8sConfig().initClient()

	pods, err := cliset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("namespace  %s, pod: %s\n", pod.Namespace, pod.Name)
	}
}

//go:embed tpls/deployment.yaml
var deployTpl string

func dynamicClient() {

	dynCli := newK8sConfig().initDynamicClient()

	deployGVR := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	deployObj := &unstructured.Unstructured{}

	if err := yaml.Unmarshal([]byte(deployTpl), deployObj); err != nil {
		log.Fatalln(err)
	}

	if _, err := dynCli.
		Resource(deployGVR).
		Namespace("default").
		Create(context.Background(), deployObj, metav1.CreateOptions{}); err != nil {
		log.Fatalln(err)
	}

	log.Println("created deployment")
}

func discoveryCli() {
	cli := newK8sConfig().initDiscoveryClient()

	res, _ := cli.ServerPreferredResources()

	for _, r := range res {
		fmt.Println(r.String())
	}

}
