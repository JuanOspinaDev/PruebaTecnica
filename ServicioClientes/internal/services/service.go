package services

import (
	"ServicioClientes/internal/models"
	"ServicioClientes/internal/repository"
)

// ClienteService define las funcionalidades para los clientes
type ClienteService interface {
	GetClientes() ([]models.Cliente, error)
}

type clienteService struct {
	clienteRepo  repository.ClienteRepository
}

func NewClienteService(clienteRepo repository.ClienteRepository) *clienteService {
	return &clienteService{
		clienteRepo:  clienteRepo,
	}
}

// el servicio GetClientes llama al metodo de repositorio correspondiente
func (cl *clienteService) GetClientes() ([]models.Cliente, error) {
	return cl.clienteRepo.GetClientes()
}
