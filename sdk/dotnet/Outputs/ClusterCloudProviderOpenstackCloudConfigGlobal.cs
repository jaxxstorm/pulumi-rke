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
    public sealed class ClusterCloudProviderOpenstackCloudConfigGlobal
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string AuthUrl;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? CaFile;
        /// <summary>
        /// Required if `domain_name` not provided. (string)
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// Required if `domain_id` not provided. (string)
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// Registry password (string)
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Region for S3 service (string)
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// Required if `tenant_name` not provided. (string)
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// Required if `tenant_id` not provided. (string)
        /// </summary>
        public readonly string? TenantName;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? TrustId;
        /// <summary>
        /// Required if `username` not provided. (string)
        /// </summary>
        public readonly string? UserId;
        /// <summary>
        /// Required if `user_id` not provided. (string)
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ClusterCloudProviderOpenstackCloudConfigGlobal(
            string authUrl,

            string? caFile,

            string? domainId,

            string? domainName,

            string password,

            string? region,

            string? tenantId,

            string? tenantName,

            string? trustId,

            string? userId,

            string? username)
        {
            AuthUrl = authUrl;
            CaFile = caFile;
            DomainId = domainId;
            DomainName = domainName;
            Password = password;
            Region = region;
            TenantId = tenantId;
            TenantName = tenantName;
            TrustId = trustId;
            UserId = userId;
            Username = username;
        }
    }
}
