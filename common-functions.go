package zabbix

import (
	"net/http"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
)

func buildRequest(url string, body string) (*http.Request , error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json-rpc")
	return req, err
}

func makeRequest(r *http.Request) (*http.Response,error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cl := &http.Client{Transport: tr}
	resp, err := cl.Do(r)
	return resp , err
}

func parseResponse(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	rbody, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	err := json.Unmarshal(rbody, &dat)
	return dat, err
}

func zApiRequest(url string, body string) (map[string]interface{},error) {
	var data map[string]interface{}
	req, err := buildRequest(url,body)
	if err != nil {
		return data , err
	}
	resp, err := makeRequest(req)
	if err != nil {
		return data , err
	}
	data, err = parseResponse(resp)
	return data , err
}