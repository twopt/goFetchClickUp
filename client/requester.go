package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/twopt/clickup/utils"
)

//Requester interface needs an API path to request JSON data
type Requester interface {
	BuildPath() string
	GetJSON(string) []byte
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//Gets JSON data for any struct that implements Requester interface
func getJSON(apiPath string) []byte {
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	req.Header.Add("Authorization", viper.GetString("token"))
	req.Header.Add("Content-Type", "application/json")

	var client HTTPClient
	client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	return resp_body
}

func Request(r Requester) {
	fmt.Println(utils.FormatJSON(r.GetJSON(r.BuildPath())))
}
