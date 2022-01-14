package httpclient

import (
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)

//Client http客户端
type Client struct {
	client *http.Client
}

type DealRespOpt func(*http.Response) unsafe.Pointer

type DealReqOpt func(*http.Request)

//NewClient 获得client
func NewClient() *Client {
	return &Client{client: &http.Client{}}
}

//GetClient 获取http客户端
func (cli *Client) GetClient() *http.Client {
	return cli.client
}

//DoGet get请求
func (cli *Client) DoGet(url string, dealResp DealRespOpt, dealReq DealReqOpt) error {
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if dealReq != nil {
		dealReq(req)
	}

	resp, err := cli.client.Do(req)
	if err != nil {
		return err
	}

	if dealResp != nil {
		dealResp(resp)
	}

	return nil
}

//DoPost post请求
func (cli *Client) DoPost(url, data string, dealResp DealRespOpt, dealReq DealReqOpt) error {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if dealReq != nil {
		dealReq(req)
	}

	resp, err := cli.client.Do(req)
	if err != nil {
		return err
	}

	if dealResp != nil {
		dealResp(resp)
	}

	return nil
}

//DoDelete delete请求
func (cli *Client) DoDelete(url, data string, dealResp DealRespOpt, dealReq DealReqOpt) error {
	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if dealReq != nil {
		dealReq(req)
	}

	resp, err := cli.client.Do(req)
	if err != nil {
		return err
	}

	if dealResp != nil {
		dealResp(resp)
	}

	return nil
}

//getBody 从response中获取json类型的body
func GetBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
