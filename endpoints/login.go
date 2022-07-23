package endpoints

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(username string, password string) (string, error) {
	requestBody := fmt.Sprintf(`
		{"username":"%s",
		"password":"%s", 
		"supported_user_types": [
			"parent",
			"child"
    ]}`, username, password)

	req, err := http.NewRequest("POST", "https://www.easistent.com/m/login", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-App-Name", "child")
	req.Header.Set("x-client-platform", "android")
	req.Header.Set("x-client-version", "11101")

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
		return string(body), fmt.Errorf("login failed with status code %d", response.StatusCode)
	}

	return string(body), nil
}
