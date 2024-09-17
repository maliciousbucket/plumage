package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

func SetArgoTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-argo-token",
		Short: "set argo token",
		Run: func(cmd *cobra.Command, args []string) {

			setArgoToken()
		},
	}
	return cmd
}

// SetArgoToken TODO: Remove
func SetArgoToken() {
	setArgoToken()
}

func setArgoToken() {
	k8sClient, err := kubeclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = k8sClient.Client.PatchArgoToLB(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Client.CreateGalahArgoAccount(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	secret, err := k8sClient.Client.GetArgoPassword(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	portForwardCmd := exec.Command("kubectl", "port-forward", "svc/argocd-helm-server", "8081:443", "-n", "argocd")
	go func() {
		defer wg.Done()
		err := portForwardCmd.Run()
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				if exitErr.Sys().(syscall.WaitStatus).Signaled() && exitErr.Sys().(syscall.WaitStatus).Signal() == syscall.SIGKILL {
					log.Println("Port-forward Closed")
					return
				}
			}
			log.Fatal("Port-forward failed: ", err)
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Secret: " + secret)
	pss := fmt.Sprintf("--password %s", secret)
	loginCommand := exec.Command("argocd", "login", "localhost:8081", "--username", "admin", pss, "--insecure")
	aout, err := loginCommand.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = loginCommand.Start(); err != nil {
		log.Fatal(err)
	}
	ad, err := io.ReadAll(aout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ad))

	tokenCmd := exec.Command("argocd", "account", "generate-token", "--account", "galah")
	stdout, err := tokenCmd.StdoutPipe()

	if err = tokenCmd.Start(); err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %s\n", string(data))

	myEnv, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	myEnv["ARGOCD_TOKEN"] = string(data)
	err = godotenv.Write(myEnv, ".env")
	if err != nil {
		log.Fatal(err)
	}

	if err := portForwardCmd.Process.Kill(); err != nil {
		log.Fatal("Failed to kill port-forward process: ", err)
	}

	wg.Wait()
}

func GetHttpToken() error {
	k8sClient, err := kubeclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = k8sClient.Client.PatchArgoToLB(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Client.CreateGalahArgoAccount(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	secret, err := k8sClient.Client.GetArgoPassword(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	address, err := k8sClient.Client.GetServiceAddress(ctx, "argocd", "argocd-helm-server")
	if err != nil {
		log.Fatal(err)
	}
	token, err := getArgoAuthToken(address, secret)
	if err != nil {
		log.Fatal(err)
	}
	if token == "" {
		return errors.New("token empty")
	}
	myEnv, err := godotenv.Read(".env")
	if err != nil {
		return err
	}
	myEnv["ARGOCD_TOKEN"] = token
	err = godotenv.Write(myEnv, ".env")
	if err != nil {
		return nil
	}
	return nil

}

func getArgoAuthToken(url, secret string) (string, error) {
	client := &http.Client{}

	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "admin",
		Password: secret,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var token string
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respData, &token)
	if err != nil {
		return "", err
	}
	key, err := getAccountAPIKey(client, token, url)
	if err != nil {
		return "", err
	}
	return key, nil
}

func getAccountAPIKey(client *http.Client, token, url string) (string, error) {
	path := fmt.Sprintf("%s/api/v1/galah/token", url)
	body := struct {
		Name string `json:"name"`
	}{
		Name: "galah",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodGet, path, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var key string
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respData, &key)
	if err != nil {
		return "", err
	}
	return key, nil
}

func syncCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sync",
	}
	return cmd
}

func actionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "actions",
	}
	return cmd
}

func createAppsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-apps",
	}
	return cmd
}
