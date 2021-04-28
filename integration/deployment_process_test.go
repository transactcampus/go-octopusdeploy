package integration

import (
	"testing"

	"github.com/transactcampus/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func GetTestDeploymentProcess(t *testing.T, client *octopusdeploy.Client, project *octopusdeploy.Project) *octopusdeploy.DeploymentProcess {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	require.NotNil(t, deploymentProcess)
	require.NoError(t, err)
	require.Equal(t, project.DeploymentProcessID, deploymentProcess.GetID())

	return deploymentProcess
}

func TestDeploymentProcessGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess := GetTestDeploymentProcess(t, client, project)
	require.NotNil(t, deploymentProcess)
}

func TestDeploymentProcessGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	deploymentProcesses, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	numberOfDeploymentProcesses := len(deploymentProcesses)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	totalDeploymentProcesses, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	assert.Equal(t, len(totalDeploymentProcesses), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}

func TestDeploymentProcessUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess := GetTestDeploymentProcess(t, client, project)
	require.NotNil(t, deploymentProcess)

	deploymentActionWindowService := octopusdeploy.NewDeploymentAction("Install Windows Service", "Octopus.WindowService")
	deploymentActionWindowService.Properties["Octopus.Action.EnabledFeatures"] = octopusdeploy.NewPropertyValue("Octopus.Features.WindowService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.FeedId"] = octopusdeploy.NewPropertyValue("feeds-nugetfeed", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.DownloadOnTentacle"] = octopusdeploy.NewPropertyValue("False", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.PackageId"] = octopusdeploy.NewPropertyValue("Newtonsoft.Json", false)
	deploymentActionWindowService.Properties["Octopus.Action.SubstituteInFiles.Enabled"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.CreateOrUpdateService"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.DisplayName"] = octopusdeploy.NewPropertyValue("my display name", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.Description"] = octopusdeploy.NewPropertyValue("my desc", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ExecutablePath"] = octopusdeploy.NewPropertyValue("bin\\Myservice.exe", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ServiceAccount"] = octopusdeploy.NewPropertyValue("LocalSystem", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ServiceName"] = octopusdeploy.NewPropertyValue("My service name", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.StartMode"] = octopusdeploy.NewPropertyValue("auto", false)
	deploymentActionWindowService.Properties["Octopus.Action.SubstituteInFiles.TargetFiles"] = octopusdeploy.NewPropertyValue("*.sh", false)
	deploymentActionWindowService.Properties["test"] = octopusdeploy.NewPropertyValue("foo", true)

	deploymentStep := octopusdeploy.NewDeploymentStep(getRandomName())
	deploymentStep.Actions = append(deploymentStep.Actions, *deploymentActionWindowService)
	deploymentStep.Properties["Octopus.Action.TargetRoles"] = octopusdeploy.NewPropertyValue("octopus-server", false)

	deploymentProcess.Steps = append(deploymentProcess.Steps, *deploymentStep)

	updatedDeploymentProcess, err := client.DeploymentProcesses.Update(deploymentProcess)
	require.NoError(t, err)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
