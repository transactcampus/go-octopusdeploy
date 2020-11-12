package examples

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

// ReplaceCertificateExample shows how to replace an existing certificate using go-octopusdeploy.
func ReplaceCertificateExample() {
	var (
		apiKey          string = "API-YOUR_API_KEY"
		certificateName string = "MyCertificate"
		octopusURL      string = "https://your_octopus_url"
		pfxFilePath     string = "path\\to\\pfxfile.pfx"
		pfxFilePassword string = "PFX-file-password"
		spaceID         string = "default"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// get certificates
	certificateList, err := client.Certificates.GetByPartialName(certificateName)
	if err != nil {
		// TODO: handle error
	}

	// find the certificate with a specific name
	certificate := certificateList[0]

	file, err := os.Open(pfxFilePath)
	if err != nil {
		// TODO: handle error
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		// TODO: handle error
	}

	// Convert file to base64
	base64Certificate := base64.StdEncoding.EncodeToString(data)

	// Replace certificate
	replacementCertificate := octopusdeploy.NewReplacementCertificate(base64Certificate, pfxFilePassword)
	_, err = client.Certificates.Replace(certificate.GetID(), replacementCertificate)
	if err != nil {
		// TODO: handle error
	}
}
