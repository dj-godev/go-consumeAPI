package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

func main() {
	url :=
		"http://data.fixer.io/api/latest?access_key=" + "3ac2f60efbdadb027dda991fd3de68a2"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {

			var result Result
			json.Unmarshal([]byte(body), &result)

			if result.Success {
				for i, v := range result.Rates {
					fmt.Println(i, v)
				}
			} else {
				var err Error
				json.Unmarshal([]byte(body), &err)
				fmt.Println(err.Error.Info)
			}

		} else {
			log.Fatal(err)
		}
		fmt.Println("Done")
	}
}
