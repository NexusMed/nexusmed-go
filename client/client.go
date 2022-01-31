package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nexusmed/nexusmed-go/graphqljson"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type GQLRequestInfo struct {
	Request *Request
}

func NewGQLRequestInfo(r *Request) *GQLRequestInfo {
	return &GQLRequestInfo{
		Request: r,
	}
}

type Url struct {
	Base string
	Path *string
}

// Client is the http client wrapper
type Client struct {
	Client             *http.Client
	ApiKey             *string
	Url                *Url
	RequestInterceptor RequestInterceptor
}

// Request represents an outgoing GraphQL request
type Request struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
}

// New creates a new http client wrapper
func New(pathURL string, interceptors ...RequestInterceptor) *Client {
	return &Client{
		Client: http.DefaultClient,
		Url: &Url{
			Base: "https://api.nexusmed.io",
			Path: &pathURL,
		},
		RequestInterceptor: ChainInterceptor(append([]RequestInterceptor{func(ctx context.Context, requestSet *http.Request, gqlInfo *GQLRequestInfo, res interface{}, next RequestInterceptorFunc) error {
			return next(ctx, requestSet, gqlInfo, res)
		}}, interceptors...)...),
	}
}

// SetApiKey sets the Authorization header
func (c *Client) SetApiKey(key string) {
	c.ApiKey = &key
}

// SetBaseUrl sets the BaseUrl
func (c *Client) SetBaseUrl(url string) {
	c.Url.Base = url
}

// GqlErrorList is the struct of a standard graphql error response
type GqlErrorList struct {
	Errors gqlerror.List `json:"errors"`
}

func (e *GqlErrorList) Error() string {
	return e.Errors.Error()
}

// HTTPError is the error when a GqlErrorList cannot be parsed
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse represent an handled error
type ErrorResponse struct {
	// populated when http status code is not OK
	NetworkError *HTTPError `json:"networkErrors"`
	// populated when http status code is OK but the server returned at least one graphql error
	GqlErrors *gqlerror.List `json:"graphqlErrors"`
}

// HasErrors returns true when at least one error is declared
func (er *ErrorResponse) HasErrors() bool {
	return er.NetworkError != nil || er.GqlErrors != nil
}

func (er *ErrorResponse) Error() string {
	content, err := json.Marshal(er)
	if err != nil {
		return err.Error()
	}

	return string(content)
}

// the response into the given object.
func (c *Client) Post(operationName, query string, respData interface{}, vars map[string]interface{}, interceptors ...RequestInterceptor) error {
	r := &Request{
		Query:         query,
		Variables:     vars,
		OperationName: operationName,
	}
	gqlInfo := NewGQLRequestInfo(r)

	requestBody, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.Url.Base, *c.Url.Path),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return fmt.Errorf("create request struct failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	if c.ApiKey != nil && *c.ApiKey != "" {
		req.Header.Set("Authorization", *c.ApiKey)
	}

	f := ChainInterceptor(append([]RequestInterceptor{c.RequestInterceptor}, interceptors...)...)

	return f(ctx, req, gqlInfo, respData, c.do)
}

func (c *Client) do(_ context.Context, req *http.Request, _ *GQLRequestInfo, res interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return parseResponse(body, resp.StatusCode, res)
}

func parseResponse(body []byte, httpCode int, result interface{}) error {
	errResponse := &ErrorResponse{}
	isKOCode := httpCode < 200 || 299 < httpCode
	if isKOCode {
		errResponse.NetworkError = &HTTPError{
			Code:    httpCode,
			Message: string(body),
		}
	}

	// some servers return a graphql error with a non OK http code, try anyway to parse the body
	if err := unmarshal(body, result); err != nil {
		if gqlErr, ok := err.(*GqlErrorList); ok {
			errResponse.GqlErrors = &gqlErr.Errors
		} else if !isKOCode { // if is KO code there is already the http error, this error should not be returned
			return err
		}
	}

	if errResponse.HasErrors() {
		return errResponse
	}

	return nil
}

// response is a GraphQL layer response from a handler.
type response struct {
	Data   json.RawMessage `json:"data"`
	Errors json.RawMessage `json:"errors"`
}

func unmarshal(data []byte, res interface{}) error {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("failed to decode data %s: %w", string(data), err)
	}

	if resp.Errors != nil && len(resp.Errors) > 0 {
		// try to parse standard graphql error
		errors := &GqlErrorList{}
		if e := json.Unmarshal(data, errors); e != nil {
			return fmt.Errorf("faild to parse graphql errors. Response content %s - %w", string(data), e)
		}

		return errors
	}

	if err := graphqljson.UnmarshalData(resp.Data, res); err != nil {
		return fmt.Errorf("failed to decode data into response %s: %w", string(data), err)
	}

	return nil
}

type RequestInterceptorFunc func(ctx context.Context, req *http.Request, gqlInfo *GQLRequestInfo, res interface{}) error

type RequestInterceptor func(ctx context.Context, req *http.Request, gqlInfo *GQLRequestInfo, res interface{}, next RequestInterceptorFunc) error

func ChainInterceptor(interceptors ...RequestInterceptor) RequestInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, req *http.Request, gqlInfo *GQLRequestInfo, res interface{}, next RequestInterceptorFunc) error {
		chainer := func(currentInter RequestInterceptor, currentFunc RequestInterceptorFunc) RequestInterceptorFunc {
			return func(currentCtx context.Context, currentReq *http.Request, currentGqlInfo *GQLRequestInfo, currentRes interface{}) error {
				return currentInter(currentCtx, currentReq, currentGqlInfo, currentRes, currentFunc)
			}
		}

		chainedHandler := next
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req, gqlInfo, res)
	}
}
