package main

import (
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"
)

func verifyCertificate(cert *x509.Certificate, hostname string) error {
	if time.Now().Before(cert.NotBefore) || time.Now().After(cert.NotAfter) {
		return fmt.Errorf("Certificatul este invalid. Posibil un atac 'Man in the middle'!")
	}

	err := cert.VerifyHostname(hostname)
	if err != nil {
		return fmt.Errorf("Numele domeniului nu corespunde certificatului.")
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Eroare la argumentele din linia de comanda!")
		return
	}
	hostname := os.Args[1]
	url := "https://" + hostname

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Eroare la request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	tlsConnState := resp.TLS
	if tlsConnState == nil || len(tlsConnState.PeerCertificates) == 0 {
		fmt.Println("Certificatul TLS nu a putut fi obtinut!")
		return
	}
	cert := tlsConnState.PeerCertificates[0]

	fmt.Printf("Certificatul este valabil pana la: %s\n", cert.NotAfter.Format("02 Jan 2006 15:04:05 "))

	err = verifyCertificate(cert, hostname)
	if err != nil {
		fmt.Printf("Site-ul este nesigur. Posibil un atac 'Man in the middle'!:\n %v\n", err)
	} else {
		fmt.Println("Site-ul este sigur!")
	}
}
