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

var (
	p12path   = flag.String("path", "", "Path to client P12")
	p12pass   = flag.String("pass", "", "Password for client P12")
	enrollPtr = flag.Bool("enroll", false, "Bool to enroll")
	renewPtr  = flag.Bool("renew", false, "Bool to renew")
	trustPtr  = flag.Bool("trust", false, "Bool to retrieve trust")
	certFile  = flag.String("cert", "someCertFile", "A PEM eoncoded certificate file.")
	keyFile   = flag.String("key", "someKeyFile", "A PEM encoded private key file.")
	caFile    = flag.String("CA", "someCertCAFile", "A PEM eoncoded CA's certificate file.")
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

/* Disabled temporarily for debugging
func renew() {

	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	resp, err := client.Get("https://goldportugal.local:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
*/

func main() {

	flag.Parse()

	fmt.Println("---Begin Flag Debug---")
	fmt.Println("path:", *p12path)
	fmt.Println("p12pass:", *p12pass)
	fmt.Println("enroll:", *enrollPtr)
	fmt.Println("renew:", *renewPtr)
	fmt.Println("trust:", *trustPtr)

	fmt.Println("---End Flag Debug---")

	if *enrollPtr {
		log.Println("Initiating EST Simple Enroll.")
	} else if *renewPtr {
		log.Println("Initiating EST Simple Reenrollment.")
		renew()
	} else if *trustPtr {
		log.Println("Retrieving Certificate Authority trust.")
		gettrust() // runs trust function
	}
}
