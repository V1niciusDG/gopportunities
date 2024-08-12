package handler

import (
	"net/http"

	"github.com/V1niciusDG/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error;err != nil{
		sendError(ctx, http.StatusNotFound, "opening not found in database")
	}
	sendSucess(ctx,"list-opening", openings)
}