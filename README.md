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