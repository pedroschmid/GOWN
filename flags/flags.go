package flags

import (
	"flag"
)

type Flag struct {
	Method string
	Url    string
	Token  string
	Body   string
	File   string
}

func BuildFlags() Flag {
	methodFlag := flag.String("METHOD", "GET", "Method of request")
	urlFlag := flag.String("URL", "", "URL of request")
	tokenFlag := flag.String("TOKEN", "", "Bearer token of request")
	bodyFlag := flag.String("BODY", "", "Body of request")
	fileFlag := flag.String("FILE", "", "File containing settings of request")

	flag.Parse()

	flags := Flag{
		Method: *methodFlag,
		Url:    *urlFlag,
		Token:  *tokenFlag,
		Body:   *bodyFlag,
		File:   *fileFlag,
	}

	return flags
}
