package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wang1309/cola-layout/adapter/web"
)

func main() {
	r := gin.Default()

	web.Init(r)

	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
