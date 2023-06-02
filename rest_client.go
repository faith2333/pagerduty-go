package pagerduty_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/faith2333/pagerduty-go/types"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sync"
)

type IRESTClient interface {
	WithHttpClient(client *http.Client) IRESTClient
	WithEndpoint(endpoint types.Endpoint) IRESTClient
	// WithToken the token will be replaced by the passed value
	WithToken(token string) IRESTClient
	// WithBody the body will be replaced by the passed value
	WithBody(body interface{}) IRESTClient
	// AddURLParam the k-v pair will be added to url params
	AddURLParam(key, value string) IRESTClient
	// WithURLParams the url params will be replaced by the passed value
	WithURLParams(urlParams map[string]string) IRESTClient
	POST() IRESTClient
	// GET pathParams is a variable you can choose to pass or not, if you pass it will be added after the url with slash
	// eg: if you request https://example.com/test and you pass the urlParams="123321" then the real url you request is https://example.com/test/123321
	// and the urlParams just the first value take effect.
	GET(pathParams ...string) IRESTClient
	PUT() IRESTClient
	DELETE() IRESTClient
	Do(ctx context.Context) ([]byte, error)
}

type defaultRestClient struct {
	lock         *sync.RWMutex
	getPathParam string
	method       string
	endpoint     types.Endpoint
	token        string
	body         interface{}
	urlParams    map[string]string
	httpClient   *http.Client
}

func NewDefaultRestClient() IRESTClient {
	return &defaultRestClient{
		lock:       &sync.RWMutex{},
		httpClient: &http.Client{},
	}
}

func (dClient *defaultRestClient) WithHttpClient(httpClient *http.Client) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.httpClient = httpClient
	return dClient
}

func (dClient *defaultRestClient) WithEndpoint(endpoint types.Endpoint) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.endpoint = endpoint
	return dClient
}

func (dClient *defaultRestClient) WithToken(token string) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.token = token
	return dClient
}

func (dClient *defaultRestClient) POST() IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "POST"
	return dClient
}

func (dClient *defaultRestClient) GET(pathParams ...string) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	if len(pathParams) != 0 {
		dClient.getPathParam = pathParams[0]
	}
	dClient.method = "GET"
	return dClient
}

func (dClient *defaultRestClient) PUT() IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "PUT"
	return dClient
}

func (dClient *defaultRestClient) DELETE() IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.method = "DELETE"
	return dClient
}

func (dClient *defaultRestClient) WithBody(body interface{}) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.body = body
	return dClient
}

func (dClient *defaultRestClient) AddURLParam(key, value string) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	if dClient.urlParams == nil {
		dClient.urlParams = make(map[string]string)
	}
	dClient.urlParams[key] = value
	return dClient
}

func (dClient *defaultRestClient) WithURLParams(urlParams map[string]string) IRESTClient {
	dClient.lock.Lock()
	defer dClient.lock.Unlock()

	dClient.urlParams = urlParams
	return dClient
}

func (dClient *defaultRestClient) Do(ctx context.Context) ([]byte, error) {
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

	url := dClient.endpoint.String()
	if dClient.getPathParam != "" {
		url += "/" + dClient.getPathParam
	}

	bodyBuf, err := json.Marshal(dClient.body)
	if err != nil {
		return nil, errors.Errorf("marshal request body %q failed:", dClient.body)
	}

	req, err := http.NewRequest(dClient.method, url, bytes.NewBuffer(bodyBuf))
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

	fmt.Println(req.URL.String())

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

	errResp := &types.ErrResponse{}
	err = json.Unmarshal(respBytes, &errResp)
	if err == nil {
		if errResp.Error != nil {
			return nil, errors.New(errResp.Error.Message)
		}
	}

	return respBytes, nil
}
