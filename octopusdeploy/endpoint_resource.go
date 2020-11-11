package octopusdeploy

import (
	"net/url"

	"github.com/go-playground/validator/v10"
)

type EndpointResource struct {
	AadClientCredentialSecret            string                    `json:"AadClientCredentialSecret,omitempty"`
	AadCredentialType                    string                    `json:"AadCredentialType,omitempty" validate:"omitempty,oneof=ClientCredential UserCredential"`
	AadUserCredentialUsername            string                    `json:"AadUserCredentialUsername,omitempty"`
	AadUserCredentialPassword            SensitiveValue            `json:"AadUserCredentialPassword,omitempty"`
	AccountID                            string                    `json:"AccountId"`
	ApplicationsDirectory                string                    `json:"ApplicationsDirectory,omitempty"`
	Authentication                       EndpointAuthentication    `json:"Authentication,omitempty"`
	CertificateSignatureAlgorithm        string                    `json:"CertificateSignatureAlgorithm,omitempty"`
	CertificateStoreLocation             string                    `json:"CertificateStoreLocation,omitempty"`
	CertificateStoreName                 string                    `json:"CertificateStoreName,omitempty"`
	ClientCertificateVariable            string                    `json:"ClientCertVariable,omitempty"`
	CloudServiceName                     string                    `json:"CloudServiceName"`
	ClusterCertificate                   string                    `json:"ClusterCertificate,omitempty"`
	ClusterURL                           *url.URL                  `json:"ClusterUrl" validate:"required,url"`
	CommunicationStyle                   string                    `json:"CommunicationStyle" validate:"required,oneof=AzureCloudService AzureServiceFabricCluster Ftp Kubernetes None OfflineDrop Ssh TentacleActive TentaclePassive"`
	ConnectionEndpoint                   string                    `json:"ConnectionEndpoint,omitempty"`
	Container                            DeploymentActionContainer `json:"Container,omitempty"`
	DefaultWorkerPoolID                  string                    `json:"DefaultWorkerPoolId"`
	Destination                          *OfflineDropDestination   `json:"Destination"`
	DotNetCorePlatform                   string
	Fingerprint                          string
	Host                                 string
	Namespace                            string `json:"Namespace,omitempty"`
	Port                                 int
	ProxyID                              string                  `json:"ProxyId,omitempty"`
	ResourceGroupName                    string                  `json:"ResourceGroupName,omitempty"`
	RunningInContainer                   bool                    `json:"RunningInContainer"`
	SecurityMode                         string                  `json:"SecurityMode,omitempty" validate:"omitempty,oneof=Unsecure SecureClientCertificate SecureAzureAD"`
	SensitiveVariablesEncryptionPassword SensitiveValue          `json:"SensitiveVariablesEncryptionPassword"`
	ServerCertificateThumbprint          string                  `json:"ServerCertThumbprint,omitempty"`
	SkipTLSVerification                  bool                    `json:"SkipTlsVerification"`
	Slot                                 string                  `json:"Slot"`
	StorageAccountName                   string                  `json:"StorageAccountName"`
	SwapIfPossible                       bool                    `json:"SwapIfPossible"`
	TentacleVersionDetails               *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
	Thumbprint                           string                  `json:"Thumbprint" validate:"required"`
	OctopusWorkingDirectory              string                  `json:"OctopusWorkingDirectory,omitempty"`
	UseCurrentInstanceCount              bool                    `json:"UseCurrentInstanceCount"`
	URI                                  *url.URL                `json:"Uri" validate:"required,uri"`
	WebAppName                           string                  `json:"WebAppName,omitempty"`
	WebAppSlotName                       int                     `json:"WebAppSlotName"`

	resource
}

type EndpointResources struct {
	Items []*EndpointResource `json:"Items"`
	PagedResults
}

// NewEndpoint creates and initializes an account resource with a name and type.
func NewEndpointResource(communicationStyle string) *EndpointResource {
	return &EndpointResource{
		CommunicationStyle: communicationStyle,
		resource:           *newResource(),
	}
}

// Validate checks the state of the feed resource and returns an error if
// invalid.
func (f EndpointResource) Validate() error {
	return validator.New().Struct(f)
}