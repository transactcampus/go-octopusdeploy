package examples

import (
	"fmt"
	"net/url"
	"os"

	// "encoding/json"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

func main() {
	var (
		// Declare working variables
		octopusURL string = os.Getenv("OCTOPUS_URL")
		apiKey     string = os.Getenv("API_KEY")
		runbookID  string = os.Getenv("RUNBOOK_ID")
		roleName   string = os.Getenv("ROLE_NAME")
		scriptBody string = "Write-Host \"Hello world\""
		spaceID    string = os.Getenv("SPACE_ID")
		stepName   string = "Run a script"
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

	/*
		rb, _ := json.Marshal(runbook)
		fmt.Printf("%s\n\t", rb)
		os.Exit(0)
		//*/

	// Get the deployment process
	runbookProcess, err := client.RunbookProcesses.GetByID(runbook.RunbookProcessID)
	// jsonEncoding, err := json.Marshal(runbookProcess)

	/*
		rb, _ := json.Marshal(runbookProcess)
		fmt.Printf("%s\n\t", rb)
		os.Exit(0)
		//*/

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

	/*
		rb, _ := json.Marshal(runbookProcess)
		fmt.Printf("%s\n\t", rb)
		os.Exit(0)
		//*/

	// Update process
	rbpr, err := client.RunbookProcesses.Update(*runbookProcess)

	/*
		rb, _ := json.Marshal(rbpr)
		fmt.Printf("%s\n\t", rb)
		os.Exit(0)
		//*/

	if err != nil {
		// TODO: handle error
	}
}
