package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	p12path := flag.String("path", "", "Path to client P12")
	p12pass := flag.String("pass", "", "Password for client P12")
	enrollPtr := flag.Bool("enroll", false, "Bool to enroll")
	renewPtr := flag.Bool("renew", false, "Bool to renew")
	trustPtr := flag.Bool("trust", false, "Bool to retrieve trust")

	flag.Parse()

	fmt.Println("path:", *p12path)
	fmt.Println("p12pass:", *p12pass)
	fmt.Println("enroll:", *enrollPtr)
	fmt.Println("renew:", *renewPtr)
	fmt.Println("trust:", *trustPtr)

	if *enrollPtr {
		fmt.Println("Starting EST Simple Enroll.")
	} else if *renewPtr {
		fmt.Println("Starting EST Simple Reenroll.")
	} else if *trustPtr {
		fmt.Println("Getting CA trust.")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("GET", "https://postman-echo.com/get", nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		trustP7 := fmt.Sprintf("-----BEGIN P7B-----\n%d\n-----END P7B-----\n", bodyText)
		fmt.Println(trustP7)
	}

}
