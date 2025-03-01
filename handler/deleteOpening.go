package handler

import (
	"fmt"
	"net/http"

	"github.com/V1niciusDG/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)


func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id== "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "urlParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err!=nil{
		sendError(ctx, http.StatusNotFound,fmt.Sprintf("opening with id: %s not found", id))
	} 

	if err := db.Delete(&opening).Error; err != nil{
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound,fmt.Sprintf("error deleting opening with id %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", opening)
}