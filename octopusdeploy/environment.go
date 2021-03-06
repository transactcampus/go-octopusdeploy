package octopusdeploy

import "github.com/go-playground/validator/v10"

type Environments struct {
	Items []*Environment `json:"Items"`
	PagedResults
}

type Environment struct {
	AllowDynamicInfrastructure bool                       `json:"AllowDynamicInfrastructure"`
	Description                string                     `json:"Description,omitempty"`
	Name                       string                     `json:"Name" validate:"required"`
	SortOrder                  int                        `json:"SortOrder"`
	UseGuidedFailure           bool                       `json:"UseGuidedFailure"`
	ExtensionSettings          []*ExtensionSettingsValues `json:"ExtensionSettings,omitempty"`

	resource
}

// NewEnvironment initializes an Environment with a name.
func NewEnvironment(name string) *Environment {
	return &Environment{
		AllowDynamicInfrastructure: false,
		Name:                       name,
		SortOrder:                  0,
		UseGuidedFailure:           false,
		resource:                   *newResource(),
	}
}

// Validate checks the state of the environment and returns an error if
// invalid.
func (e *Environment) Validate() error {
	return validator.New().Struct(e)
}
