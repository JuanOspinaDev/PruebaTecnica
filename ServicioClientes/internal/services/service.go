package services

import "ServicioClientes/internal/repository"

// ClienteService define las funcionalidades para los clientes
type ClienteService interface {
	GetClientes() error
}

type clienteService struct {
	clienteRepo  repository.ClienteRepository
}

func NewClienteService(clienteRepo repository.ClienteRepository) *clienteService {
	return &clienteService{
		clienteRepo:  clienteRepo,
	}
}