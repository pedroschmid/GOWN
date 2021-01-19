# GOWN
Tool for multiple requests and stress test

# How to run?

If u have a configuration file (should be a .json file!), run: go run main.go --FILE="./config.json"

But if u dont have a configuration file, you can pass arguments/flags by yourself, follow the list of args...

--METHOD=string
<br /> 
--URL=string
<br /> 
--TOKEN=string
<br /> 
--BODY=string

Here's an example: go run main.go --METHOD="POST" --URL="http://localhost:3000/" --BODY="{ "user": "root" }"
