package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func RequestSimple(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

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

func RequestFeatureFunc(url string) func() ([]byte, error) {
	var body []byte
	var err error

	c := make(chan struct{}, 1)

	go func() {
		defer close(c)

		var resp *http.Response
		resp, err = http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
	}()

	return func() ([]byte, error) {
		<-c
		return body, err
	}
}

var version = flag.Int("v", 1, "version v=1")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(4)

	switch *version {
	case 1:
		start := time.Now()

		// simple way
		body := RequestSimple("http://www.dmv.com")
		body2 := RequestSimple("http://www.wealthprep.ca")
		body3 := RequestSimple("http://www.carros.com.do")
		body4 := RequestSimple("http://www.mbhs.com/")

		elapsed := time.Since(start)

		fmt.Println("version #1")
		fmt.Println("length =", len(body))
		fmt.Println("length2 =", len(body2))
		fmt.Println("length3 =", len(body3))
		fmt.Println("length4 =", len(body4))
		fmt.Println(elapsed)

		break

	case 2:
		start := time.Now()

		// concurrent way
		page := RequestFunc("http://www.dmv.com")
		page2 := RequestFunc("http://www.wealthprep.ca")
		page3 := RequestFunc("http://www.carros.com.do")
		page4 := RequestFunc("http://www.mbhs.com")

		body := <-page
		body2 := <-page2
		body3 := <-page3
		body4 := <-page4

		elapsed := time.Since(start)

		fmt.Println("version #2")
		fmt.Println("length =", len(body))
		fmt.Println("length2 =", len(body2))
		fmt.Println("length3 =", len(body3))
		fmt.Println("length4 =", len(body4))
		fmt.Println(elapsed)

		break

	case 3:
		start := time.Now()

		page := RequestFeatureFunc("http://www.dmv.com")
		page2 := RequestFeatureFunc("http://www.wealthprep.ca")
		page3 := RequestFeatureFunc("http://www.carros.com.do")
		page4 := RequestFeatureFunc("http://www.mbhs.com")
		body, err := page()
		body2, err := page2()
		body3, err := page3()
		body4, err := page4()

		elapsed := time.Since(start)

		fmt.Println("version #3")
		fmt.Printf("error %v\n", err)
		fmt.Println("length =", len(body))
		fmt.Println("length2 =", len(body2))
		fmt.Println("length3 =", len(body3))
		fmt.Println("length4 =", len(body4))
		fmt.Println(elapsed)

		break

	case 4:
		start := time.Now()

		body := RequestSimple("http://www.dmv.com")
		body2 := RequestSimple("http://www.wealthprep.ca")
		body3 := RequestSimple("http://www.carros.com.do")
		body4 := RequestSimple("http://www.mbhs.com")
		body5 := RequestSimple("http://chatter.ru/chat")
		body6 := RequestSimple("http://ad.nl")
		body7 := RequestSimple("http://www.adweek.com")
		body8 := RequestSimple("http://upwork.com")

		elapsed := time.Since(start)

		fmt.Println("version #3")
		fmt.Println("length =", len(body))
		fmt.Println("length2 =", len(body2))
		fmt.Println("length3 =", len(body3))
		fmt.Println("length4 =", len(body4))
		fmt.Println("length5 =", len(body5))
		fmt.Println("length6 =", len(body6))
		fmt.Println("length7 =", len(body7))
		fmt.Println("length8 =", len(body8))
		fmt.Println(elapsed)

		break

	case 5:
		start := time.Now()

		// concurrent way
		page := RequestFunc("http://www.dmv.com")
		page2 := RequestFunc("http://www.wealthprep.ca")
		page3 := RequestFunc("http://www.carros.com.do")
		page4 := RequestFunc("http://www.mbhs.com")
		page5 := RequestFunc("http://chatter.ru/chat")
		page6 := RequestFunc("http://ad.nl")
		page7 := RequestFunc("http://www.adweek.com")
		page8 := RequestFunc("http://upwork.com")

		body := <-page
		body2 := <-page2
		body3 := <-page3
		body4 := <-page4
		body5 := <-page5
		body6 := <-page6
		body7 := <-page7
		body8 := <-page8

		elapsed := time.Since(start)

		fmt.Println("version #2")
		fmt.Println("length =", len(body))
		fmt.Println("length2 =", len(body2))
		fmt.Println("length3 =", len(body3))
		fmt.Println("length4 =", len(body4))
		fmt.Println("length5 =", len(body5))
		fmt.Println("length6 =", len(body6))
		fmt.Println("length7 =", len(body7))
		fmt.Println("length8 =", len(body8))
		fmt.Println(elapsed)

		break

	}

}
