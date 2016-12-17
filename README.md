## What is it for?
This simple app will allow you to obtain token from [yaas.io](https://www.yaas.io)
using implicit grant flow. Client for which you want to obtain token should be 
configured in  [builder](https://builder.yaas.io). Remember about adding redirect uri
(by default it this app uses localhost:8080/callback),
and scopes which you require.

## How to run?
To run this app you need to have [golang](https://golang.org/) installed
Then just type:
```
go run main.go
```
App will be started at [localhost:8080](http://localhost:8080/)