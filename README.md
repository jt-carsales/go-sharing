#Getting started with Golang

1) Install Golang : 
https://golang.org/

2) Documentation based on examples:
https://gobyexample.com/

3) Basic commands: 

go get -v .  ==> Clone the project or dependencies

go build     ==> Build the code into a binary

go run *.go  ==> Run the code in the folder 

env GOOS=linux GOARCH=amd64 go build -o builds/gosharing  ==> Cross compile for linux X64 


# go-sharing
Implementation of a basic Golang API for PoC

The goal of this PoC is to implement three Endpoints to Get,add and delete numbers from a Mysql database

Start the application with:
go run *.go

Testing:
curl -X GET http://127.0.0.1:9090/numbers/1

curl -X POST http://127.0.0.1:9090/numbers -d '{"did":"0390934667","isocc":"AU"}'

curl -X DELETE http://127.0.0.1:9090/numbers -d '{"did":"0390934667","isocc":"AU"}'



