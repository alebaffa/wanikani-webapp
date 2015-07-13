package apiutils

import (
	"io/ioutil"
	"net/http"
)

/* fetch the json response from the API */
func FetchJson() ([]byte, error) {
	resp := getHttpResponse()
	defer resp.Body.Close()
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return json, err
}

/* connect to the API and return the Http.Response */
func getHttpResponse() *http.Response {
	url := "https://www.wanikani.com/api/user/user-key/study-queue"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}
