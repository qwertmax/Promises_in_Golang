# Promises_in_Golang

Simple implementation of http.Get (not a concurrent)

```go
func RequestSimple(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
```

```bash
length = 30240
length2 = 29789
length3 = 21492
length4 = 10389
length5 = 1658
length6 = 2995
length7 = 102569
length8 = 50118
9.428153889s
```

```go
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
```

```bash
length = 30240
length2 = 29789
length3 = 21492
length4 = 10389
length5 = 1658
length6 = 2995
length7 = 102569
length8 = 50118
3.570775871s
```
