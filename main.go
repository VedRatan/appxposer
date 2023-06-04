package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/vedratan/.kube/config", "location of your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error: %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()

		// if the app is deployed inside the cluster then it will automatically retrieve the config from the incluster config

		if err != nil {
			fmt.Printf("error %s, getting incluster config\n", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if(err != nil){
		fmt.Printf("error: %s creating a clientset\n", err.Error())
	}
	ch := make(chan struct{})
	informers := informers.NewSharedInformerFactory(clientset, 10*time.Minute)
	c := newController(clientset, informers.Apps().V1().Deployments())

	informers.Start(ch)
	c.run(ch)
	fmt.Println(clientset)

}