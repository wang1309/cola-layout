package web

import (
	"github.com/gin-gonic/gin"
	"github.com/wang1309/cola-layout/adapter/web/router/customer"
	"net/http"
)

func Init(r *gin.Engine) {
	r.GET("/customer/list/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": customer.GetName(),
		})
	})
}
