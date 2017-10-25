package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)

type ExtraHeader struct {
	User  string
	Role  string
	Token string
}

func do(method, url string, in interface{}, timeout time.Duration, extra *ExtraHeader) (string, error) {
	reqJson, _ := json.Marshal(in)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if extra != nil {
		if extra.User != "" {
			req.Header.Set("User", extra.User)
		}
		if extra.Role != "" {
			req.Header.Set("Role", extra.Role)
		}
		if extra.Token != "" {
			req.Header.Set("Token", extra.Token)
		}
	}
	client := http.DefaultClient
	client.Timeout = timeout
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return string(body), err
	} else {
		return "", err
	}
}

func HttpPost(url string, in interface{}, timeout time.Duration) (string, error) {
	return do("POST", url, in, timeout, nil)
}

func HttpPut(url string, in interface{}, timeout time.Duration) (string, error) {
	return do("PUT", url, in, timeout, nil)
}

func HttpGet(url string, in interface{}, timeout time.Duration) (string, error) {
	return do("GET", url, in, timeout, nil)
}

func HttpPostExtra(url string, in interface{}, timeout time.Duration, extra *ExtraHeader) (string, error) {
	return do("POST", url, in, timeout, extra)
}

func HttpPutExtra(url string, in interface{}, timeout time.Duration, extra *ExtraHeader) (string, error) {
	return do("PUT", url, in, timeout, extra)
}

func HttpGetExtra(url string, in interface{}, timeout time.Duration, extra *ExtraHeader) (string, error) {
	return do("GET", url, in, timeout, extra)
}
