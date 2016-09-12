package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// Request creates and does the request, returning the body content if OK
func Request(method, url string) ([]byte, error) {
	path := fmt.Sprintf("http://%s:%d/v1%s", viper.GetString("address"), viper.GetInt("port"), url)
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(viper.GetString("username"), viper.GetString("password"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}
