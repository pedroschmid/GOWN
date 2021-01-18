package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Custom(method string, url string, token string, data string) {
	payload := map[string]interface{}{"id": 1, "name": "root"}

	byts, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byts))
	req.Header.Add("Authorization", FormatBearerToken(token))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Println("Request failed, ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println("response Body:", string(body))
}

func Get(url string, token string) {
	resp, err := http.Get(url)
	resp.Header.Add("Authorization", FormatBearerToken(token))

	if err != nil {
		log.Println("Request failed, ", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string([]byte(body)))
}

func Post(url string, token string, data map[string]interface{}) {
	jsonBody, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	resp.Header.Add("Authorization", FormatBearerToken(token))

	if err != nil {
		log.Println("Request failed, ", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string([]byte(body)))
}

func FormatBearerToken(token string) string {
	if token == "" {
		return ""
	}

	return fmt.Sprintf("Bearer %s", token)
}
