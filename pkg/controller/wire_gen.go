// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

// Injectors from wire.go:

func InitializeScenarioController() (ScenarioController, error) {
	restClient := NewRestClient()
	assertionController := NewAssertionController()
	stepController := NewStepController(restClient, assertionController)
	scenarioController := NewScenarioController(stepController)
	return scenarioController, nil
}
