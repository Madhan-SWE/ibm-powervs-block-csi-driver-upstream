package cloud

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

// Metadata is info about the instance on which the driver is running
type Metadata struct {
	serviceInstanceID string
	nodeInstanceId    string
	region            string
}

// GetServiceInstanceId returns service instance id of the instance
func (m *Metadata) GetServiceInstanceId() string {
	return m.serviceInstanceID
}

// GetNodeInstanceId returns node instance id of the instance
func (m *Metadata) GetNodeInstanceId() string {
	return m.nodeInstanceId
}

// GetInstanceRegion returns region of the instance
func (m *Metadata) GetInstanceRegion() string {
	return m.region
}

type KubernetesAPIClient func() (kubernetes.Interface, error)

// Get default kubernetes API client
var DefaultKubernetesAPIClient = func() (kubernetes.Interface, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// Get testing kubernetes APIclient
var TestingKubernetesAPIClient = func() (kubernetes.Interface, error) {
	// Get configuration from config file
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		fmt.Println("ERROR building configuration:", err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// Get New Metadata Service
func NewMetadataService(k8sAPIClient KubernetesAPIClient) (MetadataService, error) {
	klog.Infof("retrieving instance data from kubernetes api")
	clientset, err := k8sAPIClient()
	if err != nil {
		klog.Warningf("error creating kubernetes api client: %v", err)
	} else {
		klog.Infof("kubernetes api is available")
		return KubernetesAPIInstanceInfo(clientset)
	}
	return nil, fmt.Errorf("error getting instance data from ec2 metadata or kubernetes api")
}

// Get instance info from kubernetes API
func KubernetesAPIInstanceInfo(clientset kubernetes.Interface) (*Metadata, error) {
	var nodeName = os.Getenv("CSI_NODE_NAME")
	if nodeName == "" {
		hostName, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeName = hostName
	}
	if nodeName == "" {
		return nil, fmt.Errorf("CSI_NODE_NAME env var not set")
	}
	// get node with k8s API
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting Node %v: %v", nodeName, err)
	}

	// Get node labels
	labels := node.GetLabels()
	keysList := []string{"topology.powervs.csi.ibm.com/serviceInstanceId", "topology.powervs.csi.ibm.com/nodeInstanceId", "topology.powervs.csi.ibm.com/region"}
	instanceInfo := Metadata{}
	for _, key := range keysList {
		if val, ok := labels[key]; ok {
			switch key {
			case "topology.powervs.csi.ibm.com/serviceInstanceId":
				instanceInfo.serviceInstanceID = val
				break
			case "topology.powervs.csi.ibm.com/nodeInstanceId":
				instanceInfo.nodeInstanceId = val
				break
			case "topology.powervs.csi.ibm.com/region":
				instanceInfo.region = val
				break
			}
		} else {
			return nil, fmt.Errorf("Error getting label %v for node Node %v", key, nodeName)
		}
	}

	return &instanceInfo, nil
}
