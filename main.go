package main

import (
	"flag"

	service "./service"
)

func main() {
	methodFlag := flag.String("METHOD", "GET", "Method of request")
	urlFlag := flag.String("URL", "", "URL of request")
	tokenFlag := flag.String("TOKEN", "", "Bearer token of request")

	flag.Parse()

	startStress(*methodFlag, *urlFlag, *tokenFlag)
}

func startStress(method string, url string, token string) {
	switch method {
	case "GET":
		for i := 0; i < 10_000; i++ {
			service.Get(url, token)
		}
	}

}
