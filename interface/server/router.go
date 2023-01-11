package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(app *gin.Engine, handler *handler) {
	customerApi := app.Group("/customer/v1")

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "REST API CRUD")
	})

	customerApi.GET("/getAllData", handler.customerHandler.getAllDataCustomer())
	customerApi.POST("/add", handler.customerHandler.customerSaveData())
	customerApi.POST("/getById", handler.customerHandler.getDataCustomerById())
	customerApi.PUT("/update", handler.customerHandler.customerUpdate())
	customerApi.DELETE("/delete", handler.customerHandler.customerDelete())

}
