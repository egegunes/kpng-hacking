package main

import (
	"k8s.io/klog/v2"
	"sigs.k8s.io/kpng/client"
)

func main() {
	client.Run(handleEndpoints)
}

func handleEndpoints(items []*client.ServiceEndpoints) {
	for _, item := range items {
		klog.Info("item:", item.Service.Namespace, item.Service.Name, item.Service.Type)
		for _, endpoint := range item.Endpoints {
			klog.Info("endpoint:", endpoint.Hostname, endpoint.IPs, endpoint.Local)
		}
	}
}
