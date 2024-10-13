package router

import (
	"github.com/FayxChance/tha-backend-tezos/internal/app/controller"
	"github.com/FayxChance/tha-backend-tezos/internal/app/service"
	"github.com/FayxChance/tha-backend-tezos/internal/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router         *gin.Engine
	DelegationCtrl controller.DelegationsController
}

func (r *Router) SetupRouter(db persistence.SQLite3Database) error {
	r.Router = gin.Default()

	r.Router.GET("/xtz/delegations", r.DelegationCtrl.Delegations)

	var delegationsController controller.DelegationsController
	delegationsController.DelegationsSvc = service.DelegationsService{
		DB: db,
	}
	r.DelegationCtrl = delegationsController
	return nil
}
