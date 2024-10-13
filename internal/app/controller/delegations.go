package controller

import (
	"net/http"

	"github.com/FayxChance/tha-backend-tezos/internal/app/service"
	"github.com/gin-gonic/gin"
)

type DelegationsController struct {
	DelegationsSvc service.DelegationsService
}

func (d *DelegationsController) Delegations(c *gin.Context) {
	delegations, err := d.DelegationsSvc.Delegations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, delegations)

}
