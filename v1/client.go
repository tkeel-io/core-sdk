package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const TKeelAuthHeader = `x-tKeel-auth`

type Client struct {
	User   string
	Role   string
	Tenant string
}

func NewClient(user, role, tenant string) Client {
	return Client{User: user, Role: role, Tenant: tenant}
}

func (c *Client) Post(ctx context.Context, url string, data interface{}) Result {
	return c.invoke(ctx, url, http.MethodPost, data)
}

func (c *Client) Patch(ctx context.Context, url string, data interface{}) Result {
	return c.invoke(ctx, url, http.MethodPatch, data)
}

func (c *Client) Put(ctx context.Context, url string, data interface{}) Result {
	return c.invoke(ctx, url, http.MethodPut, data)
}

func (c *Client) Get(ctx context.Context, url string, data interface{}) Result {
	return c.invoke(ctx, url, http.MethodGet, data)
}

func (c *Client) Delete(ctx context.Context, url string, data interface{}) Result {
	return c.invoke(ctx, url, http.MethodDelete, data)
}

func (c *Client) invoke(ctx context.Context, url string, verb string, data interface{}) Result {
	var (
		err      error
		bytes    []byte
		httpReq  *http.Request
		httpResp *http.Response
	)

	if nil != data {
		// marshal payload.
		if bytes, err = json.Marshal(data); nil != err {
			return Result{Err: err}
		}

		// new request payload.
		payload := strings.NewReader(string(bytes))
		if httpReq, err = http.NewRequest(
			http.MethodPost, url, payload); nil != err {
			return Result{Err: err}
		}
	}

	// add request header.
	c.addAuthHeader(httpReq)
	httpReq.Header.Add("Content-Type", "application/json")

	// do request.
	if httpResp, err = http.DefaultClient.
		Do(httpReq); nil != err {
		return Result{Err: err}
	}

	defer httpResp.Body.Close()
	bytes, err = ioutil.ReadAll(httpResp.Body)
	return Result{Err: err, Data: bytes}
}

func (c *Client) addAuthHeader(req *http.Request) {
	authString := fmt.Sprintf("tenant=%s&user=%s&role=%s", c.Tenant, c.User, c.Role)
	req.Header.Add(TKeelAuthHeader, base64.StdEncoding.EncodeToString([]byte(authString)))
}
