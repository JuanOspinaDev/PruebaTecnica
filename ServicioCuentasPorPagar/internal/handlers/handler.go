package handlers

import (
	"ServicioCuentasPorPagar/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cuentaPendienteHandler struct {
	cuentaPendienteService services.CuentaPendienteService
}

func NewClienteHandler(cuentaPendienteService services.CuentaPendienteService) *cuentaPendienteHandler {
	return &cuentaPendienteHandler{
		cuentaPendienteService: cuentaPendienteService,
	}
}

// GetCuentas es el controlador para obtener todas las cuentas pendientes
func (cp *cuentaPendienteHandler) GetCuentas(c *gin.Context) {
	cuentas, err := cp.cuentaPendienteService.GetCuentas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cuentas)
}

// GetCuentasPorCliente es el controlador para obtener cuentas por pagar de un cliente específico
func (cp *cuentaPendienteHandler) GetCuentasPorCliente(c *gin.Context) {
	clienteIDStr := c.Param("clienteId")

	clienteID, err := strconv.ParseInt(clienteIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	cuentas, err := cp.cuentaPendienteService.GetCuentasPorCliente(clienteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cuentas)
}
