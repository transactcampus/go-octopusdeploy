package main

import (
	"fmt"
	"net/url"
	// "encoding/json"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

func main() {
	var (
		// Declare working variables
		octopusURL  string = ""
		apiKey      string = ""
		runbookID   string = ""
		roleName    string = "role-name"
		scriptBody  string = "Write-Host \"Hello world\""
		spaceID     string = ""
		stepName    string = "Run a script"
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

	// Get project
	// project, err := client.Projects.GetByName(projectName)
	runbook, err := client.Runbooks.GetByID(runbookID)
	if err != nil {
		// TODO: handle error
	}
	
	// Get the deployment process
	runbookProcess, err := client.RunbookProcesses.GetByID(runbook.RunbookProcessID)
	// jsonEncoding, err := json.Marshal(runbookProcess)
	

	if err != nil {
		// TODO: handle error
	}

	// Create new step object
	newStep := octopusdeploy.NewDeploymentStep(stepName)
	newStep.Condition = "Success"
	newStep.Properties["Octopus.Action.TargetRoles"] = roleName
	
	// Create new script action
	stepAction := octopusdeploy.NewDeploymentAction(stepName)

	stepAction.ActionType = "Octopus.Script"
	stepAction.Properties["Octopus.Action.Script.ScriptBody"] = scriptBody

	// Add step action and step to process
	newStep.Actions = append(newStep.Actions, *stepAction)
	runbookProcess.Steps = append(runbookProcess.Steps, *newStep)

	// Update process
	_, err = client.RunbookProcesses.Update(*runbookProcess)

	if err != nil {
		// TODO: handle error
	}
}
