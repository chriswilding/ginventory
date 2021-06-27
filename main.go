package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	adaptor := gorillamux.New(router)

	lambda.Start(adaptor)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
