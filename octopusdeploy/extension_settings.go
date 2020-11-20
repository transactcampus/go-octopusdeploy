package octopusdeploy

import "github.com/go-playground/validator/v10"

// ExtensionSettings Environment Extension Settings
type ExtensionSettings struct {
	ExtensionID string            `json:"ExtensionId"`
	Values      map[string]string `json:"Values"`

	//resource
}

// NewExtensionSettings initializes a ExtensionSettings with id.
func NewExtensionSettings(id string) *ExtensionSettings {
	return &ExtensionSettings{
		ExtensionID: id, // jira-integration
		Values:      map[string]string{},
		//resource:   *newResource(),
	}
}

// Validate checks the state of the deployment step and returns an error if
// invalid.
func (d ExtensionSettings) Validate() error {
	return validator.New().Struct(d)
}
