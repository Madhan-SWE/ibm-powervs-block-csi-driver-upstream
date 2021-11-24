# PowerVS-CSI-Driver
CSI Driver for IBM® Power Systems™ Virtual Server

# Overview
The Power Virtual Systems Container Storage Interface (CSI) Driver provides a CSI interface used by Container Orchestrators to manage the lifecycle ofPower Virtual System volumes.


# CSI Specification Compatibility Matrix
| PowerVS CSI Driver \ CSI Version | v1.5.0 |
| ----------------------------- | -------|
| main branch | yes |

# Features
The following CSI gRPC calls are implemented:

- **Controller Service:** CreateVolume, DeleteVolume, ControllerPublishVolume,ControllerUnpublishVolume, ControllerGetCapabilities, ValidateVolumeCapabilities
- **Node Service:** NodeStageVolume, NodeUnstageVolume, NodePublishVolume, NodeUnpublishVolume, NodeGetCapabilities, NodeGetInfo
- **Identity Service:** GetPluginInfo, GetPluginCapabilities

# CreateVolume Parameters
There are several optional parameters that could be passed into ``` CreateVolumeRequest.parameters``` map, these parameters can be configured in StorageClass, see example:

| **Parameters** | **Values** | **Default** | **Description**|
| ----------------------------- | ----------------------------- | ----------- | ----------------------------- |
| "csi.storage.k8s.io/fstype" | xfs, ext2, ext3, ext4 | ext4 | File system type that will be formatted during volume creation. This parameter is case sensitive! |


## Driver Options
There are couple driver options that can be passed as arguments when starting driver container.

| Option argument             | value sample                                      | default                                             | Description         |
|-----------------------------|---------------------------------------------------|-----------------------------------------------------|---------------------|
| endpoint                    | tcp://127.0.0.1:10000/                            | unix:///var/lib/csi/sockets/pluginproxy/csi.sock    | added to all volumes, for checking if a given volume was already created so that ControllerPublish/CreateVolume is idempotent. |
| volume-attach-limit         | 1,2,3 ...                                         | -1                                                  | Value for the maximum number of volumes attachable per node. If specified, the limit applies to all nodes. If not specified, the value is approximated from the instance type.    |
| debug           | true                                              | false                                               | if true, driver will enable the debug log level|


# PowerVS CSI Driver on Kubernetes
Following sections are Kubernetes specific. If you are Kubernetes user, use followings for driver features, installation steps and examples.

## Kubernetes Version Compatibility Matrix
| PowerVS CSI Driver \ Kubernetes Version| v1.22 |
|----------------------------------------|-------|
| main branch                          | yes   |


## Container Images:
|PowerVS CSI Driver Version | quay.io image                                       | DockerHub Image                                                                 |
|---------------------------|--------------------------------------------------|-----------------------------------------------------------------------------|
|main                     | - - -  |  - - -   |


## Features
* **Static Provisioning** - create a new or migrating existing PowerVS volumes, then create persistence volume (PV) from the PowerVS volume and consume the PV from container using persistence volume claim (PVC).
* **Dynamic Provisioning** - uses persistence volume claim (PVC) to request the Kuberenetes to create the PowerVS volume on behalf of user and consumes the volume from inside container.
* **Mount Option** - mount options could be specified in persistence volume (PV) to define how the volume should be mounted.
* **[Block Volume](https://kubernetes-csi.github.io/docs/raw-block.html)** - consumes the PowerVS volume as a raw block device for latency sensitive application eg. MySql. The corresponding CSI feature (`CSIBlockVolume`) is GA since Kubernetes 1.18.
* **[Volume Resizing](https://kubernetes-csi.github.io/docs/volume-expansion.html)** - expand the volume size. The corresponding CSI feature (`ExpandCSIVolumes`) is beta since Kubernetes 1.16.

## Prerequisites
* If you are managing PowerVS volumes using static provisioning, get yourself familiar with [Power Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).
* Get yourself familiar with how to setup Kubernetes on IBM Cloud and have a working Kubernetes cluster:
  * Enable flag `--allow-privileged=true` for `kubelet` and `kube-apiserver`
  * Enable `kube-apiserver` feature gates `--feature-gates=CSINodeInfo=true,CSIDriverRegistry=true,CSIBlockVolume=true`
  * Enable `kubelet` feature gates `--feature-gates=CSINodeInfo=true,CSIDriverRegistry=true,CSIBlockVolume=true`


## Installation
#### Set up driver permission

* Using secret object - Generate IBMCLOUD_APIKEY from the UI, put that user's credentials in [secret manifest](../deploy/kubernetes/secret.yaml), then deploy the secret
```sh
curl https://raw.githubusercontent.com/ppc64le-cloud/powervs-csi-driver/blob/main/deploy/kubernetes/secret.yaml > secret.yaml edit the IBMCLOUD_API_KEY
# Edit the secret with user credentials
kubectl apply -f secret.yaml
```

#### Deploy driver
Please see the compatibility matrix above before you deploy the driver

To deploy the CSI driver:
```sh
kubectl apply -k "github.com/ppc64le-cloud/powervs-csi-driver/deploy/kubernetes/overlays/stable"
```

Verify driver is running:
```sh
kubectl get pods -n kube-system
```

#### Deploy driver with debug mode
To view driver debug logs, run the CSI driver with `-v=5` command line option

To enable powervs debug logs, run the CSI driver with `debug=true` command line option.

## Examples
Make sure you follow the [Prerequisites](README.md#Prerequisites) before the examples:
* [Dynamic Provisioning - need to add examples](../examples/kubernetes/dynamic-provisioning)
* [Block Volume - Need to add examples](../examples/kubernetes/block-volume)
* [Configure StorageClass - Need to add examples](../examples/kubernetes/storageclass)
* [Volume Resizing - Need to add examples](../examples/kubernetes/resizing)


## Development
Please go through [CSI Spec](https://github.com/container-storage-interface/spec/blob/master/spec.md) and [General CSI driver development guideline](https://kubernetes-csi.github.io/docs/developing.html) to get some basic understanding of CSI driver before you start.

### Requirements
* Golang 1.17.+
* [Ginkgo](https://github.com/onsi/ginkgo) in your PATH for integration testing and end-to-end testing
* Docker 20.10+ for releasing

### Dependency
Dependencies are managed through go module. To build the project, first turn on go mod using `export GO111MODULE=on`, then build the project using: `make`

### Testing
* To execute all unit tests, run: `make test` - pending
* To execute sanity test run: `make test-sanity` - pending
* To execute integration tests, run: `make test-integration` - pending
* To execute e2e tests, run: `make test-e2e-single-az` and `make test-e2e-multi-az` - pending




