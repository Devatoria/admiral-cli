package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// Request creates and does the request, returning the body content if OK
func Request(method, url string, bodyData interface{}) ([]byte, int, error) {
	body, err := json.Marshal(bodyData)
	if err != nil {
		return nil, 0, err
	}

	path := fmt.Sprintf("http://%s:%d/v1%s", viper.GetString("address"), viper.GetInt("port"), url)
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}

	req.SetBasicAuth(viper.GetString("username"), viper.GetString("password"))
	if bodyData != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	return data, resp.StatusCode, nil
}
