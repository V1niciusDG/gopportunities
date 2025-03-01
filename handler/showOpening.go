package handler

import (
	"fmt"
	"net/http"

	"github.com/V1niciusDG/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id== "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "urlParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err!=nil{
		sendError(ctx, http.StatusNotFound,fmt.Sprintf("opening with id: %s not found", id))
		return
	} 
	sendSucess(ctx, "show-openings", opening)
}