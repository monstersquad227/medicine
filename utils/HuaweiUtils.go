package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"medicine/config"
	"net/http"
	"net/url"
	"strings"
)

func GetHuaweiAccessToken(code string) (string, error) {
	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("client_id", config.GlobalConfig.Huawei.ClientID)
	body.Set("client_secret", config.GlobalConfig.Huawei.ClientSecret)

	request, err := http.NewRequest(http.MethodPost, config.GlobalConfig.Huawei.Oauth2URL, strings.NewReader(body.Encode()))
	if err != nil {
		return err.Error(), err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err.Error(), err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error(), err
	}

	if response.StatusCode != http.StatusOK {
		return " ", errors.New("请求失败")
	}

	type res struct {
		Scope        string `json:"scope"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		IdToken      string `json:"id_token"`
	}
	var tokenResp res
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		return err.Error(), err
	}

	return tokenResp.AccessToken, nil
}

func GetHuaweiUserInfo(accessToken string) (string, string, error) {
	body := url.Values{}
	body.Set("access_token", accessToken)

	request, err := http.NewRequest(http.MethodPost, config.GlobalConfig.Huawei.AccountUrl, strings.NewReader(body.Encode()))
	if err != nil {
		return "", "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", "", err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	type res struct {
		UnionID           string `json:"unionID"`
		LoginMobileNumber string `json:"loginMobileNumber"`
	}
	var tokenResp res
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		return "", "", err
	}

	return tokenResp.UnionID, tokenResp.LoginMobileNumber, nil
}
