//curl with -vvv
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		fixedURL := addPrefix(url)

		resp, err := http.Get(fixedURL)
		if err != nil {
			log.Panic(err)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n\n", b)
		fmt.Printf("HTTP Status: %v\n", resp.Status)

	}
}

func addPrefix(url string) string {
	if !strings.HasPrefix(url, "http") {
		return string("http" + url)
	}
	return url
}
