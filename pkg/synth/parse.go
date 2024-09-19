package synth

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/k8s"
	"github.com/maliciousbucket/plumage/pkg/config"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	plTemp "github.com/maliciousbucket/plumage/pkg/plumage-template/manifests"
)

func FromChart(scope constructs.Construct, id string) constructs.Construct {
	mep := constructs.NewConstruct(scope, jsii.String(id))
	chart := FromManifests(mep, "chart")
	fmt.Println(chart.ApiObjects())

	return mep
}

func FromManifests(scope constructs.Construct, id string) cdk8s.Chart {

	chartAName := "yeet"
	chart := cdk8s.NewChart(scope, jsii.String(chartAName), nil)

	include := cdk8s.NewInclude(scope, jsii.String(id), &cdk8s.IncludeProps{
		Url: jsii.String("./dist/plumage.compose.yaml"),
	})

	var test cdk8s.JsonPatch

	for _, obj := range *include.ApiObjects() {
		if *obj.Metadata().Name() == "traefik-ingress-controller" {
			fmt.Println("ok")
			//test = cdk8s.JsonPatch_Test(jsii.String("/spec/minReadySeconds"), jsii.Number(1))
			//obj.AddJsonPatch(&test)

			//fmt.Println(obj.Name())
			//
			at := cdk8s.ApiObject_Of(obj)

			//fmt.Println("Api obj?")
			//fmt.Println(at)
			//fmt.Println(*at.Name())

			//patch := cdk8s.JsonPatch_Test(jsii.String("/spec/replicas"), jsii.Number(1))
			if *at.Kind() == "Deployment" {
				fmt.Println("Patch test")
				//fmt.Println(patch)
				//obj.AddJsonPatch(patch)
				fmt.Println("Adding Test")
				rep := cdk8s.JsonPatch_Replace(jsii.String("/spec/replicas"), jsii.Number(8))
				obj.AddJsonPatch(rep)
				fmt.Print(obj)

				add := cdk8s.JsonPatch_Add(jsii.String("/metadata/labels/hello"), jsii.String("yeet"))

				cdk8s.JsonPatch_Apply(obj, add)
				containerPath := "/spec/template/spec/containers/0/env"
				envt := k8s.EnvVar{
					Name:  jsii.String("hello"),
					Value: jsii.String("goodbye"),
				}

				containerPatch := cdk8s.JsonPatch_Add(jsii.String(containerPath), envt)
				at.AddJsonPatch(containerPatch)
				//cdk8s.JsonPatch_Apply(obj, containerPatch)
				name := fmt.Sprintf("hello-%s", *at.Name())
				newMap := kplus.NewConfigMap(scope, jsii.String(name), &kplus.ConfigMapProps{
					Data: &map[string]*string{
						"Key1": jsii.String("value1"),
					},
				})

				newEnvPath := "/spec/template/spec/containers/0/envFrom"

				newEnv := k8s.EnvFromSource{
					ConfigMapRef: &k8s.ConfigMapEnvSource{
						Name:     newMap.Name(),
						Optional: nil,
					},
					Prefix:    nil,
					SecretRef: nil,
				}

				newEnvPatch := cdk8s.JsonPatch_Add(jsii.String(newEnvPath), newEnv)
				at.AddJsonPatch(newEnvPatch)

			}

			if *at.Kind() == "Service" {
				//fmt.Println("Service Test test")
				//fmt.Println(*at.Name())
				//kspec := &compose.ServiceSpec{
				//	Ports: &[]*compose.ServicePort{
				//		&compose.ServicePort{
				//			Port:       jsii.Number(90),
				//			Name:       jsii.String("HELLO"),
				//			TargetPort: compose.IntOrString_FromString(jsii.String("8080")),
				//		},
				//	},
				//
				//	Type: jsii.String("loadbalancer"),
				//}
				//
				//
				//add := cdk8s.JsonPatch_Add(jsii.String("/spec"), kspec)
				//fmt.Printf("Add: %v", add)
				//at.AddJsonPatch(add)

				//newSpec := compose.KubeService_Of(at)
				//fmt.Printf("New Obj: %v", newSpec)

				ok := k8s.KubeService_Of(obj)
				fmt.Println(ok)
				var alright k8s.KubeService
				alright = ok
				fmt.Println("Alright")
				fmt.Println(alright)
				st := cdk8s.Yaml_Stringify(alright)
				fmt.Println(*st)

				hmm := ok.(k8s.KubeService)
				fmt.Println(hmm)
				fmt.Println(*hmm.Name())

				////cName := fmt.Sprintf("c-name-%s", *at.Name())
				//
				//newName := fmt.Sprintf("%s-%s", *at.Name(), *at.Kind())

				//compose.NewKubeService_Override(, scope, jsii.String("ogn"), &compose.KubeServiceProps{
				//
				//	Metadata: &compose.ObjectMeta{Name: jsii.String(ogName)},
				//	Spec:     nil,
				//})

			}

			fn := obj.Node().Children()
			fmt.Println(len(*fn))
			for _, child := range *fn {
				fmt.Println(child)
				fmt.Println(child.Node())
			}

			if *at.Kind() == "Deployment" && *at.Name() == "traefik-ingress-controller" {
				template2 := plTemp.ServiceTemplate{
					Name:      "traefik-ingress-controller",
					Namespace: "galah",
					Host:      "0.0.0.0",
					Paths: []*plTemp.ServicePath{
						&plTemp.ServicePath{
							Host:   "127.0.0.0",
							Prefix: "/traefik",
							Port:   80,
						},
					},
					LoadBalancer:   false,
					Middlewares:    nil,
					Retry:          nil,
					CircuitBreaker: nil,
					RateLimit:      nil,
					MonitoringEnv: []string{
						"OTEL_EXPORTER_OTLP_PROTOCOL",
					},
					MonitoringAliases: map[string]string{},
					Scaling: &plumagetemplate.ScalingConfig{
						TargetCpuAmount:     0,
						TargetCpuPercent:    75,
						TargetMemoryAmount:  0,
						TargetMemoryPercent: 60,
						MinReplicas:         2,
						MaxReplicas:         6,
						TargetReplicas:      5,
						Resources:           nil,
					},
				}

				cfg := config.CollectorConfig{OtlpExportProtocol: "http"}

				plTemp.NewService(scope, "shrek", &template2, at, &cfg)
			}

		}
	}

	fmt.Printf("Patch: %v", test)

	return chart
}

//func serviceNames(chart cdk8s.Chart) []string {
//
//}
