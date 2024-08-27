package main

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// 获取配置
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	coreV1 := clientset.CoreV1()

	pod, err := coreV1.Pods("kube-system").Get(context.TODO(), "coredns-64897985d-d62tb", v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	} else {
		println(pod.Namespace, pod.Name, pod.UID)
		println("The pod label are: ")
		for k, v := range pod.Labels {
			println(k, v)
		}
	}

}
