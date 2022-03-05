package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HttpRequestData struct {
	Client   *http.Client
	ProxyUrl string            `json:"proxy_url"` // ProxyUrl: "http://60.31.89.41:4231",
	Method   string            `json:"method"`
	ReqUrl   string            `json:"req_url"`
	Params   map[string]string `json:"params"`
	Data     map[string]string `json:"data"`
	Headers  map[string]string `json:"headers"`
	IsEncode bool              `json:"is_encode"`
}

// HttpRequest Customized http request, support Customized Headers、Data、Params and Proxy
func HttpRequest(requestData HttpRequestData) (string, error) {
	if requestData.Client == nil {
		return "", errors.New("client is nil")
	}
	if requestData.Method == "" {
		return "", errors.New("method is nil")
	}
	if requestData.Params != nil && len(requestData.Params) > 0 {
		param := paramToString(requestData.Params, requestData.IsEncode)
		unescape, err := url.QueryUnescape(fmt.Sprintf("%s%s%s", requestData.ReqUrl, "?", param))
		if err != nil {
			return "", err
		}
		requestData.ReqUrl = unescape
	}
	dataByte, err := json.Marshal(requestData.Data)
	if err != nil {
		return "", err
	}
	httpRequest, err := http.NewRequest(requestData.Method, requestData.ReqUrl, strings.NewReader(string(dataByte)))
	if err != nil {
		return "", err
	}
	if requestData.Headers != nil {
		for k, v := range requestData.Headers {
			httpRequest.Header.Add(k, v)
		}
	}
	if requestData.ProxyUrl != "" {
		proxy, _ := url.Parse(requestData.ProxyUrl)
		tr := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		requestData.Client.Transport = tr
	}
	resp, err := requestData.Client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	byteResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byteResp), nil
}

func paramToString(params map[string]string, isEncode bool) string {
	paramList := []string{}
	for k, v := range params {
		if v == "" {
			paramList = append(paramList, "=")
		} else {
			if !isEncode {
				paramList = append(paramList, k+"="+v)
			}
			if isEncode {
				paramList = append(paramList, k+"="+url.QueryEscape(v))
			}
		}
	}
	return strings.Join(paramList, "&")
}
