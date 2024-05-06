package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	. "github.com/Kumengda/Tarantola/runtime"
	"github.com/Kumengda/Tarantola/utils"
	"github.com/corpix/uarand"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	url2 "net/url"
	"time"
)

type HttpRequest struct {
	headers              map[string]interface{}
	timeout              int
	randomWaitTimeoutMin int
	randomWaitTimeoutMax int
	maxRetry             int
	retryFunc            func(respData []byte, respHeader http.Header, err error) bool
	client               *http.Client
}

func NewHttpRequest(headers map[string]interface{}, proxyUrl string, timeout, randomWaitTimeoutMin, randomWaitTimeoutMax int) *HttpRequest {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	if proxyUrl != "" {
		proxy, err := url2.Parse(proxyUrl)
		if err == nil {
			client.Transport = &http.Transport{
				Proxy:           http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	} else {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	return &HttpRequest{
		client:               client,
		headers:              headers,
		timeout:              timeout,
		randomWaitTimeoutMin: randomWaitTimeoutMin,
		randomWaitTimeoutMax: randomWaitTimeoutMax,
	}
}
func (r *HttpRequest) SetRetryFunc(retryFunc func(respData []byte, respHeader http.Header, err error) bool, maxRetry int) {
	r.retryFunc = retryFunc
	r.maxRetry = maxRetry
}

func (r *HttpRequest) postJson(url string, jsonData interface{}) ([]byte, http.Header, error) {
	jsonBody, err := json.Marshal(jsonData)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, nil, err
	}
	for k, v := range r.headers {
		req.Header.Set(k, v.(string))
	}
	req.Header.Set("Content-Type", "application/json")
	sleepTime := utils.RandomNum(r.randomWaitTimeoutMax, r.randomWaitTimeoutMin)
	MainInsp.Print(LEVEL_INFO, Text(fmt.Sprintf("GET Method Random sleep %ds", sleepTime)))
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, nil, err
	}

	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, nil, err
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return data, resp.Header, nil
}
func (r *HttpRequest) get(url string) ([]byte, http.Header, error) {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range r.headers {
		req.Header.Set(k, v.(string))
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	resp, err := r.client.Do(req)
	sleepTime := utils.RandomNum(r.randomWaitTimeoutMax, r.randomWaitTimeoutMin)
	MainInsp.Print(LEVEL_INFO, Text(fmt.Sprintf("GET Method Random sleep %ds", sleepTime)))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, nil, err
	}

	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, nil, err
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return data, nil, nil
}
func (r *HttpRequest) Get(url string) ([]byte, error) {
	var respData []byte
	var respHeader http.Header
	var err error
	maxRetry := r.maxRetry
	for {
		if maxRetry == -1 {
			return respData, err
		}
		respData, respHeader, err = r.get(url)
		maxRetry--
		if r.retryFunc != nil {
			isRetry := r.retryFunc(respData, respHeader, err)
			if isRetry {
				continue
			} else {
				return respData, err
			}
		} else {
			return respData, err
		}
	}
}

func (r *HttpRequest) PostJson(url string, jsonData interface{}) ([]byte, error) {
	var respData []byte
	var respHeader http.Header
	var err error
	maxRetry := r.maxRetry
	for {
		if maxRetry == -1 {
			return respData, err
		}
		respData, respHeader, err = r.postJson(url, jsonData)
		maxRetry--
		if r.retryFunc != nil {
			isRetry := r.retryFunc(respData, respHeader, err)
			if isRetry {
				continue
			} else {
				return respData, err
			}
		} else {
			return respData, err
		}
	}
}
