package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Setting struct {
	Method string                 `json:"method"`
	Url    string                 `json:"url"`
	Token  string                 `json:"token"`
	Body   map[string]interface{} `json:"body"`
}

func OpenFile(path string) []byte {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func RetrieveJSON(path string) Setting {
	var setting Setting

	content := OpenFile(path)

	err := json.Unmarshal(content, &setting)

	if err != nil {
		log.Fatal(err)
	}

	return setting
}
