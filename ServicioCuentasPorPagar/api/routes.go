package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura y retorna las rutas para el servicio CuentasPorPagar.
func SetUpRouter() *gin.Engine {
	router := gin.Default()

	return router
}
