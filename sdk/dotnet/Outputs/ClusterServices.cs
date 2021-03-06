// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rke.Outputs
{

    [OutputType]
    public sealed class ClusterServices
    {
        /// <summary>
        /// Docker image for etcd (string)
        /// </summary>
        public readonly Outputs.ClusterServicesEtcd? Etcd;
        /// <summary>
        /// Kube API options for RKE services (list maxitems:1)
        /// </summary>
        public readonly Outputs.ClusterServicesKubeApi? KubeApi;
        /// <summary>
        /// Kube Controller options for RKE services (list maxitems:1)
        /// </summary>
        public readonly Outputs.ClusterServicesKubeController? KubeController;
        /// <summary>
        /// Kubelet options for RKE services (list maxitems:1)
        /// </summary>
        public readonly Outputs.ClusterServicesKubelet? Kubelet;
        /// <summary>
        /// Kubeproxy options for RKE services (list maxitems:1)
        /// </summary>
        public readonly Outputs.ClusterServicesKubeproxy? Kubeproxy;
        /// <summary>
        /// Scheduler options for RKE services (list maxitems:1)
        /// </summary>
        public readonly Outputs.ClusterServicesScheduler? Scheduler;

        [OutputConstructor]
        private ClusterServices(
            Outputs.ClusterServicesEtcd? etcd,

            Outputs.ClusterServicesKubeApi? kubeApi,

            Outputs.ClusterServicesKubeController? kubeController,

            Outputs.ClusterServicesKubelet? kubelet,

            Outputs.ClusterServicesKubeproxy? kubeproxy,

            Outputs.ClusterServicesScheduler? scheduler)
        {
            Etcd = etcd;
            KubeApi = kubeApi;
            KubeController = kubeController;
            Kubelet = kubelet;
            Kubeproxy = kubeproxy;
            Scheduler = scheduler;
        }
    }
}
