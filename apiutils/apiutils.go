package apiutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alebaffa/wanikani-webapp/assets"
	"github.com/alebaffa/wanikani-webapp/structures"
)

/*GetStudyQueue gets the*/
func GetStudyQueue() structures.User {
	var user structures.User
	jsonResponse, err := FetchJSON("study-queue")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonResponse, &user)

	if err != nil {
		panic(err)
	}

	return user
}

/*GetKanjiList gets the list of kanji*/
func GetKanjiList(level string) structures.User {
	var user structures.User
	jsonResponse, err := FetchJSON("kanji/" + level)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonResponse, &user)

	if err != nil {
		panic(err)
	}

	return user
}

/*FetchJSON fetches the json response from the API*/
func FetchJSON(api string) ([]byte, error) {
	resp := getHTTPResponse(api)
	defer resp.Body.Close()
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return json, err
}

/* connect to the API and return the Http.Response */
func getHTTPResponse(api string) *http.Response {
	url := "https://www.wanikani.com/api/user/" + assets.APIKey + "/" + api
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}
