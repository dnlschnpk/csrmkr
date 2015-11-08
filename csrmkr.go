package main

import(
	"fmt"
	"flag"
	"crypto/rsa"
	"crypto/rand"
	"encoding/pem"
	"crypto/x509"
	"crypto/x509/pkix"
//	"time"
	"os"
)

func keyGen(rsaBits int) {


}

func main(){
	// Parsing command line flags
	rsaBits               := flag.Int("bits", 2048, "Size of RSA key to generate.")
	crtCountry            := flag.String("C", "DE", "Subject Country")
	crtState              := flag.String("ST", "Nordrhein-Westfalen", "Subject Country")
	crtLocation           := flag.String("L", "Guetersloh", "Subject Location")
	crtOrganization       := flag.String("O", "arvato Systems GmbH", "Subject Organization")
	crtOrganizationalUnit := flag.String("OU", "Application Services", "Subject Organizational Unit")
	crtCommonName         := flag.String("CN", "*.arvato-systems.de", "Subject Common Name")
	crtSAN                := flag.String("SAN", "*.arvato-systems.de,arvato-systems.de", "Subject alternative Names")
	flag.Parse()

	// Generating private key
	privateKey, err := rsa.GenerateKey(rand.Reader, *rsaBits)
	if err != nil {
            fmt.Println(err)
        }

	// Format the key for output
//	keyFile, err := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	keyOutput := os.Stdout
	pem.Encode(keyOutput, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	template := &x509.CertificateRequest{
			Subject: pkix.Name{
				CommonName:   *crtCommonName,
				Country: []string{"AU"},
				Province: []string{"Some-State"},
				Locality: []string{"MyCity"},
				Organization: []string{"Company Ltd"},
				OrganizationalUnit: []string{"IT"},
			},
			EmailAddresses: []string{"test@email.com"},
			DNSNames : []string{"example.abc123.com"},
		}

	certCSR, err := x509.CreateCertificateRequest(rand.Reader, template, privateKey)
	if err != nil {
            fmt.Println(err)
        }
//	pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
//	pem.Encode(os.Stdout,
//fmt.Println(keyBytes)
}
