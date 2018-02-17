package zabbix

import (
	"net/http"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
)

func NewZabbix(url string) (Zabbix) {
	//Zabbix type consctructor with authentication
	z := Zabbix{url,nil,1}
	return z
}

func (z *Zabbix) ZRequest(body []byte) (map[string]interface{}, error) {
	//Wrapper to call request as a class method
	return ZApiRequest(&z.ApiUrl , body)
}

func buildRequest(url *string, body []byte) (*http.Request , error) {
	req, err := http.NewRequest(http.MethodPost, *url, bytes.NewBuffer(body))
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

func ZApiRequest(url *string, body []byte) (map[string]interface{},error) {
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

func (z *Zabbix) MakeBody(method string, params interface{}) ([]byte,error) {
	return json.Marshal(ApiBody{Method: method, Jsonrpc:"2.0", Auth:z.Auth, Id:z.AuthId, Params:params})
}