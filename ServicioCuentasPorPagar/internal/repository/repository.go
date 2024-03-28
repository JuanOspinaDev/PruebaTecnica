package repository

import (
	"ServicioCuentasPorPagar/internal/models"
	"errors"
	"sync"
	"time"
)

var (
	ErrCuentaNoEncontrada = errors.New("cuenta por pagar no encontrada")
)

// CuentaPendienteRepository define la interfaz para el repositorio
type CuentaPendienteRepository interface {
	GetCuentas() ([]models.CuentaPorPagar, error)
	GetCuentasPorClienteID(clienteID int64) ([]models.CuentaPorPagar, error)
	CreateCuentasPendientes(cuenta models.CuentaPorPagar) error
}

// mockCuentasPendientesRepository es una implementación de la interfaz
type mockCuentaPendienteRepository struct {
	sync.Mutex
	cuentas map[int64]models.CuentaPorPagar
	nextID  int64
}

// NewMockCuentasPendienteRepository crea un nuevo repositorio mock
func NewMockCuentaPendienteRepository() CuentaPendienteRepository {
	repo := &mockCuentaPendienteRepository{
		cuentas: make(map[int64]models.CuentaPorPagar),
		nextID:  1,
	}

	// Inicializa el repositorio con datos mockeados
	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		PedidoID:         1001,
		Concepto:         "Servicio de reencauche",
		NumeroDeCuota:    1,
		Valor:            50000,
		FechaVencimiento: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		PedidoID:         1001,
		Concepto:         "Servicio de reencauche",
		NumeroDeCuota:    2,
		Valor:            50000,
		FechaVencimiento: time.Now().Add(60 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        3,
		PedidoID:         1002,
		Concepto:         "Alineacion y balanceo",
		NumeroDeCuota:    1,
		Valor:            80000,
		FechaVencimiento: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		PedidoID:         1003,
		Concepto:         "Mantenimiento Preventivo",
		NumeroDeCuota:    1,
		Valor:            70000,
		FechaVencimiento: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		PedidoID:         1003,
		Concepto:         "Mantenimiento Preventivo",
		NumeroDeCuota:    2,
		Valor:            70000,
		FechaVencimiento: time.Now().Add(60 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        1,
		PedidoID:         1003,
		Concepto:         "Mantenimiento Preventivo",
		NumeroDeCuota:    3,
		Valor:            70000,
		FechaVencimiento: time.Now().Add(90 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	repo.cuentas[repo.nextID] = models.CuentaPorPagar{
		ID:               repo.nextID,
		ClienteID:        4,
		PedidoID:         1004,
		Concepto:         "Reparación de llanta",
		NumeroDeCuota:    1,
		Valor:            25000,
		FechaVencimiento: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
	}
	repo.nextID++

	return repo
}

// GetCuentas retorna todas las cuentas por pagar
func (r *mockCuentaPendienteRepository) GetCuentas() ([]models.CuentaPorPagar, error) {
	r.Lock()
	defer r.Unlock()
	cuentas := make([]models.CuentaPorPagar, 0, len(r.cuentas))
	for _, cuenta := range r.cuentas {
		cuentas = append(cuentas, cuenta)
	}
	return cuentas, nil
}

// GetCuentasPorClienteID retorna todas las cuentas por pagar
func (r *mockCuentaPendienteRepository) GetCuentasPorClienteID(clienteID int64) ([]models.CuentaPorPagar, error) {
	r.Lock()
	defer r.Unlock()
	var cuentas []models.CuentaPorPagar
	for _, cuenta := range r.cuentas {
		if cuenta.ClienteID == clienteID {
			cuentas = append(cuentas, cuenta)
		}
	}
	return cuentas, nil
}

// CreateCuentasPendientes inserta una nueva cuenta por pagar en el repositorio
func (r *mockCuentaPendienteRepository) CreateCuentasPendientes(cuenta models.CuentaPorPagar) error {
	r.Lock()
	defer r.Unlock()

	// Asignar un nuevo ID a la cuenta
	cuenta.ID = r.nextID
	r.nextID++

	// Insertar la cuenta en el mapa
	r.cuentas[cuenta.ID] = cuenta

	return nil
}
