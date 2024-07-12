package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// > POST /v3/push HTTP/1.1
// > Authorization: Basic N2Q0MzFlNDJkZmE2YTZkNjkzYWMyZDA0OjVlOTg3YWM2ZDJlMDRkOTVhOWQ4ZjBkMQ==
func (c *Core) DoPostRequest(ctx context.Context, action string, data any) (*Response, error) {

	b, err := json.Marshal(data)
	if err != nil {
		c.Logger.Sugar().Errorf("json marshal error: %s", err)
		return nil, err
	}
	reader := bytes.NewReader(b)
	url := fmt.Sprintf("%s%s", c.baseURL, action)
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		c.Logger.Sugar().Errorf("init %s request error: %s", action, err)
		return nil, err
	}
	req = c.handleAuth(req)

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.Logger.Sugar().Errorf("do request error: %+v", err)
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Sugar().Errorf("read response body error: %+v", err)
		return nil, err
	}

	return &Response{
		Body:   body,
		Status: resp.StatusCode,
	}, nil
}

func (c *Core) handleAuth(req *http.Request) *http.Request {
	pair := fmt.Sprintf("%s:%s", c.appKey, c.appMasterSecret)
	authString := base64.StdEncoding.EncodeToString([]byte(pair))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authString))
	req.Header.Add("Content-Type", "application/json")
	return req
}
