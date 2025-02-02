package deployment

import (
	"github.com/gorilla/mux"
)

type DeploymentConfigRouter interface {
	Init(configRouter *mux.Router)
}

type DeploymentConfigRouterImpl struct {
	deploymentRestHandler DeploymentConfigRestHandler
}

func NewDeploymentRouterImpl(deploymentRestHandler DeploymentConfigRestHandler) *DeploymentConfigRouterImpl {
	return &DeploymentConfigRouterImpl{
		deploymentRestHandler: deploymentRestHandler,
	}
}

func (router DeploymentConfigRouterImpl) Init(configRouter *mux.Router) {
	configRouter.Path("/validate").
		HandlerFunc(router.deploymentRestHandler.CreateChartFromFile).Methods("POST")
	configRouter.Path("/upload").
		HandlerFunc(router.deploymentRestHandler.SaveChart).Methods("PUT")
	configRouter.Path("/fetch").
		HandlerFunc(router.deploymentRestHandler.GetUploadedCharts).Methods("GET")
}
