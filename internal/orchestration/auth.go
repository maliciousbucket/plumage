package orchestration

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func AddRepoCredentials(ctx context.Context, argoClient ArgoClient, envFile string) error {
	return argoClient.AddRepoCredentials(ctx, envFile)
}

func SetArgoToken(ctx context.Context, kubeClient KubeClient) error {
	err := kubeClient.CreateGalahArgoAccount(ctx, "argocd")
	if err != nil {
		return err
	}

	secret, err := kubeClient.GetArgoPassword(ctx, "argocd")
	if err != nil {
		return err
	}

	address := os.Getenv("ARGOCD_ADDRESS")
	if address == "" {
		return errors.New("ARGO_ADDRESS environment variable not set")
	}
	token, err := getArgoAuthToken(address, secret)
	if err != nil {
		return err
	}
	if token == "" {
		return errors.New("token empty")
	}
	env, err := godotenv.Read(".env.test")
	if err != nil {
		return err
	}
	env["ARGOCD_TOKEN"] = token
	err = godotenv.Write(env, ".env.test")
	if err != nil {
		return nil
	}
	return nil

}

func getArgoAuthToken(url, secret string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "admin",
		Password: secret,
	}
	address := fmt.Sprintf("https://%s/api/v1/session", url)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, address, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err

	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	token := struct {
		Token string `json:"token"`
	}{}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respData, &token)
	if err != nil {
		return "", err
	}
	key, err := getAccountAPIKey(httpClient, token.Token, url)
	if err != nil {
		return "", err
	}
	return key, nil
}

func getAccountAPIKey(client *http.Client, token, url string) (string, error) {
	path := fmt.Sprintf("https://%s/api/v1/account/galah/token", url)
	body := struct {
		Name string `json:"name"`
	}{
		Name: "galah",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	key := struct {
		Token string `json:"token"`
	}{}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respData, &key)
	if err != nil {
		return "", err
	}
	return key.Token, nil
}
