package winter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderResult(ctx *gin.Context, status int, result any) {
	ctx.JSON(status, result)
}

func RenderOkResult(ctx *gin.Context, result any) {
	RenderResult(ctx, http.StatusOK, result)
}

func RenderBadRequestResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusBadRequest, businessError)
	} else {
		RenderResult(ctx, http.StatusBadRequest, NewBadRequestBusinessError(err.Error()))
	}
}

func RenderNotFoundResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusNotFound, businessError)
	} else {
		RenderResult(ctx, http.StatusNotFound, NewNotFoundBusinessError(err.Error()))
	}
}

func RenderUnauthorizedResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusUnauthorized, businessError)
	} else {
		RenderResult(ctx, http.StatusUnauthorized, NewUnauthorizedBusinessError(err.Error()))
	}
}

func RenderForbiddenResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusForbidden, businessError)
	} else {
		RenderResult(ctx, http.StatusForbidden, NewForbiddenBusinessError(err.Error()))
	}
}

func RenderInternalServerErrorResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusInternalServerError, businessError)
	} else {
		RenderResult(ctx, http.StatusInternalServerError, NewInternalServerErrorBusinessError(err.Error()))
	}
}

func RenderServiceUnavailableResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		RenderResult(ctx, http.StatusServiceUnavailable, businessError)
	} else {
		RenderResult(ctx, http.StatusServiceUnavailable, NewServiceUnavailableBusinessError(err.Error()))
	}
}
