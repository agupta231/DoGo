package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"time"
)

const baseURL string = "server:9876/hello-world"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("Send [R] or [Q]uit?")

		if userChoice == "Q" {
			fmt.Println("Okie Dokie... Quitting!")
			return
		} else if userChoice == "R" {
			client := http.Client{
				Timeout: time.Second * 2,
			}

			req, err := http.NewRequest(http.MethodGet, baseURL, nil)

			if err != nil {
				log.Fatal(err)
			}

			res, getErr := client.Do(req)

			if getErr != nil {
				log.Fatal(err)
			}

			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(body)
		} else {
			fmt.Println("Couldn't understand input. Please try again.")
		}
	}
}
