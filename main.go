package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var flagOutput = flag.String("o", "", "write output into file") // the emty string "" is the default value. It changes if there is userinput.

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
		log.Fatal("Takes one URL as input.")
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
	var writer io.Writer = os.Stdout
	// *flagOutput != "" -> if the var set to flag.String is not the default
	if *flagOutput != "" {
		file, err := os.OpenFile(
			*flagOutput,
			os.O_RDWR|os.O_CREATE,
			0755,
		)
		if err != nil {
			fmt.Printf("Failed to create%s\n%v", *flagOutput, err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	}
	io.Copy(writer, resp.Body)
}
