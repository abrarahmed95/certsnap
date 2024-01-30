package certsnap

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/abrarahmed95/certsnap/pkg/certsnap"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ensureValidURL(value string) string {
	if !certsnap.HasValidScheme(value) {
		return "https://" + value
	}

	return value
}

func printResultsJSON(results <-chan certsnap.CertificateInfo) {
	var resultSlice []certsnap.CertificateInfo

	for result := range results {
		resultSlice = append(resultSlice, result)
	}

	encoder := json.NewEncoder(os.Stdout)

	if err := encoder.Encode(resultSlice); err != nil {
		fmt.Println("Error checking SSL certificate:", err)
	}
}

func printResultsTable(results <-chan certsnap.CertificateInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Domain", "Expiry date", "Remaining Days"})

	i := 0
	for result := range results {
		i++

		if result.Error != nil {
			t.AppendRow(table.Row{
				i,
				result.URL,
				"Error checking SSL certificate",
				nil,
			})
		} else {
			t.AppendRow(table.Row{
				i,
				result.URL,
				result.Expiry,
				result.GetRemainingDays(),
			})
		}
	}

	t.AppendSeparator()
	t.Render()
}

func printResults(results <-chan certsnap.CertificateInfo, jsonOut bool) {
	if jsonOut {
		printResultsJSON(results)
	} else {
		printResultsTable(results)
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
