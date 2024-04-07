package main

import (
	"context"
	"fmt"

	rolloutsclient "github.com/argoproj/argo-rollouts/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// 引用 rollouts client 包

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "./kube/config")
	if err != nil {
		return
	}
	client, err := rolloutsclient.NewForConfig(config)
	rolloutList, err := client.ArgoprojV1alpha1().
		Rollouts("default").List(context.Background(), v1.ListOptions{})

	for _, v := range rolloutList.Items {
		fmt.Println(v.Name)
	}
}
