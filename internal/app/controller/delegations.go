package controller

import (
	"net/http" // For handling HTTP requests and responses

	"github.com/FayxChance/tha-backend-tezos/internal/app/service" // Service layer handling delegation logic
	"github.com/FayxChance/tha-backend-tezos/model"                // Data models
	"github.com/gin-gonic/gin"                                     // Gin framework for handling routing
)

// DelegationsController handles delegation-related HTTP requests.
// It uses the DelegationsService to retrieve and return delegation data.
type DelegationsController struct {
	DelegationsSvc service.DelegationsService // Service that manages delegation logic
}

// Delegations is the handler function for the `/xtz/delegations` endpoint.
// It retrieves delegations from the service layer and returns them in JSON format.
func (d *DelegationsController) Delegations(c *gin.Context) {
	// Fetch delegations using the service layer
	delegations, err := d.DelegationsSvc.Delegations()
	if err != nil {
		// If an error occurs, return a 500 status code
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Return the delegations in JSON format, wrapping them in the Data struct
	c.JSON(200, model.Data{
		Data: delegations,
	})
}
