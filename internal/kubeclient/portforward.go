package kubeclient

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type PortForwardParams struct {
	HostPort   int
	TargetPort int
	Streams    genericiooptions.IOStreams
	StopChan   chan struct{}
	ReadyChan  chan struct{}
}

func PortForwardService(service *v1.Service, hostPort int, targetPort int, stopCh chan struct{}) error {

	config := os.Getenv("KUBECONFIG")
	if config == "" {
		return fmt.Errorf("KUBECONFIG environment variable not set")
	}

	readyChan := make(chan struct{})

	streams := genericiooptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-sigs
		fmt.Printf("Shutting down port forward from %d to port %d on %s\n ", hostPort, targetPort, service.Name)
		close(stopCh)
		wg.Done()
	}()

	params := PortForwardParams{
		HostPort:   hostPort,
		TargetPort: targetPort,
		Streams:    streams,
		StopChan:   stopCh,
		ReadyChan:  readyChan,
	}
	go func() {
		err := portForwardService(service, &params, config)
		if err != nil {
			//TODO: Not this
			panic(err)
		}
	}()
	select {
	case <-readyChan:
		break
	}
	fmt.Println("Port forwarding on port " + fmt.Sprint(hostPort))
	wg.Wait()
	return nil
}

func portForwardService(service *v1.Service, params *PortForwardParams, config string) error {
	path := fmt.Sprintf("/api/v1/namespaces/%s/services/%s/portforward", service.Namespace, service.Name)
	//hostIp := strings.TrimLeft(k.kubeClient.RESTClient().hos)

	restConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return err
	}
	if restConfig == nil {
		return fmt.Errorf("unable to build rest config")
	}
	hostIp := strings.TrimLeft(restConfig.Host, "htps:/")

	transport, upgrader, err := spdy.RoundTripperFor(restConfig)
	if err != nil {
		return err
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, &url.URL{Scheme: "https", Host: hostIp, Path: path})
	fw, err := portforward.New(dialer, []string{fmt.Sprintf("%d:%d", params.HostPort, params.TargetPort)}, params.StopChan, params.ReadyChan, os.Stdout, os.Stderr)
	if err != nil {
		return err
	}
	return fw.ForwardPorts()
}

func (k *k8sClient) portForwardPod() error {
	return nil
}
