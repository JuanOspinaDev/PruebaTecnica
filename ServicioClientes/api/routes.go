package api

import (
	"ServicioClientes/internal/handlers"
	"ServicioClientes/internal/repository"
	"ServicioClientes/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura y retorna las rutas para el servicio Clientes.
func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// Se realiza la inyeccion de dependencias
	clienteRepo := repository.NewMockClienteRepository()
	clienteService := services.NewClienteService(clienteRepo)
	clienteHandler := handlers.NewClienteHandler(clienteService)

	// Se establece el grupo de ruta clientes
	clientes := router.Group("/clientes")
	{
		// ruta para todos los clientes
		clientes.GET("/", clienteHandler.GetClientes)
	}
	return router
}
