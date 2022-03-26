package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var output = flag.String("o", "", "write output into file") // the emty string "" is the default value. It changes if there is userinput.

func checkURL(url string) string {
	if strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url

}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Can only takes one URL as input.")
		os.Exit(1)
	}
	url := args[0]
	url = checkURL(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Could not load %s: %v", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	output := os.Stdout
	io.Copy(output, resp.Body)
}
