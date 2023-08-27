package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/datagenx/license-generator/internal/validate"
)

func PostValidate(ctx *gin.Context) {

	// reading and binding the post data
	inpBody := validate.Slic{}
	ctx.ShouldBindJSON(&inpBody)

	// Validating the input License
	if err := inpBody.Validate(); err != nil {
		errMsg(ctx, http.StatusForbidden, err)
		return
	}
	msg(ctx, http.StatusOK, "License is valid.")
}
