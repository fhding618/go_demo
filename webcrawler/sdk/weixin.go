package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	wxHost = "https://api.weixin.qq.com"
	appId  = "wx9a992ed9281c4e95"
	appSec = "e007f69eb21aa8622c664cdfe20d695e"
)

type wxError struct {
	ErrCode uint64 `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type wxToken struct {
	Access_token string `json:"access_token"`
	Expires_in   uint32 `json:"expires_in"`
}

type wxClient struct {
	appid  string
	appsec string
}

func NewWXClient() *wxClient {
	return &wxClient{
		appid:  appId,
		appsec: appSec,
	}
}

func (wx *wxClient) GetAppId() string {
	return wx.appid
}

func (wx *wxClient) GetAccessToken() (*wxToken, error) {
	urlToken := "%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	url := fmt.Sprintf(urlToken, wxHost, appId, appSec)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Get WX Token status is not ok, value: %v", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	token := new(wxToken)
	json.Unmarshal(body, token)
	if token.Access_token == "" {
		wxErr := new(wxError)
		json.Unmarshal(body, wxErr)
		return nil, errors.New(
			fmt.Sprintf("Get WX Token failed code: %d err: %s", wxErr.ErrCode, wxErr.ErrMsg))
	}

	return token, nil
}
