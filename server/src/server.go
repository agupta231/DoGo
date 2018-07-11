package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "log"
    "net/http"
)

func main() {
		log.Println("API Server Started")

    router := httprouter.New()
    router.GET("/hello-world", helloWorld)

    http.ListenAndServe(":9876", router)
}

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Hello World!")
    log.Println("Hit on Hello World!")
}
