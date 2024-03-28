package api

import (
	"ServicioCuentasPorPagar/internal/handlers"
	"ServicioCuentasPorPagar/internal/repository"
	"ServicioCuentasPorPagar/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura y retorna las rutas para el servicio CuentasPorPagar.
func SetupRouter(repo repository.CuentaPendienteRepository) *gin.Engine {
	router := gin.Default()

	// Se realiza la inyeccion de dependencias
	cuentaPendienteService := services.NewCuentaPendienteService(repo)
	cuentaPendienteHandler := handlers.NewClienteHandler(cuentaPendienteService)

	// Se establece el grupo de ruta clientes
	clientes := router.Group("/cuentas")
	{
		// ruta para obtener todas las cuentas
		clientes.GET("/", cuentaPendienteHandler.GetCuentas)
		// ruta para obtener las cuentas por el ID del cliente
		clientes.GET("/:clienteId", cuentaPendienteHandler.GetCuentasPorCliente)
	}
	return router
}
