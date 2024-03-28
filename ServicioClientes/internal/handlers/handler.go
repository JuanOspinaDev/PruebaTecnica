package handlers

import (
	"ServicioClientes/internal/services"
	"net/http"

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
	clientes, err := cl.clienteService.GetClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clientes)

}
