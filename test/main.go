package main

// eyJhbGciOiJSUzI1NiIsImtpZCI6IjhuaXh5bVFReFNlRzd6YWR1YS1MdjFpcWxtc3FSTWxIRExVWU1NV1B2NmMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi12ajh4bCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjBmYmFhMTgwLWJlZjItNDkxYS05ODlmLWM3MjY0MzNmZDFmNCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTphZG1pbiJ9.JBZPsq4sVvtiaoXvP2Lg9Qr5mO_ItZ-rJIjHQxGqpW_cK2h-i-z8Fx0bYWB1h971cs7qanDma4fiTt8DsGhzmLtZ6e_LbUYboOCCNgbnuF-RGiqEImaqa1GLkDgO0edwlIGah729puTvfyftAnm3oh3TSEql7KloHfXYEUPoomQ4kgNthVQdJbCD_zllUSoOSWmYteaRCNrYXZfTsyVNFq0CLYoKZAPjCp16Ze90odSuRevhmixDIydizZNJxzBir8lGnrfbxHVxBXuWoJ9eLXNk6mlS0v7cA8gNsarkljCIRKJDaWBNN-LgtBld1NkBQ4-H9DyBHdt3Sgds37qOFA

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// GetClientByBearerToken 使用 BearerToken 获取 clientSet
func GetClientByBearerToken() *kubernetes.Clientset {
	// 指定 host 和 bearer token
	config := &rest.Config{
		Host:        "https://xxxxx:6443",
		BearerToken: "",
	}
	config.Insecure = true
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return client
}

// GetClientByConfig 使用 .kube/config 获取 clientSet
func GetClientByConfig() *kubernetes.Clientset {
	configPath := homedir.HomeDir() + "/.kube/config"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	config = metadata.ConfigFor(config)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	c := GetClientByBearerToken()
	fmt.Println(c.ServerVersion())
}
