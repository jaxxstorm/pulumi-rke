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
    public sealed class ClusterCloudProviderVsphereCloudConfigGlobal
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Datacenter;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Datacenters;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Datastore;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool? InsecureFlag;
        /// <summary>
        /// Registry password (string)
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port used for SSH communication. Default `22` (string)
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// (int)
        /// </summary>
        public readonly int? SoapRoundtripCount;
        /// <summary>
        /// Registry user (string)
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? VmName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? VmUuid;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? WorkingDir;

        [OutputConstructor]
        private ClusterCloudProviderVsphereCloudConfigGlobal(
            string? datacenter,

            string? datacenters,

            string? datastore,

            bool? insecureFlag,

            string? password,

            string? port,

            int? soapRoundtripCount,

            string? user,

            string? vmName,

            string? vmUuid,

            string? workingDir)
        {
            Datacenter = datacenter;
            Datacenters = datacenters;
            Datastore = datastore;
            InsecureFlag = insecureFlag;
            Password = password;
            Port = port;
            SoapRoundtripCount = soapRoundtripCount;
            User = user;
            VmName = vmName;
            VmUuid = vmUuid;
            WorkingDir = workingDir;
        }
    }
}