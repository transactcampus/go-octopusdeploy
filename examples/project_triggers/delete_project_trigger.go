package examples

import (
	"fmt"
	"net/url"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

func DeleteProjectTriggerExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// project trigger values
		projectTriggerID string = "project-trigger-id"
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

	// delete project trigger
	err = client.ProjectTriggers.DeleteByID(projectTriggerID)
	if err != nil {
		_ = fmt.Errorf("error deleting project trigger: %v", err)
		return
	}

	fmt.Printf("project trigger deleted: (%s)\n", projectTriggerID)
}
