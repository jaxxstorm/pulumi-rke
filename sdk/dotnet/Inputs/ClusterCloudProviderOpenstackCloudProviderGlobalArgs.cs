// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rke.Inputs
{

    public sealed class ClusterCloudProviderOpenstackCloudProviderGlobalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (string)
        /// </summary>
        [Input("authUrl", required: true)]
        public Input<string> AuthUrl { get; set; } = null!;

        /// <summary>
        /// (string)
        /// </summary>
        [Input("caFile")]
        public Input<string>? CaFile { get; set; }

        /// <summary>
        /// Required if `domain_name` not provided. (string)
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Required if `domain_id` not provided. (string)
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Registry password (string)
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Region for S3 service (string)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required if `tenant_name` not provided. (string)
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Required if `tenant_id` not provided. (string)
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        /// <summary>
        /// (string)
        /// </summary>
        [Input("trustId")]
        public Input<string>? TrustId { get; set; }

        /// <summary>
        /// Required if `username` not provided. (string)
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// Required if `user_id` not provided. (string)
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ClusterCloudProviderOpenstackCloudProviderGlobalArgs()
        {
        }
    }
}
