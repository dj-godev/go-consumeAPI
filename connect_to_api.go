package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://data.fixer.io/api/latest?access_key=3ac2f60efbdadb027dda991fd3de68a2"
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(body))
		} else {
			log.Fatal(err)
		}
		fmt.Println("Done")
	}
}
