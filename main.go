package main

import (
	"io/ioutil"
	"net/http"
)

func RequestFunc(url string) <-chan []byte {
	c := make(chan []byte, 1)

	go func() {
		var body []byte

		defer func() {
			c <- body
		}()

		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, _ = ioutil.ReadAll(resp.Body)

	}()

	return c
}

func main() {

}
