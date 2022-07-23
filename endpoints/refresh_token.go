package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LovroG05/goeasistent/objects"
)

func RefreshToken(refreshToken string, sessionToken string) (objects.Tokens, error) {
	requestBody, err := json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})
	if err != nil {
		return objects.Tokens{}, err
	}

	req, err := http.NewRequest("POST", "https://www.easistent.com/m/timetable/events", bytes.NewBuffer(requestBody))
	if err != nil {
		return objects.Tokens{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-App-Name", "child")
	req.Header.Set("x-client-platform", "android")
	req.Header.Set("x-client-version", "11101")
	req.Header.Set("cookie", "easistent_session="+sessionToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return objects.Tokens{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return objects.Tokens{}, err
	}

	if response.StatusCode != 200 {
		return objects.Tokens{}, fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	var tokens objects.Tokens
	err = json.Unmarshal(body, &tokens)
	if err != nil {
		return objects.Tokens{}, err
	}

	return tokens, nil
}
