args:
	go run main.go --METHOD='POST' --URL='http://localhost:3000/' --BODY='{"user":"root"}'
file:
	go run main.go --FILE='./config.json'