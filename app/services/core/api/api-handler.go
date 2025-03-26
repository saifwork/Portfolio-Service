// api/api.go
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/saifwork/portfolio-service.git/app/services/core/responses"
)

type APIClient interface {
	GetAPIRequest(ctx context.Context, apiURL string, responseData any) error
	PostAPIRequest(ctx context.Context, apiURL string, requestData interface{}, responseData interface{}) error
	PatchAPIRequest(ctx context.Context, apiURL string, requestData interface{}, responseData interface{}) error
}

type Client struct{}

func (c *Client) GetAPIRequest(ctx context.Context, apiURL string, responseData any) error {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body and decode it into the responseData
	if err := DecodeAPIResponse(resp.Body, responseData); err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}

	// Type assertion to check for ResponseDto
	response, ok := responseData.(*responses.ResponseDto)
	if !ok {
		return fmt.Errorf("unexpected response structure")
	}

	// Handle success and error within the ResponseDto
	if !response.Success {
		if response.Error != nil {
			return fmt.Errorf(response.Error.Message)
		}
		return fmt.Errorf("API error: unknown failure")
	}

	return nil
}

func (c *Client) PostAPIRequest(ctx context.Context, apiURL string, requestData interface{}, responseData interface{}) error {
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return fmt.Errorf("failed to parse API URL: %v", err)
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to encode request payload: %v", err)
	}

	resp, err := http.Post(parsedURL.String(), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	if err := DecodeAPIResponse(resp.Body, responseData); err != nil {
		return err
	}

	return nil
}

func (c *Client) PatchAPIRequest(ctx context.Context, apiURL string, requestData interface{}, responseData interface{}) error {
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return fmt.Errorf("failed to parse API URL: %v", err)
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to encode request payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", parsedURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create PATCH request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make PATCH API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := DecodeAPIResponse(resp.Body, responseData); err != nil {
		return err
	}

	return nil
}

func DecodeAPIResponse(body io.Reader, response interface{}) error {
	if err := json.NewDecoder(body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode API response: %v", err)
	}
	return nil
}

func ConvertDataToStruct(data any, target interface{}) error {
	if dataMap, ok := data.(map[string]any); ok {
		dataBytes, err := json.Marshal(dataMap)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %v", err)
		}

		if err := json.Unmarshal(dataBytes, target); err != nil {
			return fmt.Errorf("failed to unmarshal data into struct: %v", err)
		}

		return nil
	}

	return fmt.Errorf("data is not in the expected format")
}
