// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from ".././types/output";

export interface ClusterAuthentication {
    sans: string[];
    strategy?: string;
    webhook: outputs.ClusterAuthenticationWebhook;
}

export interface ClusterAuthenticationWebhook {
    cacheTimeout?: string;
    configFile?: string;
}

export interface ClusterAuthorization {
    mode?: string;
    options?: {[key: string]: any};
}

export interface ClusterBastionHost {
    address: string;
    port?: string;
    sshAgentAuth?: boolean;
    sshCert?: string;
    sshCertPath: string;
    sshKey?: string;
    sshKeyPath?: string;
    user: string;
}

export interface ClusterCertificate {
    certificate: string;
    commonName: string;
    config: string;
    configEnvName: string;
    configPath: string;
    envName: string;
    id: string;
    key: string;
    keyEnvName: string;
    keyPath: string;
    name: string;
    ouName: string;
    path: string;
}

export interface ClusterCloudProvider {
    awsCloudConfig?: outputs.ClusterCloudProviderAwsCloudConfig;
    awsCloudProvider?: outputs.ClusterCloudProviderAwsCloudProvider;
    azureCloudConfig?: outputs.ClusterCloudProviderAzureCloudConfig;
    azureCloudProvider?: outputs.ClusterCloudProviderAzureCloudProvider;
    customCloudConfig?: string;
    customCloudProvider?: string;
    name: string;
    openstackCloudConfig?: outputs.ClusterCloudProviderOpenstackCloudConfig;
    openstackCloudProvider?: outputs.ClusterCloudProviderOpenstackCloudProvider;
    vsphereCloudConfig?: outputs.ClusterCloudProviderVsphereCloudConfig;
    vsphereCloudProvider?: outputs.ClusterCloudProviderVsphereCloudProvider;
}

export interface ClusterCloudProviderAwsCloudConfig {
    global?: outputs.ClusterCloudProviderAwsCloudConfigGlobal;
    serviceOverrides?: outputs.ClusterCloudProviderAwsCloudConfigServiceOverride[];
}

export interface ClusterCloudProviderAwsCloudConfigGlobal {
    disableSecurityGroupIngress?: boolean;
    disableStrictZoneCheck?: boolean;
    elbSecurityGroup?: string;
    kubernetesClusterId?: string;
    kubernetesClusterTag?: string;
    roleArn?: string;
    routeTableId?: string;
    subnetId?: string;
    vpc?: string;
    zone?: string;
}

export interface ClusterCloudProviderAwsCloudConfigServiceOverride {
    key?: string;
    region?: string;
    service: string;
    signingMethod: string;
    signingName?: string;
    signingRegion?: string;
    url?: string;
}

export interface ClusterCloudProviderAwsCloudProvider {
    global?: outputs.ClusterCloudProviderAwsCloudProviderGlobal;
    serviceOverrides?: outputs.ClusterCloudProviderAwsCloudProviderServiceOverride[];
}

export interface ClusterCloudProviderAwsCloudProviderGlobal {
    disableSecurityGroupIngress?: boolean;
    disableStrictZoneCheck?: boolean;
    elbSecurityGroup?: string;
    kubernetesClusterId?: string;
    kubernetesClusterTag?: string;
    roleArn?: string;
    routeTableId?: string;
    subnetId?: string;
    vpc?: string;
    zone?: string;
}

export interface ClusterCloudProviderAwsCloudProviderServiceOverride {
    key?: string;
    region?: string;
    service: string;
    signingMethod: string;
    signingName?: string;
    signingRegion?: string;
    url?: string;
}

export interface ClusterCloudProviderAzureCloudConfig {
    aadClientCertPassword?: string;
    aadClientCertPath?: string;
    aadClientId: string;
    aadClientSecret: string;
    cloud?: string;
    cloudProviderBackoff?: boolean;
    cloudProviderBackoffDuration?: number;
    cloudProviderBackoffExponent?: number;
    cloudProviderBackoffJitter?: number;
    cloudProviderBackoffRetries?: number;
    cloudProviderRateLimit?: boolean;
    cloudProviderRateLimitBucket: number;
    cloudProviderRateLimitQps?: number;
    loadBalancerSku?: string;
    location?: string;
    maximumLoadBalancerRuleCount?: number;
    primaryAvailabilitySetName?: string;
    primaryScaleSetName?: string;
    resourceGroup?: string;
    routeTableName?: string;
    securityGroupName?: string;
    subnetName?: string;
    subscriptionId: string;
    tenantId: string;
    useInstanceMetadata?: boolean;
    useManagedIdentityExtension?: boolean;
    vmType?: string;
    vnetName?: string;
    vnetResourceGroup?: string;
}

export interface ClusterCloudProviderAzureCloudProvider {
    aadClientCertPassword?: string;
    aadClientCertPath?: string;
    aadClientId: string;
    aadClientSecret: string;
    cloud?: string;
    cloudProviderBackoff?: boolean;
    cloudProviderBackoffDuration?: number;
    cloudProviderBackoffExponent?: number;
    cloudProviderBackoffJitter?: number;
    cloudProviderBackoffRetries?: number;
    cloudProviderRateLimit?: boolean;
    cloudProviderRateLimitBucket: number;
    cloudProviderRateLimitQps?: number;
    loadBalancerSku?: string;
    location?: string;
    maximumLoadBalancerRuleCount?: number;
    primaryAvailabilitySetName?: string;
    primaryScaleSetName?: string;
    resourceGroup?: string;
    routeTableName?: string;
    securityGroupName?: string;
    subnetName?: string;
    subscriptionId: string;
    tenantId: string;
    useInstanceMetadata?: boolean;
    useManagedIdentityExtension?: boolean;
    vmType?: string;
    vnetName?: string;
    vnetResourceGroup?: string;
}

export interface ClusterCloudProviderOpenstackCloudConfig {
    blockStorage?: outputs.ClusterCloudProviderOpenstackCloudConfigBlockStorage;
    global: outputs.ClusterCloudProviderOpenstackCloudConfigGlobal;
    loadBalancer?: outputs.ClusterCloudProviderOpenstackCloudConfigLoadBalancer;
    metadata?: outputs.ClusterCloudProviderOpenstackCloudConfigMetadata;
    route?: outputs.ClusterCloudProviderOpenstackCloudConfigRoute;
}

export interface ClusterCloudProviderOpenstackCloudConfigBlockStorage {
    bsVersion?: string;
    ignoreVolumeAz?: boolean;
    trustDevicePath?: boolean;
}

export interface ClusterCloudProviderOpenstackCloudConfigGlobal {
    authUrl: string;
    caFile?: string;
    domainId?: string;
    domainName?: string;
    password: string;
    region?: string;
    tenantId?: string;
    tenantName?: string;
    trustId?: string;
    userId?: string;
    username?: string;
}

export interface ClusterCloudProviderOpenstackCloudConfigLoadBalancer {
    createMonitor?: boolean;
    floatingNetworkId?: string;
    lbMethod?: string;
    lbProvider?: string;
    lbVersion?: string;
    manageSecurityGroups?: boolean;
    monitorDelay?: string;
    monitorMaxRetries?: number;
    monitorTimeout?: string;
    subnetId?: string;
    useOctavia?: boolean;
}

export interface ClusterCloudProviderOpenstackCloudConfigMetadata {
    requestTimeout?: number;
    searchOrder?: string;
}

export interface ClusterCloudProviderOpenstackCloudConfigRoute {
    routerId?: string;
}

export interface ClusterCloudProviderOpenstackCloudProvider {
    blockStorage?: outputs.ClusterCloudProviderOpenstackCloudProviderBlockStorage;
    global: outputs.ClusterCloudProviderOpenstackCloudProviderGlobal;
    loadBalancer?: outputs.ClusterCloudProviderOpenstackCloudProviderLoadBalancer;
    metadata?: outputs.ClusterCloudProviderOpenstackCloudProviderMetadata;
    route?: outputs.ClusterCloudProviderOpenstackCloudProviderRoute;
}

export interface ClusterCloudProviderOpenstackCloudProviderBlockStorage {
    bsVersion?: string;
    ignoreVolumeAz?: boolean;
    trustDevicePath?: boolean;
}

export interface ClusterCloudProviderOpenstackCloudProviderGlobal {
    authUrl: string;
    caFile?: string;
    domainId?: string;
    domainName?: string;
    password: string;
    region?: string;
    tenantId?: string;
    tenantName?: string;
    trustId?: string;
    userId?: string;
    username?: string;
}

export interface ClusterCloudProviderOpenstackCloudProviderLoadBalancer {
    createMonitor?: boolean;
    floatingNetworkId?: string;
    lbMethod?: string;
    lbProvider?: string;
    lbVersion?: string;
    manageSecurityGroups?: boolean;
    monitorDelay?: string;
    monitorMaxRetries?: number;
    monitorTimeout?: string;
    subnetId?: string;
    useOctavia?: boolean;
}

export interface ClusterCloudProviderOpenstackCloudProviderMetadata {
    requestTimeout?: number;
    searchOrder?: string;
}

export interface ClusterCloudProviderOpenstackCloudProviderRoute {
    routerId?: string;
}

export interface ClusterCloudProviderVsphereCloudConfig {
    disk: outputs.ClusterCloudProviderVsphereCloudConfigDisk;
    global: outputs.ClusterCloudProviderVsphereCloudConfigGlobal;
    network: outputs.ClusterCloudProviderVsphereCloudConfigNetwork;
    virtualCenters: outputs.ClusterCloudProviderVsphereCloudConfigVirtualCenter[];
    workspace: outputs.ClusterCloudProviderVsphereCloudConfigWorkspace;
}

export interface ClusterCloudProviderVsphereCloudConfigDisk {
    scsiControllerType?: string;
}

export interface ClusterCloudProviderVsphereCloudConfigGlobal {
    datacenter?: string;
    datacenters: string;
    datastore?: string;
    insecureFlag?: boolean;
    password?: string;
    port?: string;
    soapRoundtripCount?: number;
    user?: string;
    vmName?: string;
    vmUuid?: string;
    workingDir?: string;
}

export interface ClusterCloudProviderVsphereCloudConfigNetwork {
    publicNetwork?: string;
}

export interface ClusterCloudProviderVsphereCloudConfigVirtualCenter {
    datacenters: string;
    name: string;
    password: string;
    port?: string;
    soapRoundtripCount?: number;
    user: string;
}

export interface ClusterCloudProviderVsphereCloudConfigWorkspace {
    datacenter: string;
    defaultDatastore?: string;
    folder?: string;
    resourcepoolPath?: string;
    server: string;
}

export interface ClusterCloudProviderVsphereCloudProvider {
    disk: outputs.ClusterCloudProviderVsphereCloudProviderDisk;
    global: outputs.ClusterCloudProviderVsphereCloudProviderGlobal;
    network: outputs.ClusterCloudProviderVsphereCloudProviderNetwork;
    virtualCenters: outputs.ClusterCloudProviderVsphereCloudProviderVirtualCenter[];
    workspace: outputs.ClusterCloudProviderVsphereCloudProviderWorkspace;
}

export interface ClusterCloudProviderVsphereCloudProviderDisk {
    scsiControllerType?: string;
}

export interface ClusterCloudProviderVsphereCloudProviderGlobal {
    datacenter?: string;
    datacenters: string;
    datastore?: string;
    insecureFlag?: boolean;
    password?: string;
    port?: string;
    soapRoundtripCount?: number;
    user?: string;
    vmName?: string;
    vmUuid?: string;
    workingDir?: string;
}

export interface ClusterCloudProviderVsphereCloudProviderNetwork {
    publicNetwork?: string;
}

export interface ClusterCloudProviderVsphereCloudProviderVirtualCenter {
    datacenters: string;
    name: string;
    password: string;
    port?: string;
    soapRoundtripCount?: number;
    user: string;
}

export interface ClusterCloudProviderVsphereCloudProviderWorkspace {
    datacenter: string;
    defaultDatastore?: string;
    folder?: string;
    resourcepoolPath?: string;
    server: string;
}

export interface ClusterControlPlaneHost {
    address: string;
    nodeName: string;
}

export interface ClusterDns {
    nodeSelector?: {[key: string]: any};
    provider?: string;
    reverseCidrs?: string[];
    upstreamNameservers?: string[];
}

export interface ClusterEtcdHost {
    address: string;
    nodeName: string;
}

export interface ClusterInactiveHost {
    address: string;
    nodeName: string;
}

export interface ClusterIngress {
    dnsPolicy?: string;
    extraArgs?: {[key: string]: any};
    nodeSelector?: {[key: string]: any};
    options?: {[key: string]: any};
    provider?: string;
}

export interface ClusterMonitoring {
    nodeSelector?: {[key: string]: any};
    options?: {[key: string]: any};
    provider: string;
}

export interface ClusterNetwork {
    calicoNetworkProvider?: outputs.ClusterNetworkCalicoNetworkProvider;
    canalNetworkProvider?: outputs.ClusterNetworkCanalNetworkProvider;
    flannelNetworkProvider?: outputs.ClusterNetworkFlannelNetworkProvider;
    mtu?: number;
    options: {[key: string]: any};
    plugin?: string;
    weaveNetworkProvider?: outputs.ClusterNetworkWeaveNetworkProvider;
}

export interface ClusterNetworkCalicoNetworkProvider {
    cloudProvider: string;
}

export interface ClusterNetworkCanalNetworkProvider {
    iface: string;
}

export interface ClusterNetworkFlannelNetworkProvider {
    iface: string;
}

export interface ClusterNetworkWeaveNetworkProvider {
    password: string;
}

export interface ClusterNode {
    address: string;
    dockerSocket: string;
    hostnameOverride: string;
    internalAddress: string;
    labels?: {[key: string]: any};
    nodeName?: string;
    port?: string;
    roles: string[];
    rolesDeprecated?: string;
    sshAgentAuth: boolean;
    sshCert?: string;
    sshCertPath?: string;
    sshKey: string;
    sshKeyPath: string;
    taints?: outputs.ClusterNodeTaint[];
    user: string;
}

export interface ClusterNodeTaint {
    effect?: string;
    key: string;
    value: string;
}

export interface ClusterPrivateRegistry {
    isDefault?: boolean;
    password?: string;
    url: string;
    user?: string;
}

export interface ClusterRestore {
    restore?: boolean;
    snapshotName?: string;
}

export interface ClusterRotateCertificates {
    caCertificates?: boolean;
    services?: string[];
}

export interface ClusterRunningSystemImages {
    alpine?: string;
    calicoCni?: string;
    calicoControllers?: string;
    calicoCtl?: string;
    calicoFlexVol?: string;
    calicoNode?: string;
    canalCni?: string;
    canalFlannel?: string;
    canalFlexVol?: string;
    canalNode?: string;
    certDownloader?: string;
    coredns?: string;
    corednsAutoscaler?: string;
    dnsmasq?: string;
    etcd?: string;
    flannel?: string;
    flannelCni?: string;
    ingress?: string;
    ingressBackend?: string;
    kubeDns?: string;
    kubeDnsAutoscaler?: string;
    kubeDnsSidecar?: string;
    kubernetes?: string;
    kubernetesServicesSidecar?: string;
    metricsServer?: string;
    nginxProxy?: string;
    nodelocal?: string;
    podInfraContainer?: string;
    weaveCni?: string;
    weaveNode?: string;
    windowsPodInfraContainer?: string;
}

export interface ClusterServices {
    etcd: outputs.ClusterServicesEtcd;
    kubeApi: outputs.ClusterServicesKubeApi;
    kubeController: outputs.ClusterServicesKubeController;
    kubelet: outputs.ClusterServicesKubelet;
    kubeproxy: outputs.ClusterServicesKubeproxy;
    scheduler: outputs.ClusterServicesScheduler;
}

export interface ClusterServicesEtcd {
    backupConfig: outputs.ClusterServicesEtcdBackupConfig;
    caCert: string;
    cert: string;
    creation: string;
    externalUrls: string[];
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    gid?: number;
    image: string;
    key: string;
    path: string;
    retention: string;
    snapshot?: boolean;
    uid?: number;
}

export interface ClusterServicesEtcdBackupConfig {
    enabled?: boolean;
    intervalHours?: number;
    retention?: number;
    s3BackupConfig?: outputs.ClusterServicesEtcdBackupConfigS3BackupConfig;
    safeTimestamp?: boolean;
}

export interface ClusterServicesEtcdBackupConfigS3BackupConfig {
    accessKey?: string;
    bucketName?: string;
    customCa?: string;
    endpoint?: string;
    folder?: string;
    region?: string;
    secretKey?: string;
}

export interface ClusterServicesEtcdDeprecated {
    backupConfig: outputs.ClusterServicesEtcdDeprecatedBackupConfig;
    caCert: string;
    cert: string;
    creation: string;
    externalUrls: string[];
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    gid?: number;
    image: string;
    key: string;
    path: string;
    retention: string;
    snapshot?: boolean;
    uid?: number;
}

export interface ClusterServicesEtcdDeprecatedBackupConfig {
    enabled?: boolean;
    intervalHours?: number;
    retention?: number;
    s3BackupConfig?: outputs.ClusterServicesEtcdDeprecatedBackupConfigS3BackupConfig;
    safeTimestamp?: boolean;
}

export interface ClusterServicesEtcdDeprecatedBackupConfigS3BackupConfig {
    accessKey?: string;
    bucketName?: string;
    customCa?: string;
    endpoint?: string;
    folder?: string;
    region?: string;
    secretKey?: string;
}

export interface ClusterServicesKubeApi {
    alwaysPullImages?: boolean;
    auditLog?: outputs.ClusterServicesKubeApiAuditLog;
    eventRateLimit?: outputs.ClusterServicesKubeApiEventRateLimit;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
    podSecurityPolicy?: boolean;
    secretsEncryptionConfig?: outputs.ClusterServicesKubeApiSecretsEncryptionConfig;
    serviceClusterIpRange: string;
    serviceNodePortRange: string;
}

export interface ClusterServicesKubeApiAuditLog {
    configuration: outputs.ClusterServicesKubeApiAuditLogConfiguration;
    enabled?: boolean;
}

export interface ClusterServicesKubeApiAuditLogConfiguration {
    format?: string;
    maxAge?: number;
    maxBackup?: number;
    maxSize?: number;
    path?: string;
    policy: string;
}

export interface ClusterServicesKubeApiDeprecated {
    alwaysPullImages?: boolean;
    auditLog?: outputs.ClusterServicesKubeApiDeprecatedAuditLog;
    eventRateLimit?: outputs.ClusterServicesKubeApiDeprecatedEventRateLimit;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
    podSecurityPolicy?: boolean;
    secretsEncryptionConfig?: outputs.ClusterServicesKubeApiDeprecatedSecretsEncryptionConfig;
    serviceClusterIpRange: string;
    serviceNodePortRange: string;
}

export interface ClusterServicesKubeApiDeprecatedAuditLog {
    configuration: outputs.ClusterServicesKubeApiDeprecatedAuditLogConfiguration;
    enabled?: boolean;
}

export interface ClusterServicesKubeApiDeprecatedAuditLogConfiguration {
    format?: string;
    maxAge?: number;
    maxBackup?: number;
    maxSize?: number;
    path?: string;
    policy: string;
}

export interface ClusterServicesKubeApiDeprecatedEventRateLimit {
    enabled?: boolean;
}

export interface ClusterServicesKubeApiDeprecatedSecretsEncryptionConfig {
    enabled?: boolean;
}

export interface ClusterServicesKubeApiEventRateLimit {
    enabled?: boolean;
}

export interface ClusterServicesKubeApiSecretsEncryptionConfig {
    enabled?: boolean;
}

export interface ClusterServicesKubeController {
    clusterCidr: string;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
    serviceClusterIpRange: string;
}

export interface ClusterServicesKubeControllerDeprecated {
    clusterCidr: string;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
    serviceClusterIpRange: string;
}

export interface ClusterServicesKubeProxyDeprecated {
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
}

export interface ClusterServicesKubeSchedulerDeprecated {
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
}

export interface ClusterServicesKubelet {
    clusterDnsServer: string;
    clusterDomain?: string;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    failSwapOn: boolean;
    generateServingCertificate?: boolean;
    image: string;
    infraContainerImage: string;
}

export interface ClusterServicesKubeletDeprecated {
    clusterDnsServer: string;
    clusterDomain?: string;
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    failSwapOn: boolean;
    generateServingCertificate?: boolean;
    image: string;
    infraContainerImage: string;
}

export interface ClusterServicesKubeproxy {
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
}

export interface ClusterServicesScheduler {
    extraArgs: {[key: string]: any};
    extraBinds: string[];
    extraEnvs: string[];
    image: string;
}

export interface ClusterSystemImages {
    alpine?: string;
    calicoCni?: string;
    calicoControllers?: string;
    calicoCtl?: string;
    calicoFlexVol?: string;
    calicoNode?: string;
    canalCni?: string;
    canalFlannel?: string;
    canalFlexVol?: string;
    canalNode?: string;
    certDownloader?: string;
    coredns?: string;
    corednsAutoscaler?: string;
    dnsmasq?: string;
    etcd?: string;
    flannel?: string;
    flannelCni?: string;
    ingress?: string;
    ingressBackend?: string;
    kubeDns?: string;
    kubeDnsAutoscaler?: string;
    kubeDnsSidecar?: string;
    kubernetes?: string;
    kubernetesServicesSidecar?: string;
    metricsServer?: string;
    nginxProxy?: string;
    nodelocal?: string;
    podInfraContainer?: string;
    weaveCni?: string;
    weaveNode?: string;
    windowsPodInfraContainer?: string;
}

export interface ClusterUpgradeStrategy {
    drain?: boolean;
    drainInput: outputs.ClusterUpgradeStrategyDrainInput;
    maxUnavailableControlplane?: string;
    maxUnavailableWorker?: string;
}

export interface ClusterUpgradeStrategyDrainInput {
    deleteLocalData?: boolean;
    force?: boolean;
    gracePeriod?: number;
    ignoreDaemonSets?: boolean;
    timeout?: number;
}

export interface ClusterWorkerHost {
    address: string;
    nodeName: string;
}
