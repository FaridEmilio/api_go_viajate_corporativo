package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/dtos/commonsdtos"
)

type RunEndpoint interface {
	RunEndpoint(method, endpoint string, headers map[string]string, body interface{}, queryParams map[string]string, logRequest bool, response interface{}) error
}

type runendpoint struct {
	HTTPClient *http.Client
}

func NewRunEndpoint(client *http.Client) RunEndpoint {
	return &runendpoint{
		HTTPClient: client,
	}
}

func (r *runendpoint) RunEndpoint(method, endpoint string, headers map[string]string, body interface{}, queryParams map[string]string, logRequest bool, response interface{}) error {
	finalURL, err := buildURL(endpoint, queryParams)
	if err != nil {
		logs.Error("URL Parsing Error: " + err.Error())
		return err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		logs.Error("JSON Marshal Error: " + err.Error())
		return err
	}

	req, err := http.NewRequest(method, finalURL, bytes.NewBuffer(reqBody))
	if err != nil {
		logs.Error("Request Creation Error: " + err.Error())
		return err
	}

	applyHeaders(req, headers)
	return r.performRequest(req, response)
}

func buildURL(endpoint string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	query := parsedURL.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

func applyHeaders(req *http.Request, headers map[string]string) {
	defaultHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	for k, v := range defaultHeaders {
		if _, exists := headers[k]; !exists {
			req.Header.Set(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (r *runendpoint) performRequest(req *http.Request, output interface{}) error {
	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		logs.Error("HTTP Request Failed: " + err.Error())
		return fmt.Errorf("error during HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error("Reading Response Body Failed: " + err.Error())
		return err
	}

	if resp.StatusCode >= 400 {
		return r.handleAPIError(resp.StatusCode, body)
	}

	if output != nil {
		if err := json.Unmarshal(body, output); err != nil {
			logs.Error("JSON Unmarshal Error: " + err.Error())
			return err
		}
	}

	return nil
}

func (r *runendpoint) handleAPIError(statusCode int, body []byte) error {
	apiErr := commonsdtos.APIError{
		Code:        strconv.Itoa(statusCode),
		Description: "Ocurrió un error procesando la petición",
	}

	switch statusCode {
	case http.StatusUnauthorized:
		apiErr.Description = "Unauthorized"
	case http.StatusInternalServerError:
		apiErr.Description = "Internal Server Error"
	}

	_ = json.Unmarshal(body, &apiErr)
	return &apiErr
}
