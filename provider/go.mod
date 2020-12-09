module github.com/jaxxstorm/pulumi-rke/provider/v2

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.15.3
	github.com/pulumi/pulumi/pkg/v2 v2.15.1-0.20201202214525-260620430c4c
	github.com/pulumi/pulumi/sdk/v2 v2.15.1-0.20201202214525-260620430c4c
	github.com/rancher/terraform-provider-rke v1.1.5
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190916202348-b4ddaad3f8a3
	k8s.io/client-go => k8s.io/client-go v0.18.0
)
