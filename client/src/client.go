package main

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"log"
)

const baseURL string = "http://server:9876/hello-world"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Send [R] or [Q]uit?")

		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.Replace(userChoice, "\n", "", -1)

		if strings.Compare("Q", userChoice) == 0 {
			fmt.Println("Okie Dokie... Quitting!")
			return
		} else if strings.Compare("R", userChoice) == 0 {
			client := http.Client{
				Timeout: time.Second * 2,
			}

			req, err := http.NewRequest(http.MethodGet, baseURL, nil)

			if err != nil {
				log.Println("HTTP Request Creation Error")
				log.Println(err)
			}

			res, getErr := client.Do(req)

			if getErr != nil {
				log.Println("HTTP Request Execution Error")
				log.Println(getErr)
			}

			body, _ := ioutil.ReadAll(res.Body)
			log.Println(string(body))
			fmt.Println(string(body))
		} else {
			log.Println("Couldn't understand input. Please try again.")
		}
	}
}
