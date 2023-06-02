package pagerduty_go

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/faith2333/pagerduty-go/types"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sync"
)

type defaultPagerDutyClient struct {
	lock       *sync.RWMutex
	method     string
	endpoint   types.Endpoint
	token      string
	body       interface{}
	urlParams  map[string]string
	httpClient *http.Client
}

func NewDefaultPagerDutyClient() Interface {
	return &defaultPagerDutyClient{
		lock:       &sync.RWMutex{},
		httpClient: &http.Client{},
	}
}

func (dClient *defaultPagerDutyClient) WithHttpClient(httpClient *http.Client) Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.httpClient = httpClient
	return dClient
}

func (dClient *defaultPagerDutyClient) WithEndpoint(endpoint types.Endpoint) Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.endpoint = endpoint
	return dClient
}

func (dClient *defaultPagerDutyClient) WithToken(token string) Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.token = token
	return dClient
}

func (dClient *defaultPagerDutyClient) POST() Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "POST"
	return dClient
}

func (dClient *defaultPagerDutyClient) GET() Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "GET"
	return dClient
}

func (dClient *defaultPagerDutyClient) PUT() Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "PUT"
	return dClient
}

func (dClient *defaultPagerDutyClient) DELETE() Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "DELETE"
	return dClient
}

func (dClient *defaultPagerDutyClient) WithBody(body interface{}) Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.body = body
	return dClient
}

func (dClient *defaultPagerDutyClient) WithURLParams(urlParams map[string]string) Interface {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.urlParams = urlParams
	return dClient
}

func (dClient *defaultPagerDutyClient) Do(ctx context.Context) (*types.Response, error) {
	dClient.lock.RLock()
	defer dClient.lock.RUnlock()

	select {
	case <-ctx.Done():
		return nil, errors.New("context done received")
	default:
	}
	if dClient.method == "" {
		return nil, errors.New("please call one of the POST,GET,PUT,DELETE before call Do method")
	}
	if dClient.endpoint == "" {
		return nil, errors.New("please call WithEndpoint method and pass non-empty endpoint into it before call Do method")
	}
	if dClient.token == "" {
		return nil, errors.New("please call WithToken method and pass non-empty token into it before call Do method")
	}

	bodyBuf, err := json.Marshal(dClient.body)
	if err != nil {
		return nil, errors.Errorf("marshal request body %q failed:", dClient.body)
	}

	req, err := http.NewRequest(dClient.method, dClient.endpoint.String(), bytes.NewBuffer(bodyBuf))
	if err != nil {
		return nil, errors.Wrap(err, "make request failed")
	}

	// Add url params for request
	if dClient.urlParams != nil {
		q := req.URL.Query()
		for k, v := range dClient.urlParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add("Authorization", "Token token="+dClient.token)
	req.Header.Add("Content-Type", "application/json")

	rawResp, err := dClient.httpClient.Do(req)
	defer func() {
		_ = rawResp.Body.Close()
	}()
	if err != nil {
		return nil, errors.Wrap(err, "request pagerduty failed")
	}

	if rawResp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("request pagerduty failed, response code is %d", rawResp.StatusCode)
	}

	respBytes, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body failed")
	}

	resp := &types.Response{}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response body failed")
	}

	return resp, nil
}
