package certsnap

import (
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/abrarahmed95/certsnap/pkg/certsnap"
)

func ensureValidURL(value string) string {
	if !certsnap.HasValidScheme(value) {
		return "https://" + value
	}

	return value
}

func printResultDetails(result certsnap.CertificateInfo, jsonOut bool) {
	if jsonOut {
		fmt.Println(result.ToJSON())
	} else {
		fmt.Println(result.ToString())
	}
}

func printResults(results <-chan certsnap.CertificateInfo, jsonOut bool) {
	for result := range results {
		if result.Error != nil {
			log.Printf("Error checking SSL certificate for %s: %v\n", result.URL, result.Error)
		} else {
			printResultDetails(result, jsonOut)
		}
	}
}

func checkCertificate(host string, results chan<- certsnap.CertificateInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	expiry, err := certsnap.GetCertificateExpiryDate(host)

	results <- certsnap.CertificateInfo{
		URL:    host,
		Expiry: expiry,
		Error:  err,
	}
}

func ValidateAndCheckCertificates(args []string, wg *sync.WaitGroup, jsonOut bool) {
	if len(args) < 1 {
		log.Fatal(`You didn't specify the additional arguments\Use the --help flag for comprehensive help on how to use this tool`)
	}

	results := make(chan certsnap.CertificateInfo, len(args))

	for _, value := range args {
		validURL := ensureValidURL(value)

		parsedURL, err := url.Parse(validURL)

		if err != nil {
			log.Fatal(err)
		}

		if parsedURL.Scheme == "" || parsedURL.Host == "" {
			log.Fatal("Invalid URL")
		}

		wg.Add(1)
		go checkCertificate(parsedURL.Host, results, wg)
	}

	wg.Wait()
	close(results)

	printResults(results, jsonOut)
}
