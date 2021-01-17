package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string, token string) {
	var request = CreateRequest("GET", url, token)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func CreateRequest(method string, url string, token string) *http.Request {
	request, error := http.NewRequest(method, url, nil)

	if error != nil {
		log.Fatalln(error)
	}

	if token != "" {
		bearerToken := FormatBearerToken(token)

		request.Header.Add("Authorization", bearerToken)
	}

	return request
}

func FormatBearerToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
