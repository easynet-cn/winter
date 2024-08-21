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
		status := http.StatusBadRequest

		if status != businessError.Status {
			status = businessError.Status
		}

		RenderResult(ctx, status, businessError)
	} else {
		RenderResult(ctx, http.StatusBadRequest, NewBadRequestBusinessError(err.Error()))
	}
}

func RenderNotFoundResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		status := http.StatusNotFound

		if status != businessError.Status {
			status = businessError.Status
		}

		RenderResult(ctx, status, businessError)
	} else {
		RenderResult(ctx, http.StatusNotFound, NewNotFoundBusinessError(err.Error()))
	}
}

func RenderUnauthorizedResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		status := http.StatusUnauthorized

		if status != businessError.Status {
			status = businessError.Status
		}

		RenderResult(ctx, status, businessError)
	} else {
		RenderResult(ctx, http.StatusUnauthorized, NewUnauthorizedBusinessError(err.Error()))
	}
}

func RenderForbiddenResult(ctx *gin.Context, err error) {
	if businessError, ok := err.(*BusinessError); ok {
		status := http.StatusForbidden

		if status != businessError.Status {
			status = businessError.Status
		}

		RenderResult(ctx, status, businessError)
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
		status := http.StatusServiceUnavailable

		if status != businessError.Status {
			status = businessError.Status
		}

		RenderResult(ctx, status, businessError)
	} else {
		RenderResult(ctx, http.StatusServiceUnavailable, NewServiceUnavailableBusinessError(err.Error()))
	}
}
