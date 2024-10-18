package chaos

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

func NewTestRunRBAC(scope constructs.Construct, ns string) (kplus.RoleBinding, string) {
	account := newJobServiceAccount(scope, "k6-account", ns)
	role := newJobRole(scope, "k6-role", ns)

	return role.Bind(account), *account.Metadata().Name()
}

func newJobServiceAccount(scope constructs.Construct, id string, ns string) kplus.ServiceAccount {
	name := fmt.Sprintf("k6-%s", ns)
	account := kplus.NewServiceAccount(scope, jsii.String(id), &kplus.ServiceAccountProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(name),
			//Namespace: jsii.String(ns),
		},
	})
	return account
}

func newJobRole(scope constructs.Construct, id string, ns string) kplus.Role {
	name := fmt.Sprintf("k6-%s", ns)
	role := kplus.NewRole(scope, jsii.String(id), &kplus.RoleProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(name),
			//Namespace: jsii.String(ns),
		},
		Rules: &[]*kplus.RolePolicyRule{
			{
				Resources: &[]kplus.IApiResource{
					kplus.ApiResource_Custom(&kplus.ApiResourceOptions{
						ApiGroup:     jsii.String("k6.io"),
						ResourceType: jsii.String("testruns"),
					}),
				},
				Verbs: &[]*string{
					jsii.String("create"),
					jsii.String("delete"),
					jsii.String("get"),
					jsii.String("list"),
					jsii.String("patch"),
					jsii.String("update"),
					jsii.String("watch"),
				},
			},
		},
	})
	return role
}
