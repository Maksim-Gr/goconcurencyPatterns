package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(homeDir, ".kube", "config")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("Error getting k8s config: %v\n", err)
		os.Exit(1)
	}
	clentset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error gettng kubernetes config %v\n", err)
		os.Exit(1)
	}
	namespace := "dev011"
	pods, err := ListPods(namespace, clentset)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %v\n", pod.Name)
	}

	var message string
	if message == "" {
		message = "Total pods in all namespaces"
	} else {
		message = fmt.Sprintf("Total pods in the namespace: %s", namespace)
	}
	fmt.Printf("%s %d\n", message, len(pods.Items))
}

func ListPods(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	fmt.Println("get available pods for namespace")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}
	return pods, nil
}
