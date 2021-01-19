package main

import (
	flags "./flags"
	services "./services"
	settings "./settings"
)

func HandleStress(f flags.Flag) {
	if f.File != "" {
		data := settings.RetrieveJSON(f.File)

		switch data.Method {
		case "GET":
			for {
				services.Get(data.Url, data.Token)
			}
		case "POST":
			for {
				services.Post(data.Url, data.Token, data.Body)
			}
		}
	}

	for {
		services.Custom(f.Method, f.Url, f.Token, f.Body)
	}
}

func main() {
	flag := flags.BuildFlags()

	HandleAttack(flag)
}
