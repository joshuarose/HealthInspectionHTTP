package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://data.cincinnati-oh.gov/resource/2c8u-zmu9.json")
	if err == nil {
		fmt.Println(resp)
	}
}
