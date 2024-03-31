package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// NewModelSplit godoc
// @Summary      create a new model split
// @Description  create a new model split
// @Tags         Model split
// @Accept       json
// @Produce      json
// @Param        request body requests.ModelSplitTransactionReq true "model split"
// @Success      200  {object} responses.SuccessRes
// @Router       /v2/model-split [post]
func NewModelSplit(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.ModelSplitTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.CreateModelSplit(userIDUint, req.ModelSplitName, req.ModelSplits)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("MODEL_SPLIT_CREATED", fmt.Sprintf("model split %s created", req.ModelSplitName)))
}
