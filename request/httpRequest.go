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
	proxyUrl             string
	timeout              int
	randomWaitTimeoutMin int
	randomWaitTimeoutMax int
}

func NewHttpRequest(headers map[string]interface{}, proxyUrl string, timeout, randomWaitTimeoutMin, randomWaitTimeoutMax int) *HttpRequest {
	return &HttpRequest{
		headers:              headers,
		proxyUrl:             proxyUrl,
		timeout:              timeout,
		randomWaitTimeoutMin: randomWaitTimeoutMin,
		randomWaitTimeoutMax: randomWaitTimeoutMax,
	}
}

func (r *HttpRequest) PostJson(url string, jsonData interface{}) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(r.timeout) * time.Second}
	if r.proxyUrl != "" {
		proxy, err := url2.Parse(r.proxyUrl)
		if err == nil {
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		} else {
			return nil, err
		}

	}

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jsonBody, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	for k, v := range r.headers {
		req.Header.Set(k, v.(string))
	}
	req.Header.Set("Content-Type", "application/json")
	sleepTime := utils.RandomNum(r.randomWaitTimeoutMax, r.randomWaitTimeoutMin)
	MainInsp.Print(LEVEL_INFO, Text(fmt.Sprintf("GET Method Random sleep %ds", sleepTime)))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return data, nil
}

func (r *HttpRequest) Get(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(r.timeout) * time.Second,
	}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if r.proxyUrl != "" {
		proxy, err := url2.Parse(r.proxyUrl)
		if err == nil {
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		} else {
			return nil, err
		}

	}

	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range r.headers {
		req.Header.Set(k, v.(string))
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	resp, err := client.Do(req)
	sleepTime := utils.RandomNum(r.randomWaitTimeoutMax, r.randomWaitTimeoutMin)
	MainInsp.Print(LEVEL_INFO, Text(fmt.Sprintf("GET Method Random sleep %ds", sleepTime)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return data, nil
}
