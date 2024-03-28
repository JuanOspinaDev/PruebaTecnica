package services

import (
	"ServicioCuentasPorPagar/internal/models"
	"ServicioCuentasPorPagar/internal/repository"
)

// CuentaPendienteService define las funcionalidades para los clientes
type CuentaPendienteService interface {
	GetCuentas() ([]models.CuentaPorPagar, error)
	GetCuentasPorCliente(clienteID int64) ([]models.CuentaPorPagar, error)
}

type cuentaPendienteService struct {
	clienteRepo repository.CuentaPendienteRepository
}

func NewClienteService(clienteRepo repository.CuentaPendienteRepository) *cuentaPendienteService {
	return &cuentaPendienteService{
		clienteRepo: clienteRepo,
	}
}

// GetClientes llama al metodo de repositorio correspondiente
func (cp *cuentaPendienteService) GetCuentas() ([]models.CuentaPorPagar, error) {
	return cp.clienteRepo.GetCuentas()
}

// GetClientes llama al metodo de repositorio correspondiente
func (cp *cuentaPendienteService) GetCuentasPorCliente(clienteID int64) ([]models.CuentaPorPagar, error) {
	return cp.clienteRepo.GetCuentasPorClienteID(clienteID)
}
