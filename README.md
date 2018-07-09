# go-octopusdeploy
A Go wrapper for the Octopus Deploy REST API.

Written to be used in the [Octopus Deploy Terraform Provider](https://github.com/MattHodge/terraform-provider-octopusdeploy).

[![Build status](https://ci.appveyor.com/api/projects/status/5t5gbqjyl8hpou52?svg=true)](https://ci.appveyor.com/project/MattHodge/go-octopusdeploy)

> :warning: The Octopus Deploy REST Client is in heavy development.

# Go Dependencies
* Dependencies are managed using [dep](https://golang.github.io/dep/docs/new-project.html)

```bash
# Vendor new modules
dep ensure
```

# Using the main.go Example

```bash
export OCTOPUS_URL=http://localhost:8081/
export OCTOPUS_APIKEY=API-FAKEAPIKEYFAKEAPIKEY

go run main.go # creates a project
```

# Links
* Octopus Deploy API Examples: https://github.com/OctopusDeploy/OctopusDeploy-Api
* Octopus Deploy Clients: https://github.com/OctopusDeploy/OctopusClients
