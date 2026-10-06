package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	*http.Client
}

func New() *Client {
	var cl Client
	transport := &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSClientConfig:     &tls.Config{},
		MaxIdleConnsPerHost: 10,
	}

	cl.Client = &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
	return &cl

}

func (cl *Client) ValidateToken(token string) (bool, error) {
	introspectURL := "https://auth.redbeaver.ru/realms/local-dev/protocol/openid-connect/token/introspect"

	clientID := "local-go-api"
	clientSecret := "Jfr9m6RUXhLtLe1PrxdGn5jDXigTtnCaOFts5t3eAEx606Z83mbpKl6kAw8bLyPAXytAEo4uVQQYxSmUROCKrP"

	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest(
		http.MethodPost,
		introspectURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := cl.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf(
			"keycloak introspection returned %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result struct {
		Active   bool   `json:"active"`
		Sub      string `json:"sub"`
		Username string `json:"username"`
		ClientID string `json:"client_id"`
		Exp      int64  `json:"exp"`
		Iss      string `json:"iss"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return false, fmt.Errorf("failed to parse keycloak response: %w", err)
	}

	if !result.Active {
		return false, nil
	}

	return true, nil
}
