package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http status code: %v\n", resp.StatusCode)
		return nil, err
	}

	reader := bufio.NewReader(resp.Body)
	e, err := determineEncoding(reader)
	if err != nil {
		fmt.Printf("get encoding error: %v\n", err)
		return nil, err
	}
	r := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(r)
}

// determineEncoding is func to check the text is which encoding
func determineEncoding(r *bufio.Reader) (encoding.Encoding, error) {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8, nil
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e, nil
}
