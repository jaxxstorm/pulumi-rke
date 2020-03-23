// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rke

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides RKE cluster resource. This can be used to create RKE clusters and retrieve their information.
//
// > This content is derived from https://github.com/rancher/terraform-provider-rke/blob/master/website/docs/r/cluster.html.markdown.
type Cluster struct {
	pulumi.CustomResourceState

	// Use services.etcd instead (list maxitems:1)
	Services_Etcd ClusterServices_EtcdPtrOutput `pulumi:"Services_Etcd"`
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi ClusterServices_KubeApiPtrOutput `pulumi:"Services_KubeApi"`
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController ClusterServices_KubeControllerPtrOutput `pulumi:"Services_KubeController"`
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy ClusterServices_KubeProxyPtrOutput `pulumi:"Services_KubeProxy"`
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler ClusterServices_KubeSchedulerPtrOutput `pulumi:"Services_KubeScheduler"`
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet ClusterServices_KubeletPtrOutput `pulumi:"Services_Kubelet"`
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout pulumi.IntOutput `pulumi:"addonJobTimeout"`
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons pulumi.StringPtrOutput `pulumi:"addons"`
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes pulumi.StringArrayOutput `pulumi:"addonsIncludes"`
	// (Computed) RKE k8s cluster api server url (string)
	ApiServerUrl pulumi.StringOutput `pulumi:"apiServerUrl"`
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication ClusterAuthenticationOutput `pulumi:"authentication"`
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization ClusterAuthorizationOutput `pulumi:"authorization"`
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost ClusterBastionHostPtrOutput `pulumi:"bastionHost"`
	// (Computed/Sensitive) RKE k8s cluster CA certificate (string)
	CaCrt pulumi.StringOutput `pulumi:"caCrt"`
	// Specify a certificate dir path (string)
	CertDir pulumi.StringPtrOutput `pulumi:"certDir"`
	// (Computed/Sensitive) RKE k8s cluster certificates (string)
	Certificates ClusterCertificateArrayOutput `pulumi:"certificates"`
	// (Computed/Sensitive) RKE k8s cluster client certificate (string)
	ClientCert pulumi.StringOutput `pulumi:"clientCert"`
	// (Computed/Sensitive) RKE k8s cluster client key (string)
	ClientKey pulumi.StringOutput `pulumi:"clientKey"`
	// Calico cloud provider (string)
	CloudProvider ClusterCloudProviderPtrOutput `pulumi:"cloudProvider"`
	// Cluster CIDR option for kube controller service (string)
	ClusterCidr pulumi.StringOutput `pulumi:"clusterCidr"`
	// Cluster DNS Server option for kubelet service (string)
	ClusterDnsServer pulumi.StringOutput `pulumi:"clusterDnsServer"`
	// Cluster Domain option for kubelet service. Default `cluster.local` (string)
	ClusterDomain pulumi.StringOutput `pulumi:"clusterDomain"`
	// RKE k8s cluster name used in the kube config (string)
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// (Computed) RKE k8s cluster control plane nodes (list)
	ControlPlaneHosts ClusterControlPlaneHostArrayOutput `pulumi:"controlPlaneHosts"`
	// Use custom certificates from a cert dir (string)
	CustomCerts pulumi.BoolPtrOutput `pulumi:"customCerts"`
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation pulumi.IntPtrOutput `pulumi:"delayOnCreation"`
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind pulumi.BoolPtrOutput `pulumi:"dind"`
	// DinD RKE cluster dns (string)
	DindDnsServer pulumi.StringPtrOutput `pulumi:"dindDnsServer"`
	// DinD RKE cluster storage driver (string)
	DindStorageDriver pulumi.StringPtrOutput `pulumi:"dindStorageDriver"`
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck pulumi.BoolPtrOutput `pulumi:"disablePortCheck"`
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns ClusterDnsOutput `pulumi:"dns"`
	// (Computed) RKE k8s cluster etcd nodes (list)
	EtcdHosts ClusterEtcdHostArrayOutput `pulumi:"etcdHosts"`
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion pulumi.BoolPtrOutput `pulumi:"ignoreDockerVersion"`
	// (Computed) RKE k8s cluster inactive nodes (list)
	InactiveHosts ClusterInactiveHostArrayOutput `pulumi:"inactiveHosts"`
	// Docker image for ingress (string)
	Ingress ClusterIngressOutput `pulumi:"ingress"`
	// (Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)
	InternalKubeConfigYaml pulumi.StringOutput `pulumi:"internalKubeConfigYaml"`
	// (Computed) RKE k8s cluster admin user (string)
	KubeAdminUser pulumi.StringOutput `pulumi:"kubeAdminUser"`
	// (Computed/Sensitive) RKE k8s cluster kube config yaml (string)
	KubeConfigYaml pulumi.StringOutput `pulumi:"kubeConfigYaml"`
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring ClusterMonitoringOutput `pulumi:"monitoring"`
	// (list maxitems:1)
	Network ClusterNetworkOutput `pulumi:"network"`
	// RKE k8s cluster nodes (list)
	Nodes ClusterNodeArrayOutput `pulumi:"nodes"`
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs pulumi.StringArrayOutput `pulumi:"nodesConfs"`
	// RKE k8s directory path (string)
	PrefixPath pulumi.StringOutput `pulumi:"prefixPath"`
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries ClusterPrivateRegistryArrayOutput `pulumi:"privateRegistries"`
	// Restore cluster. Default `false` (bool)
	Restore ClusterRestoreOutput `pulumi:"restore"`
	// (Computed/Sensitive) RKE k8s cluster config yaml (string)
	RkeClusterYaml pulumi.StringOutput `pulumi:"rkeClusterYaml"`
	// (Computed/Sensitive) RKE k8s cluster state (string)
	RkeState pulumi.StringOutput `pulumi:"rkeState"`
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates ClusterRotateCertificatesPtrOutput `pulumi:"rotateCertificates"`
	// (Computed) RKE k8s cluster running system images list (list)
	RunningSystemImages ClusterRunningSystemImagesOutput `pulumi:"runningSystemImages"`
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services ClusterServicesOutput `pulumi:"services"`
	// SSH Agent Auth enable (bool)
	SshAgentAuth pulumi.BoolPtrOutput `pulumi:"sshAgentAuth"`
	// SSH Certificate path (string)
	SshCertPath pulumi.StringOutput `pulumi:"sshCertPath"`
	// SSH Private Key path (string)
	SshKeyPath pulumi.StringOutput `pulumi:"sshKeyPath"`
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages ClusterSystemImagesPtrOutput `pulumi:"systemImages"`
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly pulumi.BoolPtrOutput `pulumi:"updateOnly"`
	// (Computed) RKE k8s cluster worker nodes (list)
	WorkerHosts ClusterWorkerHostArrayOutput `pulumi:"workerHosts"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("rke:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("rke:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd *ClusterServices_Etcd `pulumi:"Services_Etcd"`
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi *ClusterServices_KubeApi `pulumi:"Services_KubeApi"`
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController *ClusterServices_KubeController `pulumi:"Services_KubeController"`
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy *ClusterServices_KubeProxy `pulumi:"Services_KubeProxy"`
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler *ClusterServices_KubeScheduler `pulumi:"Services_KubeScheduler"`
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet *ClusterServices_Kubelet `pulumi:"Services_Kubelet"`
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout *int `pulumi:"addonJobTimeout"`
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons *string `pulumi:"addons"`
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes []string `pulumi:"addonsIncludes"`
	// (Computed) RKE k8s cluster api server url (string)
	ApiServerUrl *string `pulumi:"apiServerUrl"`
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication *ClusterAuthentication `pulumi:"authentication"`
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization *ClusterAuthorization `pulumi:"authorization"`
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost *ClusterBastionHost `pulumi:"bastionHost"`
	// (Computed/Sensitive) RKE k8s cluster CA certificate (string)
	CaCrt *string `pulumi:"caCrt"`
	// Specify a certificate dir path (string)
	CertDir *string `pulumi:"certDir"`
	// (Computed/Sensitive) RKE k8s cluster certificates (string)
	Certificates []ClusterCertificate `pulumi:"certificates"`
	// (Computed/Sensitive) RKE k8s cluster client certificate (string)
	ClientCert *string `pulumi:"clientCert"`
	// (Computed/Sensitive) RKE k8s cluster client key (string)
	ClientKey *string `pulumi:"clientKey"`
	// Calico cloud provider (string)
	CloudProvider *ClusterCloudProvider `pulumi:"cloudProvider"`
	// Cluster CIDR option for kube controller service (string)
	ClusterCidr *string `pulumi:"clusterCidr"`
	// Cluster DNS Server option for kubelet service (string)
	ClusterDnsServer *string `pulumi:"clusterDnsServer"`
	// Cluster Domain option for kubelet service. Default `cluster.local` (string)
	ClusterDomain *string `pulumi:"clusterDomain"`
	// RKE k8s cluster name used in the kube config (string)
	ClusterName *string `pulumi:"clusterName"`
	// (Computed) RKE k8s cluster control plane nodes (list)
	ControlPlaneHosts []ClusterControlPlaneHost `pulumi:"controlPlaneHosts"`
	// Use custom certificates from a cert dir (string)
	CustomCerts *bool `pulumi:"customCerts"`
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation *int `pulumi:"delayOnCreation"`
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind *bool `pulumi:"dind"`
	// DinD RKE cluster dns (string)
	DindDnsServer *string `pulumi:"dindDnsServer"`
	// DinD RKE cluster storage driver (string)
	DindStorageDriver *string `pulumi:"dindStorageDriver"`
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck *bool `pulumi:"disablePortCheck"`
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns *ClusterDns `pulumi:"dns"`
	// (Computed) RKE k8s cluster etcd nodes (list)
	EtcdHosts []ClusterEtcdHost `pulumi:"etcdHosts"`
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion *bool `pulumi:"ignoreDockerVersion"`
	// (Computed) RKE k8s cluster inactive nodes (list)
	InactiveHosts []ClusterInactiveHost `pulumi:"inactiveHosts"`
	// Docker image for ingress (string)
	Ingress *ClusterIngress `pulumi:"ingress"`
	// (Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)
	InternalKubeConfigYaml *string `pulumi:"internalKubeConfigYaml"`
	// (Computed) RKE k8s cluster admin user (string)
	KubeAdminUser *string `pulumi:"kubeAdminUser"`
	// (Computed/Sensitive) RKE k8s cluster kube config yaml (string)
	KubeConfigYaml *string `pulumi:"kubeConfigYaml"`
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring *ClusterMonitoring `pulumi:"monitoring"`
	// (list maxitems:1)
	Network *ClusterNetwork `pulumi:"network"`
	// RKE k8s cluster nodes (list)
	Nodes []ClusterNode `pulumi:"nodes"`
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs []string `pulumi:"nodesConfs"`
	// RKE k8s directory path (string)
	PrefixPath *string `pulumi:"prefixPath"`
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries []ClusterPrivateRegistry `pulumi:"privateRegistries"`
	// Restore cluster. Default `false` (bool)
	Restore *ClusterRestore `pulumi:"restore"`
	// (Computed/Sensitive) RKE k8s cluster config yaml (string)
	RkeClusterYaml *string `pulumi:"rkeClusterYaml"`
	// (Computed/Sensitive) RKE k8s cluster state (string)
	RkeState *string `pulumi:"rkeState"`
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates *ClusterRotateCertificates `pulumi:"rotateCertificates"`
	// (Computed) RKE k8s cluster running system images list (list)
	RunningSystemImages *ClusterRunningSystemImages `pulumi:"runningSystemImages"`
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services *ClusterServices `pulumi:"services"`
	// SSH Agent Auth enable (bool)
	SshAgentAuth *bool `pulumi:"sshAgentAuth"`
	// SSH Certificate path (string)
	SshCertPath *string `pulumi:"sshCertPath"`
	// SSH Private Key path (string)
	SshKeyPath *string `pulumi:"sshKeyPath"`
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages *ClusterSystemImages `pulumi:"systemImages"`
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly *bool `pulumi:"updateOnly"`
	// (Computed) RKE k8s cluster worker nodes (list)
	WorkerHosts []ClusterWorkerHost `pulumi:"workerHosts"`
}

type ClusterState struct {
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd ClusterServices_EtcdPtrInput
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi ClusterServices_KubeApiPtrInput
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController ClusterServices_KubeControllerPtrInput
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy ClusterServices_KubeProxyPtrInput
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler ClusterServices_KubeSchedulerPtrInput
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet ClusterServices_KubeletPtrInput
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout pulumi.IntPtrInput
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons pulumi.StringPtrInput
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes pulumi.StringArrayInput
	// (Computed) RKE k8s cluster api server url (string)
	ApiServerUrl pulumi.StringPtrInput
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication ClusterAuthenticationPtrInput
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization ClusterAuthorizationPtrInput
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost ClusterBastionHostPtrInput
	// (Computed/Sensitive) RKE k8s cluster CA certificate (string)
	CaCrt pulumi.StringPtrInput
	// Specify a certificate dir path (string)
	CertDir pulumi.StringPtrInput
	// (Computed/Sensitive) RKE k8s cluster certificates (string)
	Certificates ClusterCertificateArrayInput
	// (Computed/Sensitive) RKE k8s cluster client certificate (string)
	ClientCert pulumi.StringPtrInput
	// (Computed/Sensitive) RKE k8s cluster client key (string)
	ClientKey pulumi.StringPtrInput
	// Calico cloud provider (string)
	CloudProvider ClusterCloudProviderPtrInput
	// Cluster CIDR option for kube controller service (string)
	ClusterCidr pulumi.StringPtrInput
	// Cluster DNS Server option for kubelet service (string)
	ClusterDnsServer pulumi.StringPtrInput
	// Cluster Domain option for kubelet service. Default `cluster.local` (string)
	ClusterDomain pulumi.StringPtrInput
	// RKE k8s cluster name used in the kube config (string)
	ClusterName pulumi.StringPtrInput
	// (Computed) RKE k8s cluster control plane nodes (list)
	ControlPlaneHosts ClusterControlPlaneHostArrayInput
	// Use custom certificates from a cert dir (string)
	CustomCerts pulumi.BoolPtrInput
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation pulumi.IntPtrInput
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind pulumi.BoolPtrInput
	// DinD RKE cluster dns (string)
	DindDnsServer pulumi.StringPtrInput
	// DinD RKE cluster storage driver (string)
	DindStorageDriver pulumi.StringPtrInput
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck pulumi.BoolPtrInput
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns ClusterDnsPtrInput
	// (Computed) RKE k8s cluster etcd nodes (list)
	EtcdHosts ClusterEtcdHostArrayInput
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion pulumi.BoolPtrInput
	// (Computed) RKE k8s cluster inactive nodes (list)
	InactiveHosts ClusterInactiveHostArrayInput
	// Docker image for ingress (string)
	Ingress ClusterIngressPtrInput
	// (Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)
	InternalKubeConfigYaml pulumi.StringPtrInput
	// (Computed) RKE k8s cluster admin user (string)
	KubeAdminUser pulumi.StringPtrInput
	// (Computed/Sensitive) RKE k8s cluster kube config yaml (string)
	KubeConfigYaml pulumi.StringPtrInput
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion pulumi.StringPtrInput
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring ClusterMonitoringPtrInput
	// (list maxitems:1)
	Network ClusterNetworkPtrInput
	// RKE k8s cluster nodes (list)
	Nodes ClusterNodeArrayInput
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs pulumi.StringArrayInput
	// RKE k8s directory path (string)
	PrefixPath pulumi.StringPtrInput
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries ClusterPrivateRegistryArrayInput
	// Restore cluster. Default `false` (bool)
	Restore ClusterRestorePtrInput
	// (Computed/Sensitive) RKE k8s cluster config yaml (string)
	RkeClusterYaml pulumi.StringPtrInput
	// (Computed/Sensitive) RKE k8s cluster state (string)
	RkeState pulumi.StringPtrInput
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates ClusterRotateCertificatesPtrInput
	// (Computed) RKE k8s cluster running system images list (list)
	RunningSystemImages ClusterRunningSystemImagesPtrInput
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services ClusterServicesPtrInput
	// SSH Agent Auth enable (bool)
	SshAgentAuth pulumi.BoolPtrInput
	// SSH Certificate path (string)
	SshCertPath pulumi.StringPtrInput
	// SSH Private Key path (string)
	SshKeyPath pulumi.StringPtrInput
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages ClusterSystemImagesPtrInput
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly pulumi.BoolPtrInput
	// (Computed) RKE k8s cluster worker nodes (list)
	WorkerHosts ClusterWorkerHostArrayInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd *ClusterServices_Etcd `pulumi:"Services_Etcd"`
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi *ClusterServices_KubeApi `pulumi:"Services_KubeApi"`
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController *ClusterServices_KubeController `pulumi:"Services_KubeController"`
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy *ClusterServices_KubeProxy `pulumi:"Services_KubeProxy"`
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler *ClusterServices_KubeScheduler `pulumi:"Services_KubeScheduler"`
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet *ClusterServices_Kubelet `pulumi:"Services_Kubelet"`
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout *int `pulumi:"addonJobTimeout"`
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons *string `pulumi:"addons"`
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes []string `pulumi:"addonsIncludes"`
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication *ClusterAuthentication `pulumi:"authentication"`
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization *ClusterAuthorization `pulumi:"authorization"`
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost *ClusterBastionHost `pulumi:"bastionHost"`
	// Specify a certificate dir path (string)
	CertDir *string `pulumi:"certDir"`
	// Calico cloud provider (string)
	CloudProvider *ClusterCloudProvider `pulumi:"cloudProvider"`
	// RKE k8s cluster name used in the kube config (string)
	ClusterName *string `pulumi:"clusterName"`
	// Use custom certificates from a cert dir (string)
	CustomCerts *bool `pulumi:"customCerts"`
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation *int `pulumi:"delayOnCreation"`
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind *bool `pulumi:"dind"`
	// DinD RKE cluster dns (string)
	DindDnsServer *string `pulumi:"dindDnsServer"`
	// DinD RKE cluster storage driver (string)
	DindStorageDriver *string `pulumi:"dindStorageDriver"`
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck *bool `pulumi:"disablePortCheck"`
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns *ClusterDns `pulumi:"dns"`
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion *bool `pulumi:"ignoreDockerVersion"`
	// Docker image for ingress (string)
	Ingress *ClusterIngress `pulumi:"ingress"`
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring *ClusterMonitoring `pulumi:"monitoring"`
	// (list maxitems:1)
	Network *ClusterNetwork `pulumi:"network"`
	// RKE k8s cluster nodes (list)
	Nodes []ClusterNode `pulumi:"nodes"`
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs []string `pulumi:"nodesConfs"`
	// RKE k8s directory path (string)
	PrefixPath *string `pulumi:"prefixPath"`
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries []ClusterPrivateRegistry `pulumi:"privateRegistries"`
	// Restore cluster. Default `false` (bool)
	Restore *ClusterRestore `pulumi:"restore"`
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates *ClusterRotateCertificates `pulumi:"rotateCertificates"`
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services *ClusterServices `pulumi:"services"`
	// SSH Agent Auth enable (bool)
	SshAgentAuth *bool `pulumi:"sshAgentAuth"`
	// SSH Certificate path (string)
	SshCertPath *string `pulumi:"sshCertPath"`
	// SSH Private Key path (string)
	SshKeyPath *string `pulumi:"sshKeyPath"`
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages *ClusterSystemImages `pulumi:"systemImages"`
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly *bool `pulumi:"updateOnly"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd ClusterServices_EtcdPtrInput
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi ClusterServices_KubeApiPtrInput
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController ClusterServices_KubeControllerPtrInput
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy ClusterServices_KubeProxyPtrInput
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler ClusterServices_KubeSchedulerPtrInput
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet ClusterServices_KubeletPtrInput
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout pulumi.IntPtrInput
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons pulumi.StringPtrInput
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes pulumi.StringArrayInput
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication ClusterAuthenticationPtrInput
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization ClusterAuthorizationPtrInput
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost ClusterBastionHostPtrInput
	// Specify a certificate dir path (string)
	CertDir pulumi.StringPtrInput
	// Calico cloud provider (string)
	CloudProvider ClusterCloudProviderPtrInput
	// RKE k8s cluster name used in the kube config (string)
	ClusterName pulumi.StringPtrInput
	// Use custom certificates from a cert dir (string)
	CustomCerts pulumi.BoolPtrInput
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation pulumi.IntPtrInput
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind pulumi.BoolPtrInput
	// DinD RKE cluster dns (string)
	DindDnsServer pulumi.StringPtrInput
	// DinD RKE cluster storage driver (string)
	DindStorageDriver pulumi.StringPtrInput
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck pulumi.BoolPtrInput
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns ClusterDnsPtrInput
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion pulumi.BoolPtrInput
	// Docker image for ingress (string)
	Ingress ClusterIngressPtrInput
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion pulumi.StringPtrInput
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring ClusterMonitoringPtrInput
	// (list maxitems:1)
	Network ClusterNetworkPtrInput
	// RKE k8s cluster nodes (list)
	Nodes ClusterNodeArrayInput
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs pulumi.StringArrayInput
	// RKE k8s directory path (string)
	PrefixPath pulumi.StringPtrInput
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries ClusterPrivateRegistryArrayInput
	// Restore cluster. Default `false` (bool)
	Restore ClusterRestorePtrInput
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates ClusterRotateCertificatesPtrInput
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services ClusterServicesPtrInput
	// SSH Agent Auth enable (bool)
	SshAgentAuth pulumi.BoolPtrInput
	// SSH Certificate path (string)
	SshCertPath pulumi.StringPtrInput
	// SSH Private Key path (string)
	SshKeyPath pulumi.StringPtrInput
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages ClusterSystemImagesPtrInput
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly pulumi.BoolPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

