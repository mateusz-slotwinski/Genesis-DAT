package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

type SciencesController struct{}

func (v SciencesController) GetGroups(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", gin.H{})
}
