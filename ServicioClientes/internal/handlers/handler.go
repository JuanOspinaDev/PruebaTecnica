package handlers

import (
	"ServicioClientes/internal/services"

	"github.com/gin-gonic/gin"
)

type clienteHandler struct {
	clienteService services.ClienteService
}

func NewClienteHandler(clienteService services.ClienteService) *clienteHandler {
	return &clienteHandler{
		clienteService: clienteService,
	}
}

// GetClientes define un controlador para obtener una lista de los clientes
func (cl *clienteHandler) GetClientes(c *gin.Context) {
	cl.clienteService.GetClientes()
}
