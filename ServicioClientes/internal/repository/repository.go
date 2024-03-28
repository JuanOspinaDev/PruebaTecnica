package repository

import (
	"ServicioClientes/internal/models"
	"sync"
)

// Interfaz para repositorio
type ClienteRepository interface {
	GetClientes() ([]models.Cliente, error)
}

// mockClienteRepository implementa la interfaz
type mockClienteRepository struct {
	mu       sync.Mutex
	clientes map[string]models.Cliente
}

// NewMockClienteRepository inicializa y retorna un repositorio
func NewMockClienteRepository() ClienteRepository {
	repo := &mockClienteRepository{
		clientes: make(map[string]models.Cliente),
	}
	
	repo.clientes["1"] = models.Cliente{ID: "1", Nombre: "Juan Pérez", Email: "juancolpe@example.com", Telefono: "311123450"}
	repo.clientes["2"] = models.Cliente{ID: "2", Nombre: "Ana López", Email: "lopezana123@example.com", Telefono: "3187654321"}
	repo.clientes["3"] = models.Cliente{ID: "3", Nombre: "Paula Ospina", Email: "paulaosp@example.com", Telefono: "3234567890"}
	repo.clientes["4"] = models.Cliente{ID: "4", Nombre: "Daniel Zapata", Email: "DanielZ@example.com", Telefono: "3123454321"}
	repo.clientes["5"] = models.Cliente{ID: "5", Nombre: "Franchesca Gomez", Email: "FranGom@example.com", Telefono: "3434567890"}
	repo.clientes["6"] = models.Cliente{ID: "6", Nombre: "David Gutierrez", Email: "DavidGut01@example.com", Telefono: "3227654321"}


	return repo
}

// GetClientes retorna todos los clientes del repositorio
func (r *mockClienteRepository) GetClientes() ([]models.Cliente, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	lista := make([]models.Cliente, 0, len(r.clientes))
	for _, cliente := range r.clientes {
		lista = append(lista, cliente)
	}

	return lista, nil
}
