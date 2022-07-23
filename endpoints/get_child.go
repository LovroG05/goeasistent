package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetChild(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://www.easistent.com/m/me/child", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-App-Name", "child")
	req.Header.Set("x-client-platform", "android")
	req.Header.Set("x-client-version", "11101")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return string(body), fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	return string(body), nil
}
