package winter

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RenderResult(ctx *gin.Context, status int, result any) {
	ctx.JSON(status, result)
}

func RenderSuccessResult(ctx *gin.Context, result any) {
	RenderResult(ctx, http.StatusOK, result)
}

func RenderErrorResult(ctx *gin.Context, status int, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		statusCode := status

		if statusCode != businessError.Status {
			statusCode = businessError.Status
		}

		RenderResult(ctx, statusCode, businessError)
	} else {
		RenderResult(ctx, status, NewBusinessError(status, strconv.Itoa(status), err.Error()))
	}
}

func RenderBadRequestResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusBadRequest, err)
}

func RenderNotFoundResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusNotFound, err)
}

func RenderUnauthorizedResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusUnauthorized, err)
}

func RenderForbiddenResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusForbidden, err)
}

func RenderInternalServerErrorResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusInternalServerError, err)
}

func RenderServiceUnavailableResult(ctx *gin.Context, err error) {
	RenderErrorResult(ctx, http.StatusServiceUnavailable, err)
}
