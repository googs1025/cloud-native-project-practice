package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func InitKarmadaClient(karmadaConfigPath string) (dynamic.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", karmadaConfigPath)
	if err != nil {
		return nil, err
	}
	dClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return dClient, nil
}

func main() {
	c, err := InitKarmadaClient("./karmada/karmada-config")
	if err != nil {
		log.Fatal(err)
	}
	// 获取 karmada 自定义资源
	clusterGVR := schema.GroupVersionResource{
		Group:    "cluster.karmada.io",
		Version:  "v1alpha1",
		Resource: "clusters",
	}

	// list
	list, err := c.Resource(clusterGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list.Items {
		fmt.Println(item.GetName())
	}
}
