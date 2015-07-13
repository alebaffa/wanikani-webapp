package apiutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alebaffa/wanikani-webapp/assets"
	"github.com/alebaffa/wanikani-webapp/structures"
)

func GetStudyQueue() structures.User {
	var user structures.User
	jsonResponse, err := FetchJson("study-queue")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonResponse, &user)

	if err != nil {
		panic(err)
	}

	return user
}

func GetKanjiList(level string) structures.User {
	var user structures.User
	jsonResponse, err := FetchJson("kanji/" + level)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonResponse, &user)

	if err != nil {
		panic(err)
	}

	return user
}

/* fetch the json response from the API */
func FetchJson(api string) ([]byte, error) {
	resp := getHttpResponse(api)
	defer resp.Body.Close()
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return json, err
}

/* connect to the API and return the Http.Response */
func getHttpResponse(api string) *http.Response {
	url := "https://www.wanikani.com/api/user/" + assets.APIKey + "/" + api
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}
