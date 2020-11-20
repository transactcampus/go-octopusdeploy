package examples

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
)

// UpdateEnvironmentExtensionSettingsExample Test function
func UpdateEnvironmentExtensionSettingsExample() {
	var (
		octopusURL    string = os.Getenv("OCTOPUS_URL")
		apiKey        string = os.Getenv("API_KEY")
		spaceID       string = os.Getenv("SPACE_ID")
		environmentID string = os.Getenv("ENVRIONMENT_ID")
		extensionID   string = os.Getenv("EXTENSION_ID")
		valueKey      string = "JiraEnvironmentType"
		valueValue    string = "staging"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		fmt.Printf("%v", err)
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// delete environment
	environment, err := client.Environments.GetByID(environmentID)
	if err != nil {
		fmt.Printf("%v", err)
		_ = fmt.Errorf("error deleting environment: %v", err)
		return
	}

	//jsonEnv, err := json.Marshal(environment)
	//fmt.Printf("%s\n\t", jsonEnv)

	for _, extensionSetting := range environment.ExtensionSettings {
		if extensionSetting.ExtensionID == extensionID {
			extensionSetting.Values[valueKey] = valueValue
		}
	}

	resp, err := client.Environments.Update(environment)
	if err != nil {
		fmt.Printf("%v", err)
		_ = fmt.Errorf("error deleting environment: %v", err)
		return
	}

	jsonEnv, err := json.Marshal(resp)
	fmt.Printf("%s\n\t", jsonEnv)

}

func main() {
	UpdateEnvironmentExtensionSettingsExample()
}
