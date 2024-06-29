package main

import (
	"fmt"
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
}
