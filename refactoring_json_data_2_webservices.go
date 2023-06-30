package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apis map[int]string

func main() {
	apis = make(map[int]string)
	apis[1] =
		"http://data.fixer.io/api/latest?access_key=" +
			"3ac2f60efbdadb027dda991fd3de68a2"
	apis[2] =
		"http://api.openweathermap.org/data/2.5/weather?" +
			"q=gurugram&appid=e7704bc895b4a8d2dfd4a29d404285b6"
	fetchData(1)
	fetchData(2)
}

func fetchData(API int) {

	url := apis[API]

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)
			switch API {
			case 1:
				if result["success"] == true {
					fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					fmt.Println(result["error"].(map[string]interface{})["info"])
				}
			case 2: // for the openweathermap.org API
				if result["main"] != nil {
					fmt.Println(result["main"].(map[string]interface{})["temp"])
				} else {
					fmt.Println(result["message"])
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
