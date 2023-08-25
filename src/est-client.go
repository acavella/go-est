package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func gettrust() {

	viper.SetConfigFile("config.env")
	viper.ReadInConfig()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", viper.GetString("CAURL"), nil)
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
	trustP7 := fmt.Sprintf("-----BEGIN P7B-----\n%s\n-----END P7B-----\n", bodyText)

	fmt.Println(trustP7)
}

func main() {

	p12path := flag.String("path", "", "Path to client P12")
	p12pass := flag.String("pass", "", "Password for client P12")
	enrollPtr := flag.Bool("enroll", false, "Bool to enroll")
	renewPtr := flag.Bool("renew", false, "Bool to renew")
	trustPtr := flag.Bool("trust", false, "Bool to retrieve trust")

	flag.Parse()

	fmt.Println("---Begin Flag Validation---")
	fmt.Println("path:", *p12path)
	fmt.Println("p12pass:", *p12pass)
	fmt.Println("enroll:", *enrollPtr)
	fmt.Println("renew:", *renewPtr)
	fmt.Println("trust:", *trustPtr)
	fmt.Println("---End Flag Validation---")

	if *enrollPtr {
		log.Println("Initiating EST Simple Enroll.")
	} else if *renewPtr {
		log.Println("Initiating EST Simple Reenrollment.")
	} else if *trustPtr {
		log.Println("Retrieving Certificate Authority trust.")
		gettrust() // runs trust function
	}
}
