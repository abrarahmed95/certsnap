package certsnap

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type CertificateInfo struct {
	URL    string    `json:"url"`
	Expiry time.Time `json:"expiry"`
	Error  error     `json:"-"`
}

func (ci *CertificateInfo) ToJSON() string {
	data, err := json.Marshal(ci)

	if err != nil {
		log.Fatal("Error converting to JSON")
	}

	return string(data)
}

func (ci *CertificateInfo) ToString() string {
	remainingDays := int(time.Until(ci.Expiry).Hours() / 24)

	return fmt.Sprintf("SSL certificate for %s expires on %s. Remaining time: %d days", ci.URL, ci.Expiry.Format("2006-01-02 15:04:05"), remainingDays)
}

func GetCertificateExpiryDate(url string) (time.Time, error) {
	targetUrl := fmt.Sprintf("%s:443", url)

	conn, err := tls.Dial("tcp", targetUrl, nil)

	if err != nil {
		return time.Time{}, err
	}

	defer conn.Close()

	certificates := conn.ConnectionState().PeerCertificates

	if len(certificates) == 0 {
		fmt.Println("no certificates found")
	}

	return certificates[0].NotAfter, nil
}

func HasValidScheme(value string) bool {
	return !(!strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://"))
}
