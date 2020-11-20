package octopusdeploy

// RunbookProcesses defines a collection of runbooks processes with built-in support for paged
// results.
type RunbookProcesses struct {
	Items []*RunbookProcess `json:"Items"`
	PagedResults
}

// RunbookProcess represents a deployment process.
type RunbookProcess struct {
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	RunbookID      string           `json:"RunbookId,omitempty"`
	SpaceID        string           `json:"SpaceId,omitempty"`
	Steps          []DeploymentStep `json:"Steps"`
	Version        int32            `json:"Version"`

	resource
}

// NewRunbookProcess creates and initializes a runbook process.
func NewRunbookProcess(runbookID string, projectID string) *RunbookProcess {
	return &RunbookProcess{
		RunbookID: runbookID,
		ProjectID: projectID,
		resource:  *newResource(),
	}
}
