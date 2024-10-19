package kubeclient

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	v1 "k8s.io/api/core/v1"
	"log"
	"sync"
	"time"
)

const (
	argoCDServerName = "argo-helm-argocd-server"
)

type ServiceInfo struct {
	Name        string
	Namespace   string
	CreatedAt   time.Time
	Labels      map[string]string
	Annotations map[string]string
	Status      string
}

func (k *k8sClient) CheckArgoExists(ctx context.Context, ns string) (*ServiceInfo, error) {
	log.Printf("CheckArgoExists: ns: %v\n", ns)
	res, err := k.getService(ctx, ns, argoCDServerName)

	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return &ServiceInfo{
		Name:        res.Name,
		Namespace:   res.Namespace,
		CreatedAt:   res.CreationTimestamp.Time,
		Labels:      res.Labels,
		Annotations: res.Annotations,
		Status:      res.Status.String(),
	}, nil
}

func (k *k8sClient) WaitAllArgoPods(ctx context.Context, ns string) error {
	pods, err := k.getRelatedPods(ctx, ns, "argocd")

	if err != nil {
		return err
	}
	if pods == nil || len(pods.Items) == 0 {
		return fmt.Errorf("no Argo CD pods found in namespace %s", ns)
	}
	log.Println("Waiting for Argo CD pods to be ready...")
	for _, pod := range pods.Items {
		log.Printf("Waiting for Argo CD pod %s to be ready", pod.Name)
	}
	return k.waitServicePods(ctx, ns, pods)
}

func (k *k8sClient) SetupArgoLb(ctx context.Context, ns, envFile string) error {
	hasIp, err := k.checkArgoExternalIp(ctx, ns, envFile)
	if err != nil {
		return err
	}
	if !hasIp {
		info, err := k.patchArgoToLB(ctx, ns, envFile)
		if err != nil {
			return err
		}
		if info != nil {
			log.Printf("\n Created Argo CD lb service %+v", info)
		}
	}
	log.Println("Argo CD lb service has been setup")
	return nil
}

func (k *k8sClient) checkArgoExternalIp(ctx context.Context, ns, envFIle string) (bool, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()
	if err := k.WaitAllArgoPods(timeoutCtx, "argocd"); err != nil {
		return false, err
	}

	service, err := k.getService(ctx, ns, argoCDServerName)
	if err != nil {
		return false, err
	}
	if service == nil {
		return false, fmt.Errorf("no Argo CD service found in namespace %s", ns)
	}

	if service.Spec.Type != v1.ServiceTypeLoadBalancer {
		return false, nil
	}

	if service.Spec.ExternalIPs == nil {
		return false, nil
	}

	if service.Spec.ExternalIPs != nil && len(service.Spec.ExternalIPs) != 0 {
		if err = setArgoAddress(service.Spec.ExternalIPs[0], envFIle); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil

}

func (k *k8sClient) patchArgoToLB(ctx context.Context, ns string, env string) (*LoadBalancerInfo, error) {
	info, updateErr := k.exposeServiceAsLoadBalancer(ctx, ns, argoCDServerName)
	if updateErr != nil {
		return nil, updateErr
	}
	if info == nil {
		return nil, fmt.Errorf("error getting updated Argo CD server info")
	}
	if info.ExternalIPs == nil || len(info.ExternalIPs) == 0 {
		return nil, fmt.Errorf("error getting external IP for Argo CD Server")
	}

	if err := setArgoAddress(info.ExternalIPs[0], env); err != nil {
		return nil, err
	}

	return info, nil
}

func setArgoAddress(address, env string) error {
	envFile := ".env"
	if env != "" {
		envFile += "." + env
	}
	var appEnv map[string]string
	appEnv, err := godotenv.Read(envFile)
	if err != nil {
		return fmt.Errorf("error reading env file %s: %v", envFile, err)
	}
	appEnv["ARGOCD_ADRESS"] = address
	if err = godotenv.Write(appEnv, envFile); err != nil {
		return fmt.Errorf("error updating env file %s: %v", envFile, err)
	}
	return nil
}

func (k *k8sClient) portForwardToArgoCD(ctx context.Context, ns string, port int, wg *sync.WaitGroup, stopCh chan struct{}) error {

	service, err := k.getService(ctx, ns, argoCDServerName)
	if err != nil {
		return err
	}
	if service == nil {
		return fmt.Errorf("argocd-server service not found")
	}
	var pfErr error
	go func(pfErr error) {
		defer wg.Done()
		err = PortForwardService(service, port, 8080, stopCh)
		if err != nil {
			pfErr = err
		}
		select {
		case <-stopCh:
			return
		case <-ctx.Done():
			return
		}
	}(pfErr)
	return pfErr

}
